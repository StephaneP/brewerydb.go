package brewerydb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const API_ENDPOINT = "http://api.brewerydb.com"
const API_VERSION = "v2"

var (
	ErrBadRequest       = errors.New("400: Bad Request")
	ErrInvalidSignature = errors.New("401: Invalid Signature")
)

type (
	Beers []Beer

	BeerResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    *Beer  `json:"data"`
	}

	Beer struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Style Style  `json:"style"`
	}

	Breweries []Brewery
	Brewery   struct{}

	Category struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		CreateData string `json:"createData"`
	}

	Client struct {
		ApiKey     string
		httpClient *http.Client
	}

	Styles []Style
	Style  struct {
		Id     int    `json:"id"`
		SrmMax string `json:"srmMax"`
		IbuMax string `json:"ibuMax"`
		SrmMin string `json:"srmMin"`
		IbuMin string `json:"ibuMin"`
	}
)

func NewClient(apiKey string) *Client {
	return &Client{apiKey, &http.Client{}}
}

func (c *Client) Call(method, resource, payload string) (response []byte, err error) {
	rawUrl := fmt.Sprintf("%s/%s/%s", API_ENDPOINT, API_VERSION, resource)

	if len(c.ApiKey) > 0 {
		rawUrl += fmt.Sprintf("?key=%s", c.ApiKey)
	}

	log.Println(rawUrl)

	req, err := http.NewRequest(method, rawUrl, strings.NewReader(payload))
	if err != nil {
		return
	}

	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	response, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New("HTTP Error :" + res.Status)
		return
	}

	return
}

func (c *Client) GetBeer(id string) (beer *Beer, err error) {
	resource := fmt.Sprintf("beer/%s", id)
	r, err := c.Call("GET", resource, "")
	if err != nil {
		return
	}

	var response BeerResponse
	err = json.Unmarshal(r, &response)
	if err != nil {
		return
	}

	return response.Data, nil
}

func (c *Client) GetBeers() (*[]Beers, error) {
	return &[]Beers{}, nil
}

func (c *Client) GetBrewery() (*Brewery, error) {
	return &Brewery{}, nil
}

func (c *Client) GetBreweries() (*[]Breweries, error) {
	return &[]Breweries{}, nil
}
