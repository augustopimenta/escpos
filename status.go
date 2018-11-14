package escpos

type status byte

const (
	StatusPaperAdequate status = 16
	StatusPaperNearEnd  status = 19
	StatusPaperEnd      status = 28
)

func (e *Escpos) GetStatus() (status, error) {
	s := make([]byte, 1)

	e.dev.Write([]byte{gs, 0x72, 1})
	_, err := e.dev.Read(s)

	if err != nil {
		return 0, err
	}

	return status(s[0]), nil
}
