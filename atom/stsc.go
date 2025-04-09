package atom

import "encoding/binary"

// StscBox - Sample-to-chunk atom Box
// Box Type: stsc
// Container: Sample-to-chunk atom Box (stsc)
// Mandatory: Yes
// Quantity: Exactly one.
type StscBox struct {
	*Box
	Version           byte
	Flags             uint32
	EntryCount        uint32
	SampletoChunkData SampletoChunk
}
type SampletoChunk struct {
	Firstchunk          uint32
	SamplesPerChunk     uint32
	SampleDescriptionID uint32
}

func (b *StscBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.EntryCount = binary.BigEndian.Uint32(data[4:8])
	b.SampletoChunkData.Firstchunk = binary.BigEndian.Uint32(data[8:12])
	b.SampletoChunkData.SamplesPerChunk = binary.BigEndian.Uint32(data[12:16])
	b.SampletoChunkData.SampleDescriptionID = binary.BigEndian.Uint32(data[16:20])
	return nil
}
