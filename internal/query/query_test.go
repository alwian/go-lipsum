package query_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alwian/go-lipsum/internal/query"
)

// BuildURL

func TestBuildURLReturnsCorrectURL(t *testing.T) {
	var input = &query.Params{
		Amount: 5,
		What:   "words",
	}

	var output = query.BuildURL(input)

	if output != fmt.Sprintf("https://lipsum.com/feed/json?what=%v&amount=%v&start=yes", input.What, input.Amount) {
		t.Fail()
	}
}

// Execute

func TestExecuteNewRequestError(t *testing.T) {
	_, err := query.Execute("invalidscheme://abcd")

	if err == nil || err.Error() != "Get \"invalidscheme://abcd\": unsupported protocol scheme \"invalidscheme\"" {
		t.Fail()
	}
}

func TestExecuteRequestError(t *testing.T) {
	_, err := query.Execute("https://")

	if err == nil || err.Error() != "Get \"https:\": http: no Host in request URL" {
		t.Fail()
	}
}

func TestExecuteGetRequestBadResponseCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	_, err := query.Execute(ts.URL)

	if err == nil || err.Error() != "request returned 400" {
		t.Fail()
	}
}

func TestExecuteInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()

	_, err := query.Execute(ts.URL)

	if err == nil || err.Error() != "EOF" {
		t.Fail()
	}
}

func TestExecuteSuccessfulResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.UserAgent() == "go-lipsum" {
			res, _ := json.Marshal(query.LipsumReturn{
				Feed: query.LipsumInfo{
					Lipsum:     "test",
					Generated:  "test",
					DonateLink: "test",
					CreditLink: "test",
					CreditName: "test",
				},
			})

			w.Write(res)
		}
	}))
	defer ts.Close()

	res, err := query.Execute(ts.URL)

	if err != nil {
		t.Fail()
	}

	if res.Lipsum != "test" || res.Generated != "test" || res.DonateLink != "test" || res.CreditLink != "test" || res.CreditName != "test" {
		t.Fail()
	}
}
