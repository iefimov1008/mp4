package atom

// StblBox - Sample Table Box
// Box Type: stbl
// Container: Media Information Box (minf)
// Mandatory: Yes
// Quantity: Exactly one.
type StblBox struct {
	*Box
	Stts *SttsBox
	Stsd *StsdBox
	Stss *StssBox
	Ctts *CttsBox
	Stsc *StscBox
}

func (b *StblBox) parse() error {
	boxes := readBoxes(b.Reader, b.Start+BoxHeaderSize, b.Size-BoxHeaderSize)

	for _, box := range boxes {
		switch box.Name {
		case "stts":
			b.Stts = &SttsBox{Box: box}
			b.Stts.parse()

		case "stsd":
			b.Stsd = &StsdBox{Box: box}
			b.Stsd.parse()

		case "stss":
			b.Stss = &StssBox{Box: box}
			b.Stss.parse()
		case "ctts":
			b.Ctts = &CttsBox{Box: box}
			b.Ctts.parse()
		case "stsc":
			b.Stsc = &StscBox{Box: box}
			b.Stsc.parse()
		}
	}
	return nil
}
