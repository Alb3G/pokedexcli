package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationArea, error) {
	url := BASE_URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	// Check if cache contains any data before making the request
	if data, ok := c.cache.Get(url); ok {
		var locationArea LocationArea
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	// Make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	// Add response to Cache
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, nil
	}

	c.cache.Add(url, data)
	return locationArea, nil
}
