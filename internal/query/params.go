package query

import (
	"fmt"
	"slices"
)

var validTypes = []string{"paras", "words", "bytes"}
var apiLimits = map[string]uint32{
	"paras": 150,
	"bytes": 100000,
	"words": 10000,
}

// Parameters that the https://lispum.com API accepts.
type Params struct {
	Amount uint32
	What   string
}

// Validates a set of parameters to ensure they fit the ones expected by the https://lispum.com API.
func ValidateParams(params *Params) error {
	if !slices.Contains(validTypes, params.What) {
		return fmt.Errorf("unsupported what param '%v'", params.What)
	}

	if params.Amount < 1 {
		return fmt.Errorf("invalid amount '%v' < 1", params.Amount)
	} else if params.Amount > apiLimits[params.What] {
		return fmt.Errorf("invalid amount '%v' > %v", params.Amount, apiLimits[params.What])
	}

	return nil
}
