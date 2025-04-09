package atom

import "encoding/binary"

// StcoBox - Chunk offset atom Box
// Box Type: stco
// Container: Chunk offset atom Box (stco)
// Mandatory: Yes
// Quantity: Exactly one.
type StcoBox struct {
	*Box
	Version          byte
	Flags            uint32
	EntryCount       uint32
	ChunkOffsetTable []uint32
}

func (b *StcoBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.EntryCount = binary.BigEndian.Uint32(data[4:8])

	for i := 0; i < int(b.EntryCount); i++ {
		b.ChunkOffsetTable = append(b.ChunkOffsetTable, binary.BigEndian.Uint32(data[i*4+8:i*4+12]))
	}
	return nil
}
