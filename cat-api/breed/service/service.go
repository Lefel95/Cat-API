package service

import (
	"cat-api/breed"
	"cat-api/models"
)

type service struct {
	repo breed.Repository
	gate breed.Gateway
}
func NewService(r breed.Repository, g breed.Gateway) breed.Service {
	return &service{
		repo: r,
		gate: g,
	}
}

func (s *service) GetBreedByName(breedName string) (*models.Breed, error) {
	breed, err := s.repo.GetBreedByName(breedName)

	if err != nil {
		return nil, err
	}

	if breed.ID == "" {
		breed, err = s.gate.GetBreedByName(breedName)
		if err != nil {
			return nil, err
		}

		err := s.repo.InsertBreed(breed)

		if err != nil {
			return breed, err
		}
	}

	return breed, nil
}