package cli_base

import (
	"strconv"
	"strings"
)

func StrToFloat(s string) (float64, error) {
	vFloat64, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, err
	}

	return vFloat64, nil
}
