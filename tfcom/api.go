package tfcom

import (
	"io"
	"net/http"
)

const version = "v0"

var (
	ApiKey    string
	ApiUrl    = "https://api.typeform.com/" + version
	UserAgent = "xsb/typeform-go"
)

func newRequest(verb, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(verb, ApiUrl+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-TOKEN", ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}
