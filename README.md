# escpos
Golang package for handling ESC-POS thermal printer commands

## Instalation

```bash
go get -u github.com/augustopimenta/escpos
```

## Usage example

```golang
package main

import (
	"github.com/augustopimenta/escpos"
	"os"
)

func main() {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	p := escpos.New(f)
	p.Init()

	p.FontSize(2, 2)
	p.Font(escpos.FontB)
	p.FontAlign(escpos.AlignCenter)
	p.Writeln("Hello World!")
	p.Feed()

	p.FontSize(1, 1)
	p.Font(escpos.FontA)
	p.FontAlign(escpos.AlignLeft)
	p.Writeln("Lorem ipsum primis potenti in purus vestibulum amet enim, fames orci dapibus tempor...")

	p.FeedN(10)

	p.FullCut()
}
```
