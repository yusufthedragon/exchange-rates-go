package exchangerates

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SendRequest sends request to API and process the response.
func SendRequest(endpoint string) (ResponseParameter, error) {
	var req, err = http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return ResponseParameter{}, err
	}

	var client = &http.Client{}
	var resp, errRequest = client.Do(req)

	if errRequest != nil {
		return ResponseParameter{}, errRequest
	}

	defer resp.Body.Close()

	// Read the response body
	var body, errResponse = ioutil.ReadAll(resp.Body)

	if errResponse != nil {
		return ResponseParameter{}, errResponse
	}

	// Decode the response JSON into target struct
	var response ResponseParameter

	json.Unmarshal(body, &response)

	return response, nil
}
