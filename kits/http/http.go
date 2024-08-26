package httputils

import (
	"io/ioutil"
	"net/http"
)

// Get sends a GET request to the specified URL and returns the response as a string.
func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
