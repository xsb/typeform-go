package tfio

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (fd FormDescription) CreateForm() error {

	b, err := json.Marshal(fd)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := newRequest("POST", "/forms", bytes.NewBuffer(b))
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

	//Debug
	fmt.Println(string(body))

	return nil
}
