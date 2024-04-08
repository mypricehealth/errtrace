//go:build ignore

package foo

import (
	"strconv"

	_ "github.com/mypricehealth/errtrace"
)

func Unwrapped(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i + 42, nil
}
