// Package lipsum provides methods for generating lorem ipsum text via the site https://lipsum.com.
package lipsum

import (
	"github.com/alwian/go-lipsum/internal/query"
)

func generate(params *query.Params) (string, error) {
	err := query.ValidateParams(params)

	if err != nil {
		return "", err
	}

	res, err := query.Execute(query.BuildURL(params))

	if err != nil {
		return "", err
	}

	return res.Lipsum, nil
}

// Retrieves the requested number of lorem ipsum bytes.
//
// The amount requested can range from 1-100000, anything outside of this range will result in an error.
func Bytes(count uint32) (string, error) {
	res, err := generate(&query.Params{
		Amount: count,
		What:   "bytes",
	})

	return res, err
}

// Retrieves the requested number of lorem ipsum words.
//
// The amount requested can range from 1-10000, anything outside of this range will result in an error.
func Words(count uint16) (string, error) {
	res, err := generate(&query.Params{
		Amount: uint32(count),
		What:   "words",
	})

	return res, err
}

// Retrieves the requested number of lorem ipsum paragraphs.
//
// The amount requested can range from 1-150, anything outside of this range will result in an error.
func Paragraphs(count uint8) (string, error) {
	res, err := generate(&query.Params{
		Amount: uint32(count),
		What:   "paras",
	})
	return res, err
}
