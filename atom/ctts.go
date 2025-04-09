package atom

import "encoding/binary"

// CttsBox - Composition offset atom Box
// Box Type: ctts
// Container: Composition offset atom Box (ctts)
// Mandatory: Yes
// Quantity: Exactly one.
type CttsBox struct {
	*Box
	Version    byte
	Flags      uint32
	EntryCount uint32
}

func (b *CttsBox) parse() error {
	data := b.ReadBoxData()

	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.EntryCount = binary.BigEndian.Uint32(data[4:8])
	return nil
}
