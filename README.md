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

	ec := escpos.New(f)
	ec.Init()
	ec.Font(escpos.FontB)
	ec.FontAlign(escpos.AlignCenter)
	ec.Write("Hello World!")
	ec.Feed()
	ec.FullCut()
}
```
