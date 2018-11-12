package escpos

import "fmt"

type fontfamily byte

const (
	FontA fontfamily = 0
	FontB fontfamily = 1
)

func (e *Escpos) Font(f fontfamily) {
	e.dev.Write([]byte{esc, 0x4D, byte(f)})
}

type fontalign byte

const (
	AlignLeft   fontalign = 0
	AlignCenter fontalign = 1
	AlignRight  fontalign = 2
)

func (e *Escpos) FontAlign(a fontalign) {
	e.dev.Write([]byte{esc, 0x61, byte(a)})
}

func (e *Escpos) FontSize(width, height uint8) {
	if width >= 1 && width <= 8 && height >= 1 && height <= 8 {
		e.dev.Write([]byte{esc, 0x21, ((width - 1) << 4) | (height - 1)})
	} else {
		panic(fmt.Sprintf("Wrong font size: (%d x %d)", width, height))
	}
}

func (e *Escpos) FontUnderline(on bool) {
	e.dev.Write([]byte{esc, 0x2D, boolToByte(on)})
}

func (e *Escpos) FontBold(on bool) {
	e.dev.Write([]byte{esc, 0x45, boolToByte(on)})
}
