package errtrace_test

import (
	"fmt"

	"github.com/mypricehealth/errtrace"
	"github.com/mypricehealth/errtrace/internal/tracetest"
)

func f1() error {
	return errtrace.Wrap(f2())
}

func f2() error {
	return errtrace.Wrap(f3())
}

func f3() error {
	return errtrace.New("failed")
}

func Example_trace() {
	got := errtrace.FormatString(f1())

	// make trace agnostic to environment-specific location
	// and less sensitive to line number changes.
	fmt.Println(tracetest.MustClean(got))

	// Output:
	//failed
	//
	//github.com/mypricehealth/errtrace_test.f3
	//	/path/to/errtrace/example_trace_test.go:3
	//github.com/mypricehealth/errtrace_test.f2
	//	/path/to/errtrace/example_trace_test.go:2
	//github.com/mypricehealth/errtrace_test.f1
	//	/path/to/errtrace/example_trace_test.go:1
}
