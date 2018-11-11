package escpos

import "golang.org/x/text/encoding/charmap"

type charset byte

const (
	CharsetPC437 charset = 0
	CharsetPC850 charset = 2
	CharsetPC860 charset = 3
	CharsetPC863 charset = 4
	CharsetPC865 charset = 5
	CharsetPC858 charset = 19
)

func (e *Escpos) Charset(c charset) {
	switch c {
	case CharsetPC437:
		e.enc = charmap.CodePage437.NewEncoder()
	case CharsetPC850:
		e.enc = charmap.CodePage850.NewEncoder()
	case CharsetPC858:
		e.enc = charmap.CodePage858.NewEncoder()
	case CharsetPC860:
		e.enc = charmap.CodePage860.NewEncoder()
	case CharsetPC863:
		e.enc = charmap.CodePage863.NewEncoder()
	case CharsetPC865:
		e.enc = charmap.CodePage865.NewEncoder()
	}

	e.dev.Write([]byte{esc, 0x74, byte(c)})
}
