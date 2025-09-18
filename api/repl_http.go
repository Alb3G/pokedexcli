package api

import (
	"encoding/json"
	"net/http"

	"github.com/Alb3G/pokedexcli/types"
)

func GetLocationAreas(url string) (types.LocationArea, error) {
	// Make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.LocationArea{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return types.LocationArea{}, err
	}
	defer res.Body.Close()

	var locationArea types.LocationArea
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&locationArea)

	return locationArea, nil
}
