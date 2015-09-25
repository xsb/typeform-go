package tfcom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetResponses(uid string) ([]byte, error) {

	params := "&completed=true"
	path := "/form/" + uid + "?key=" + ApiKey + params

	client := &http.Client{}
	req, err := newRequest("GET", path, nil)
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

	output := Output{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(output.Responses)

	return b, nil
}
