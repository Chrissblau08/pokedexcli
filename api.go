package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/Chrissblau08/pokedexcli.git/internal/pokecache"
)

type LocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonCatch struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type LocationByNameResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

var CacheInstance = pokecache.NewCache(10 * time.Second)
var BASE_URL string = "https://pokeapi.co/api/v2/"

func getLocationsArea(url string) ([]string, string, string, error) {
	if url == "" {
		url = BASE_URL + "location-area?limit=20" // Default: erste 20
	}

	if data, ok := CacheInstance.Get(url); ok {
		var cached LocationsResponse
		if err := json.Unmarshal(data, &cached); err == nil {
			names := make([]string, len(cached.Results))
			for i, loc := range cached.Results {
				names[i] = loc.Name
			}
			return names, cached.Next, cached.Previous, nil
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", "", errors.New("request failed: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", "", err
	}

	CacheInstance.Add(url, body)

	var data LocationsResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, "", "", err
	}

	names := make([]string, len(data.Results))
	for i, loc := range data.Results {
		names[i] = loc.Name
	}

	return names, data.Next, data.Previous, nil
}

func getPokemonEncountersbyLocationArea(areaName string) ([]string, error) {
	FULL_URL := BASE_URL + "location-area/" + areaName

	if data, ok := CacheInstance.Get(FULL_URL); ok {
		var cached LocationsResponse
		if err := json.Unmarshal(data, &cached); err == nil {
			names := make([]string, len(cached.Results))
			for i, pok := range cached.Results {
				names[i] = pok.Name
			}
			return names, nil
		}
	}

	resp, err := http.Get(FULL_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	CacheInstance.Add(FULL_URL, body)

	var data LocationByNameResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	names := make([]string, len(data.PokemonEncounters))
	for i, pokemon := range data.PokemonEncounters {
		names[i] = pokemon.Pokemon.Name
	}

	return names, nil
}

func getPokemonByNameInfo(pokemon string) (PokemonCatch, error) {
	FULL_URL := BASE_URL + "pokemon/" + pokemon

	resp, err := http.Get(FULL_URL)
	if err != nil {
		return PokemonCatch{}, err
	}
	defer resp.Body.Close()

	if data, ok := CacheInstance.Get(FULL_URL); ok {
		var cached PokemonCatch
		if err := json.Unmarshal(data, &cached); err == nil {
			return cached, nil
		}
	}

	if resp.StatusCode != http.StatusOK {
		return PokemonCatch{}, errors.New("request failed: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonCatch{}, err
	}

	CacheInstance.Add(FULL_URL, body)

	var data PokemonCatch
	if err := json.Unmarshal(body, &data); err != nil {
		return PokemonCatch{}, err
	}

	return data, nil
}
