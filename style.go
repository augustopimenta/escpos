package escpos

type fontfamily byte

const (
	FontA fontfamily = 0
	FontB fontfamily = 1
)

func (e *Escpos) Font(f fontfamily) {
	font := 0

	switch f {
	case FontA:
		font = 0
	case FontB:
		font = 1
	}

	e.dev.Write([]byte{esc, 0x4D, byte(font)})
}

func (e *Escpos) FontSize() {

}

func (e *Escpos) FontUnderline() {

}

func (e *Escpos) FontBold() {

}
