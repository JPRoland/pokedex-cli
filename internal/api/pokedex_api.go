package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jproland/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	fullURL := baseURL + "/location-area"

	if pageURL != nil {
		fullURL = *pageURL
	}

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaRes := LocationAreaResponse{}
		err := json.Unmarshal(cacheData, &locationAreaRes)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreaRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status: %v, %v", res.StatusCode, res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullURL, data)

	locationAreaRes := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaRes)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaRes, nil
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	fullURL := baseURL + "/location-area/" + locationAreaName

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(cacheData, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status: %v, %v", res.StatusCode, res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(cacheData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status: %v, %v", res.StatusCode, res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
