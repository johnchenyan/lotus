package storage

import (
	"context"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
)

type postItem struct {
	partIndex uint64
	di        *dline.Info
	ts        *types.TipSet
	params    *miner.SubmitWindowedPoStParams
}

type submitRes struct {
	partIndex   uint64
	di          *dline.Info
	ts          *types.TipSet
	submitState SubmitState
	err         error
	abort       context.CancelFunc
}

type fastSubmitHandler struct {
	api             changeHandlerAPI
	hcs             chan *headChange
	posts           map[abi.ChainEpoch]map[uint64]*postItem
	postNotify      chan *postItem
	submitRes       map[abi.ChainEpoch]map[uint64]*submitRes
	submitResNotify chan *submitRes
	currentCtx      context.Context
	currentTs       *types.TipSet
	revert          *types.TipSet
	shutdownCtx     context.Context
	shutdown        context.CancelFunc
}

func newFastSubmitter(api changeHandlerAPI) *fastSubmitHandler {
	ctx, cancel := context.WithCancel(context.Background())
	return &fastSubmitHandler{
		api:             api,
		hcs:             make(chan *headChange),
		postNotify:      make(chan *postItem, 16),
		posts:           make(map[abi.ChainEpoch]map[uint64]*postItem),
		submitResNotify: make(chan *submitRes),
		submitRes:       make(map[abi.ChainEpoch]map[uint64]*submitRes),
		shutdownCtx:     ctx,
		shutdown:        cancel,
	}
}

func (s *fastSubmitHandler) run() {
	defer func() {
		for _, ss := range s.submitRes {
			for _, s := range ss {
				if s.abort != nil {
					s.abort()
				}
			}
		}
	}()

	for s.shutdownCtx.Err() == nil {
		select {
		case <-s.shutdownCtx.Done():
			return

		case post := <-s.postNotify:
			s.addPost(post)
			s.trySubmit()

		case hc := <-s.hcs:
			log.Infof("fast submit fastSubmitHandler::run chain head notify, revert [%v], advance [%v]", hc.revert, hc.advance) // to remove
			s.currentTs = hc.advance
			s.currentCtx = hc.ctx
			s.revert = hc.revert

			//s.currentDI = hc.di
			s.tryClear()
			s.trySubmit()

		case res := <-s.submitResNotify:
			s.processSubmitResult(res)
		}
	}
}

func (s *fastSubmitHandler) processSubmitResult(res *submitRes) {
	if res.err != nil {
		// Submit failed so inform the API and go back to the start state
		{
			// s.api.failPost(res.err, res.ts, res.di)
			log.Warnf("Aborted window post Submitting (Deadline: %+v at partition index: %d)", res.di, res.partIndex)
			// s.api.onAbort(res.ts, res.di)
		}
		res.submitState = SubmitStateStart
		return
	}
	// Submit succeeded so move to complete state
	res.submitState = SubmitStateComplete
}

func (s *fastSubmitHandler) trySubmit() {
	if s.currentTs == nil || len(s.posts) == 0 {
		return
	}

	for _, val := range s.posts {
		for partIndex, post := range val {
			if post == nil || post.params == nil || post.di == nil {
				continue
			}
			if s.currentTs.Height() < post.di.Open+SubmitConfidence || s.currentTs.Height() >= post.di.Close {
				continue
			}
			res := s.checkOrInitRes(*post.di, partIndex, post.ts)
			if res == nil || res.submitState != SubmitStateStart {
				continue
			}
			log.Infof("fast submit fastSubmitHandler::trySubmit, deadlines [%v], partition [%v]", res.di.Index, res.partIndex)
			res.submitState = SubmitStateSubmitting
			res.abort = s.api.startSubmitPoST(s.currentCtx, s.currentTs, post.di, []miner.SubmitWindowedPoStParams{*post.params}, func(err error) {
				res.err = err
				s.submitResNotify <- res
			})
		}
	}
}

func (s *fastSubmitHandler) addPost(post *postItem) {
	if post == nil {
		return
	}
	if _, ok := s.posts[post.di.Open]; !ok {
		s.posts[post.di.Open] = make(map[uint64]*postItem)
	}
	log.Infof("fast submit fastSubmitHandler::addPost, deadline [%v], partition [%v]", post.di.Index, post.partIndex)
	s.posts[post.di.Open][post.partIndex] = post
}

func (s *fastSubmitHandler) checkOrInitRes(di dline.Info, partIndex uint64, ts *types.TipSet) *submitRes {
	if _, ok := s.submitRes[di.Open]; !ok {
		s.submitRes[di.Open] = make(map[uint64]*submitRes)
	}
	if _, ok := s.submitRes[di.Open][partIndex]; !ok {
		s.submitRes[di.Open][partIndex] = &submitRes{
			submitState: SubmitStateStart,
			di:          &di,
			ts:          ts,
		}
	}
	return s.submitRes[di.Open][partIndex]
}

func (s *fastSubmitHandler) tryClear() {
	for open, val := range s.submitRes {
		for _, res := range val {
			revertedToPrevDL := s.revert != nil && s.revert.Height() < res.di.Open
			expired := s.currentTs.Height() >= res.di.Close
			if res.submitState == SubmitStateSubmitting && (revertedToPrevDL || expired) {
				if res.abort != nil {
					log.Warnf("fast submit fastSubmitHandler::tryClear SubmitStateSubmitting abort, deadlines [%v], partition [%v]", res.di.Index, res.partIndex)
					res.abort()
				}
			} else if res.submitState == SubmitStateComplete && revertedToPrevDL {
				log.Warnf("fast submit fastSubmitHandler::tryClear reset submitState, deadlines [%v], partition [%v]", res.di.Index, res.partIndex)
				res.submitState = SubmitStateStart
			}

			if expired {
				log.Infof("fast submit fastSubmitHandler::tryClear delete submitRes, deadlines [%v], partition [%v]", res.di.Index, res.partIndex)
				delete(s.submitRes, open)
			}
		}
	}
	for open, val := range s.posts {
		for _, res := range val {
			if s.currentTs.Height() >= res.di.Close {
				log.Infof("fast submit fastSubmitHandler::tryClear delete post, deadlines [%v], partition [%v]", res.di.Index, res.partIndex)
				delete(s.posts, open)
				break
			}
		}
	}
}
