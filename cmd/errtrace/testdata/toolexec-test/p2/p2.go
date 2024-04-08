package p2

import (
	"github.com/mypricehealth/errtrace"

	"github.com/mypricehealth/errtrace/cmd/errtrace/testdata/toolexec-test/p3"
)

// CallP3 calls p3, and wraps the error.
func CallP3() error {
	return errtrace.Wrap(p3.ReturnErr()) // @trace
}
