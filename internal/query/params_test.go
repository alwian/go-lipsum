package query_test

import (
	"testing"

	"github.com/alwian/go-lipsum/internal/query"
)

// validateParams

func TestValidateParamsWithValidParams(t *testing.T) {
	input := &query.Params{
		What:   "words",
		Amount: 5,
	}

	err := query.ValidateParams(input)

	if err != nil {
		t.Fail()
	}
}

func TestValidateParamsWithAmountTooLow(t *testing.T) {
	input := &query.Params{
		What:   "words",
		Amount: 0,
	}

	err := query.ValidateParams(input)

	if err == nil || err.Error() != "invalid amount '0' < 1" {
		t.Fail()
	}
}

func TestValidateParamsWithAmountTooHigh(t *testing.T) {
	input := &query.Params{
		What:   "words",
		Amount: 1000000,
	}

	err := query.ValidateParams(input)

	if err == nil || err.Error() != "invalid amount '1000000' > 10000" {
		t.Fail()
	}
}

func TestValidateParamsWithInvalidWhat(t *testing.T) {
	input := &query.Params{
		What:   "things",
		Amount: 5,
	}

	err := query.ValidateParams(input)

	if err == nil || err.Error() != "unsupported what param 'things'" {
		t.Fail()
	}
}
