package breed

import "cat-api/models"

//Repository sets the base methods to implement the access to local data of our service
type Repository interface {
	GetBreedByName(breedName string) (*models.Breed, error)
}
