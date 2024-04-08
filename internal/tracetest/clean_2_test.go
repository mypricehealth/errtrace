package tracetest

import "github.com/mypricehealth/errtrace"

// Separate file to verify how Clean handles separate files.

func f1() error {
	return errtrace.Wrap(f2())
}

func f2() error {
	if err := f3(); err != nil {
		return errtrace.Errorf("f3: %w", err)
	}

	return nil
}

func f3() error {
	return errtrace.New("err")
}
