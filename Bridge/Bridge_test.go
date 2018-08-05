package Bridge

import (
	"strings"
	"testing"
)

func TestPrintApi1(t *testing.T) {
	api1 := PrinterIml1{}
	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n",
			err.Error())
	}

	api2 := PrinterImpl2{}
	err = api2.PrintMessage("Hello")
	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\nActual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}

	api2 = PrinterImpl2{Writer: &testWriter}
	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n",
			err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \n%s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}

}
