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

func (e *Escpos) Write(s string) {
	str, err := e.enc.String(s)

	if err != nil {
		panic(err)
	}

	e.dev.Write([]byte(str))
}

func New(dev io.ReadWriter) *Escpos {
	escpos := &Escpos{dev: dev}
	escpos.Charset(CharsetPC437)

	return escpos
}
