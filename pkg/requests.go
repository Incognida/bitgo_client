package pkg

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func isJSON(payload string) (bool, error) {
	var js json.RawMessage
	if err := json.Unmarshal([]byte(payload), &js); err != nil {
		return false, err
	}
	return true, nil
}

func query(method, url, token string, body io.Reader) (string, error) {
	bearer := "Bearer " + token
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	return string(res), nil
}
