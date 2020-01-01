package breed

import (
	"cat-api/models"
)

//Gateway is a interface to bring external breed data
type Gateway interface {
	GetBreedByName(breedName string) (*models.Breed, error)
}
