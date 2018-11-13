package escpos

func (e *Escpos) LineSpacing(spacing byte) {
	e.dev.Write([]byte{esc, 0x33, spacing})
}

func (e *Escpos) CharacterSpacing(spacing byte) {
	e.dev.Write([]byte{esc, 0x20, spacing})
}

func (e *Escpos) Margin(high, low byte) {
	e.dev.Write([]byte{gs, 0x4C, low, high})
}
