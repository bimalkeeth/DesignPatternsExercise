package Bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterApi interface {
	PrintMessage(string) error
}

type PrinterIml1 struct{}

func (p *PrinterIml1) PrintMessage(msg string) error {

	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("content received on Writer was empty")
	return
}
