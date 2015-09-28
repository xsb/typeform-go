package tfio

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Form struct {
	Id         string
	Title      string
	Self       string
	FormRender string
}

type FormDescription struct {
	Title  string        `json:"title"`
	Fields []interface{} `json:"fields"`
}

func NewForm(title string) FormDescription {
	f := FormDescription{}
	f.Title = title
	return f
}

func (fd FormDescription) CreateForm() ([]byte, error) {

	b, err := json.Marshal(fd)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := newRequest("POST", "/forms", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
