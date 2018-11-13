package escpos

import (
	"io"

	"golang.org/x/text/encoding"
)

const (
	esc = 0x1B
	gs  = 0x1D
	lf  = 0x0A
)

type Escpos struct {
	dev io.ReadWriter
	enc *encoding.Encoder
}

// Init printer cleaning your old configs
func (e *Escpos) Init() {
	e.dev.Write([]byte{esc, 0x40})
}

// Feed skip one line of paper
func (e *Escpos) Feed() {
	e.dev.Write([]byte{lf})
}

// Feed skip n lines of paper
func (e *Escpos) FeedN(n byte) {
	e.dev.Write([]byte{esc, 0x64, n})
}

// SelfTest start self test of printer
func (e *Escpos) SelfTest() {
	e.dev.Write([]byte{gs, 0x28, 0x41, 0x02, 0x00, 0x00, 0x02})
}

// Write print text
func (e *Escpos) Write(text string) {
	str, err := e.enc.String(text)

	if err != nil {
		panic(err)
	}

	e.dev.Write([]byte(str))
}

// New create new Escpos struct and set default enconding
func New(dev io.ReadWriter) *Escpos {
	escpos := &Escpos{dev: dev}
	escpos.Charset(CharsetPC437)

	return escpos
}
