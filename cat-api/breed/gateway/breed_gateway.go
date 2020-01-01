package gateway

import (
	"cat-api/breed"
	"cat-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type catAPI struct {
	baseurl string
	key     string
}

//NewGateway returns a new instance of breed.Gateway
func NewGateway() breed.Gateway {
	return &catAPI{
		baseurl: "https://api.thecatapi.com/v1/breeds",
		key:     "",
	}
}

//GetBreedByName receives a name and tries to found the desired breed in API
func (api *catAPI) GetBreedByName(breedName string) (*models.Breed, error) {
	breeds, err := api.getBreeds()

	if err != nil {
		return nil, err
	}

	for _, breed := range breeds {
		if breed.Name == breedName {
			return &breed, nil
		}
	}

	return nil, fmt.Errorf("Could not find breed: %s", breedName)
}

func (api *catAPI) getBreeds() ([]models.Breed, error) {
	req, err := http.NewRequest(http.MethodGet, api.baseurl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("x-api-key", api.key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	var breeds []models.Breed

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil
}
