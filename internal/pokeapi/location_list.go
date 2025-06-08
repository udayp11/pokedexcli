package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (location, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		locations := location{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return location{}, err
		}

		return locations, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return location{}, err
	}

	locations := location{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return location{}, err
	}

	c.cache.Add(url, dat)
	return locations, nil
}
