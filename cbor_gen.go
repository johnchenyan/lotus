package sealing

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *SealTicket) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	// t.BlockHeight (uint64) (uint64)
	if len("BlockHeight") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BlockHeight\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("BlockHeight")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("BlockHeight")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.BlockHeight))); err != nil {
		return err
	}

	// t.TicketBytes ([]uint8) (slice)
	if len("TicketBytes") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TicketBytes\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("TicketBytes")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("TicketBytes")); err != nil {
		return err
	}

	if len(t.TicketBytes) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.TicketBytes was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.TicketBytes)))); err != nil {
		return err
	}
	if _, err := w.Write(t.TicketBytes); err != nil {
		return err
	}
	return nil
}

func (t *SealTicket) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SealTicket: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.BlockHeight (uint64) (uint64)
		case "BlockHeight":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.BlockHeight = uint64(extra)
			// t.TicketBytes ([]uint8) (slice)
		case "TicketBytes":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.TicketBytes: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.TicketBytes = make([]byte, extra)
			if _, err := io.ReadFull(br, t.TicketBytes); err != nil {
				return err
			}

		default:
		}
	}

	return nil
}
func (t *SealSeed) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	// t.BlockHeight (uint64) (uint64)
	if len("BlockHeight") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BlockHeight\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("BlockHeight")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("BlockHeight")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.BlockHeight))); err != nil {
		return err
	}

	// t.TicketBytes ([]uint8) (slice)
	if len("TicketBytes") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"TicketBytes\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("TicketBytes")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("TicketBytes")); err != nil {
		return err
	}

	if len(t.TicketBytes) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.TicketBytes was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.TicketBytes)))); err != nil {
		return err
	}
	if _, err := w.Write(t.TicketBytes); err != nil {
		return err
	}
	return nil
}

func (t *SealSeed) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SealSeed: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.BlockHeight (uint64) (uint64)
		case "BlockHeight":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.BlockHeight = uint64(extra)
			// t.TicketBytes ([]uint8) (slice)
		case "TicketBytes":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.TicketBytes: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.TicketBytes = make([]byte, extra)
			if _, err := io.ReadFull(br, t.TicketBytes); err != nil {
				return err
			}

		default:
		}
	}

	return nil
}
func (t *Piece) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{163}); err != nil {
		return err
	}

	// t.DealID (uint64) (uint64)
	if len("DealID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealID\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("DealID")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("DealID")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.DealID))); err != nil {
		return err
	}

	// t.Size (uint64) (uint64)
	if len("Size") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Size\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Size")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Size")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Size))); err != nil {
		return err
	}

	// t.CommP ([]uint8) (slice)
	if len("CommP") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommP\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("CommP")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("CommP")); err != nil {
		return err
	}

	if len(t.CommP) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.CommP was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.CommP)))); err != nil {
		return err
	}
	if _, err := w.Write(t.CommP); err != nil {
		return err
	}
	return nil
}

func (t *Piece) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Piece: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.DealID (uint64) (uint64)
		case "DealID":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.DealID = uint64(extra)
			// t.Size (uint64) (uint64)
		case "Size":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.Size = uint64(extra)
			// t.CommP ([]uint8) (slice)
		case "CommP":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.CommP: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.CommP = make([]byte, extra)
			if _, err := io.ReadFull(br, t.CommP); err != nil {
				return err
			}

		default:
		}
	}

	return nil
}
func (t *SectorInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{173}); err != nil {
		return err
	}

	// t.State (uint64) (uint64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("State")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("State")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.State))); err != nil {
		return err
	}

	// t.SectorID (uint64) (uint64)
	if len("SectorID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"SectorID\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("SectorID")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("SectorID")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorID))); err != nil {
		return err
	}

	// t.Nonce (uint64) (uint64)
	if len("Nonce") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Nonce\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Nonce")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Nonce")); err != nil {
		return err
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Nonce))); err != nil {
		return err
	}

	// t.Pieces ([]sealing.Piece) (slice)
	if len("Pieces") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Pieces\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Pieces")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Pieces")); err != nil {
		return err
	}

	if len(t.Pieces) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Pieces was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Pieces)))); err != nil {
		return err
	}
	for _, v := range t.Pieces {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.CommD ([]uint8) (slice)
	if len("CommD") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommD\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("CommD")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("CommD")); err != nil {
		return err
	}

	if len(t.CommD) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.CommD was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.CommD)))); err != nil {
		return err
	}
	if _, err := w.Write(t.CommD); err != nil {
		return err
	}

	// t.CommR ([]uint8) (slice)
	if len("CommR") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommR\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("CommR")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("CommR")); err != nil {
		return err
	}

	if len(t.CommR) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.CommR was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.CommR)))); err != nil {
		return err
	}
	if _, err := w.Write(t.CommR); err != nil {
		return err
	}

	// t.Proof ([]uint8) (slice)
	if len("Proof") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proof\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Proof")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Proof")); err != nil {
		return err
	}

	if len(t.Proof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Proof)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Proof); err != nil {
		return err
	}

	// t.Ticket (sealing.SealTicket) (struct)
	if len("Ticket") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Ticket\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Ticket")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Ticket")); err != nil {
		return err
	}

	if err := t.Ticket.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PreCommitMessage (cid.Cid) (struct)
	if len("PreCommitMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PreCommitMessage\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("PreCommitMessage")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("PreCommitMessage")); err != nil {
		return err
	}

	if t.PreCommitMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.PreCommitMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PreCommitMessage: %w", err)
		}
	}

	// t.Seed (sealing.SealSeed) (struct)
	if len("Seed") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Seed\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Seed")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Seed")); err != nil {
		return err
	}

	if err := t.Seed.MarshalCBOR(w); err != nil {
		return err
	}

	// t.CommitMessage (cid.Cid) (struct)
	if len("CommitMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"CommitMessage\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("CommitMessage")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("CommitMessage")); err != nil {
		return err
	}

	if t.CommitMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.CommitMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommitMessage: %w", err)
		}
	}

	// t.FaultReportMsg (cid.Cid) (struct)
	if len("FaultReportMsg") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FaultReportMsg\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("FaultReportMsg")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("FaultReportMsg")); err != nil {
		return err
	}

	if t.FaultReportMsg == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.FaultReportMsg); err != nil {
			return xerrors.Errorf("failed to write cid field t.FaultReportMsg: %w", err)
		}
	}

	// t.LastErr (string) (string)
	if len("LastErr") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"LastErr\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("LastErr")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("LastErr")); err != nil {
		return err
	}

	if len(t.LastErr) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.LastErr was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.LastErr)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.LastErr)); err != nil {
		return err
	}
	return nil
}

