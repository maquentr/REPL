package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	//check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit !!!!")
		LocationAreasRespp := LocationAreasResp{}
		err := json.Unmarshal(dat, &LocationAreasRespp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return LocationAreasRespp, nil
	}
	fmt.Println("cache miss !!!!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	LocationAreasRespp := LocationAreasResp{}
	err = json.Unmarshal(dat, &LocationAreasRespp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return LocationAreasRespp, nil

}
func (c *Client) GetLocationAreas(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache hit !!!!")
		LocationAreaa := LocationArea{}
		err := json.Unmarshal(dat, &LocationAreaa)
		if err != nil {
			return LocationArea{}, err
		}
		return LocationAreaa, nil
	}
	fmt.Println("cache miss !!!!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	LocationAreaa := LocationArea{}
	err = json.Unmarshal(dat, &LocationAreaa)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, dat)

	return LocationAreaa, nil

}
