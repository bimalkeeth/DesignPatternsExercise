package Adapter

import (
	"github.com/fatih/color"
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "Hello world!"
	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnMsg := adapter.PrintStored()
	if returnMsg != "Legacy Printer:Adapter: Hello world!\n" {
		t.Errorf("Message didn't match: %s\n", returnMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnMsg = adapter.PrintStored()
	if returnMsg != "Hello world!" {
		t.Errorf("Message didn't match: %s\n", returnMsg)
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Println("Prints cyan text with an underline.")
	}
}