func (t *SectorInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SectorInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(br)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.State (uint64) (uint64)
		case "State":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.State = uint64(extra)
			// t.SectorID (uint64) (uint64)
		case "SectorID":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.SectorID = uint64(extra)
			// t.Nonce (uint64) (uint64)
		case "Nonce":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint64 field")
			}
			t.Nonce = uint64(extra)
			// t.Pieces ([]sealing.Piece) (slice)
		case "Pieces":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.Pieces: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}
			if extra > 0 {
				t.Pieces = make([]Piece, extra)
			}
			for i := 0; i < int(extra); i++ {

				var v Piece
				if err := v.UnmarshalCBOR(br); err != nil {
					return err
				}

				t.Pieces[i] = v
			}

			// t.CommD ([]uint8) (slice)
		case "CommD":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.CommD: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.CommD = make([]byte, extra)
			if _, err := io.ReadFull(br, t.CommD); err != nil {
				return err
			}
			// t.CommR ([]uint8) (slice)
		case "CommR":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.CommR: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.CommR = make([]byte, extra)
			if _, err := io.ReadFull(br, t.CommR); err != nil {
				return err
			}
			// t.Proof ([]uint8) (slice)
		case "Proof":

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.Proof = make([]byte, extra)
			if _, err := io.ReadFull(br, t.Proof); err != nil {
				return err
			}
			// t.Ticket (sealing.SealTicket) (struct)
		case "Ticket":

			{

				if err := t.Ticket.UnmarshalCBOR(br); err != nil {
					return err
				}

			}
			// t.PreCommitMessage (cid.Cid) (struct)
		case "PreCommitMessage":

			{

				pb, err := br.PeekByte()
				if err != nil {
					return err
				}
				if pb == cbg.CborNull[0] {
					var nbuf [1]byte
					if _, err := br.Read(nbuf[:]); err != nil {
						return err
					}
				} else {

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PreCommitMessage: %w", err)
					}

					t.PreCommitMessage = &c
				}

			}
			// t.Seed (sealing.SealSeed) (struct)
		case "Seed":

			{

				if err := t.Seed.UnmarshalCBOR(br); err != nil {
					return err
				}

			}
			// t.CommitMessage (cid.Cid) (struct)
		case "CommitMessage":

			{

				pb, err := br.PeekByte()
				if err != nil {
					return err
				}
				if pb == cbg.CborNull[0] {
					var nbuf [1]byte
					if _, err := br.Read(nbuf[:]); err != nil {
						return err
					}
				} else {

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommitMessage: %w", err)
					}

					t.CommitMessage = &c
				}

			}
			// t.FaultReportMsg (cid.Cid) (struct)
		case "FaultReportMsg":

			{

				pb, err := br.PeekByte()
				if err != nil {
					return err
				}
				if pb == cbg.CborNull[0] {
					var nbuf [1]byte
					if _, err := br.Read(nbuf[:]); err != nil {
						return err
					}
				} else {

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.FaultReportMsg: %w", err)
					}

					t.FaultReportMsg = &c
				}

			}
			// t.LastErr (string) (string)
		case "LastErr":

			{
				sval, err := cbg.ReadString(br)
				if err != nil {
					return err
				}

				t.LastErr = string(sval)
			}

		default:
		}
	}

	return nil
}
