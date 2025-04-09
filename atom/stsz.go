package atom

import "encoding/binary"

// StszBox - Sample size atom Box
// Box Type: sttz
// Container: Sample Size Table Box (stsz)
// Mandatory: Yes
// Quantity: Exactly one.
type StszBox struct {
	*Box
	Version         byte
	Flags           uint32
	SampleSize      uint32
	EntryCount      uint32
	SampleSizeTable []uint32
}

func (b *StszBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.SampleSize = binary.BigEndian.Uint32(data[4:8])
	b.EntryCount = binary.BigEndian.Uint32(data[8:12])
	if b.SampleSize == 0 {
		for i := 0; i < int(b.EntryCount); i++ {
			b.SampleSizeTable = append(b.SampleSizeTable, binary.BigEndian.Uint32(data[i*4+12:i*4+16]))
		}
	}
	return nil
}
