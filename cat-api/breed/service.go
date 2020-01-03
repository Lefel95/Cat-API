package breed

import "cat-api/models"

type Service interface {
	GetBreedByName(breedName string) (*models.Breed, error)
}