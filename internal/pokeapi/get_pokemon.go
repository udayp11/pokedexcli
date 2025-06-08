package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(area string) (GetPokemon, error) {

	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		locations := GetPokemon{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return GetPokemon{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GetPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return GetPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetPokemon{}, err
	}

	locations := GetPokemon{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return GetPokemon{}, err
	}

	c.cache.Add(url, dat)
	return locations, nil
}
