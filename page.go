package escpos

func (e *Escpos) LineSpacing(spacing byte) {
	e.dev.Write([]byte{esc, 0x33, spacing})
}

func (e *Escpos) CharacterSpacing(spacing byte) {
	e.dev.Write([]byte{esc, 0x20, spacing})
}

func (e *Escpos) Margin(n uint16) {
	e.dev.Write([]byte{gs, 0x4C, byte(n & 0xff), byte(n >> 8)})
}
