package api

import (
	"net/http"
)

func GetLocationAreas(url string) (string, error) {
	// Make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("", "")

	return "", nil
}
