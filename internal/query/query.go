package query

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The details contained in a response from the https://lipsum.com API.
type LipsumInfo struct {
	Lipsum     string
	Generated  string
	DonateLink string
	CreditLink string
	CreditName string
}

// A response from the https://lipsum.com API.
type LipsumReturn struct {
	Feed LipsumInfo
}

// A function which takes in a set of parameters and builds a request URL with them.
type URLBuilder func(params *Params) string

// Builds a URL for querying https://lipsum.com.
func BuildURL(params *Params) string {
	return fmt.Sprintf("https://lipsum.com/feed/json?what=%v&amount=%v&start=yes", params.What, params.Amount)
}

// Makes a request to the https://lipsum.com API to retrieve lorem ipsum text
func Execute(url string) (*LipsumInfo, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return &LipsumInfo{}, err
	}

	req.Header.Add("User-Agent", "go-lipsum")

	res, err := client.Do(req)

	if err != nil {
		return &LipsumInfo{}, err
	} else if res.StatusCode != http.StatusOK {
		return &LipsumInfo{}, fmt.Errorf("request returned %v", res.StatusCode)
	}

	unmarshalled := LipsumReturn{}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&unmarshalled)
	res.Body.Close()

	if err != nil {
		return &LipsumInfo{}, err
	}

	return &unmarshalled.Feed, nil
}
