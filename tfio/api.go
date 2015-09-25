package tfio

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const version = "v0.4"

var (
	ApiToken  string
	ApiUrl    = "https://api.typeform.io/" + version
	UserAgent = "xsb/typeform-go"
)

type ApiMetadata struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Version       string `json:"version"`
	Documentation string `json:"documentation"`
	Support       string `json:"support"`
	Time          string `json:"time"`
}

func newRequest(verb, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(verb, ApiUrl+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-TOKEN", ApiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", UserAgent)
	return req, nil

}

func PrintApiMetadata() error {

	client := &http.Client{}
	req, err := newRequest("GET", "/", nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	metadata := ApiMetadata{}
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		return err
	}

	fmt.Println("Name:", metadata.Name)
	fmt.Println("Description:", metadata.Description)
	fmt.Println("Version:", metadata.Version)
	fmt.Println("Documentation:", metadata.Documentation)
	fmt.Println("Support:", metadata.Support)
	fmt.Println("Time:", metadata.Time)

	return nil
}
