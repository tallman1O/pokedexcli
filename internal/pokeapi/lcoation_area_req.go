package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationAreaResponse{}, fmt.Errorf("error: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	LocationAreaResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &LocationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return LocationAreaResp, nil
}
