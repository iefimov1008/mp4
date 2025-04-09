package atom

import (
	"encoding/binary"
)

// StssBox - Sync sample atom Box
// Box Type: stss
// Container: Sync sample atom Box (stss)
// Mandatory: Yes
// Quantity: Exactly one.
type StssBox struct {
	*Box
	Version         byte
	Flags           uint32
	EntriesNumber   uint32
	SyncSampleTable []uint32
}

func (b *StssBox) parse() error {
	data := b.ReadBoxData()
	b.Version = data[0]
	b.Flags = binary.BigEndian.Uint32(data[0:4])
	b.EntriesNumber = binary.BigEndian.Uint32(data[4:8])
	for i := 8; i < len(data); i += 4 {
		b.SyncSampleTable = append(b.SyncSampleTable, uint32(binary.BigEndian.Uint32(data[i:i+4])))
	}

	return nil
}
