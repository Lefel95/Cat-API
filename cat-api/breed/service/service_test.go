package service

import (
	"cat-api/breed/mocks"
	"cat-api/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGetBreedByNameFromRepository(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repo := mocks.NewMockRepository(controller)
	gate := mocks.NewMockGateway(controller)

	service := NewService(repo, gate)

	test := struct {
		Input string
		ExpectedResult *models.Breed
		ExpectedError error
	}{Input: "synx",
		ExpectedResult: &models.Breed{
			Adaptability:     5,
			AffectionLevel:   2,
			AltNames:         "",
			CFAURL:           "http://test.com/api/synx",
			ChildFriendly:    2,
			CountryCode:      "BR",
			CountryCodes:     "BR, EUA",
			Description:      "Just a test cat",
			DogFriendly:      2,
			EnergyLevel:      6,
			Experimental:     3,
			Grooming:         4,
			Hairless:         0,
			HealthIssues:     0,
			HypoAllergenic:   5,
			ID:               "synx",
			Indoor:           2,
			Intelligence:     5,
			Lap:              3,
			LifeSpan:         "5-10 years",
			Name:             "Synx Examples",
			Natural:          2,
			Origin:           "N/A",
			Rare:             2,
			Rex:              3,
			SheddingLevel:    4,
			ShortLegs:        2,
			SocialNeeds:      0,
			StrangerFriendly: 0,
			SupressedTail:    0,
			Temperament:      "Calm",
			VCAHospitalsURL:  "http://test.com/api/synx",
			VetStreetURL:     "http://test.com/api/synx",
			Vocalisation:     0,
			Weight:           models.Weight{
				Imperial: "3 foot",
				Metric: "0.5 meters",
			},
			WikipediaURL:     "http://test.com/api/synx",
		},
		ExpectedError: nil,}

	repo.EXPECT().GetBreedByName(test.Input).Return(test.ExpectedResult, test.ExpectedError)

	ret, err := service.GetBreedByName(test.Input)

	if assert.Equal(t, test.ExpectedError, err, "errors we expected: %s, but not nil has return: %s", test.ExpectedError, err) {
		assert.Equal(t, test.ExpectedResult, ret, "We expected return: %#v but got: %#v", test.ExpectedResult, ret)
	}
}

func TestGetBreedByNameFromGateway(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repo := mocks.NewMockRepository(controller)
	gate := mocks.NewMockGateway(controller)

	service := NewService(repo, gate)

	test := struct {
		Input string
		ExpectedResult *models.Breed
		ExpectedError error
	}{Input: "synx",
		ExpectedResult: &models.Breed{
			Adaptability:     5,
			AffectionLevel:   2,
			AltNames:         "",
			CFAURL:           "http://test.com/api/synx",
			ChildFriendly:    2,
			CountryCode:      "BR",
			CountryCodes:     "BR, EUA",
			Description:      "Just a test cat",
			DogFriendly:      2,
			EnergyLevel:      6,
			Experimental:     3,
			Grooming:         4,
			Hairless:         0,
			HealthIssues:     0,
			HypoAllergenic:   5,
			ID:               "synx",
			Indoor:           2,
			Intelligence:     5,
			Lap:              3,
			LifeSpan:         "5-10 years",
			Name:             "Synx Examples",
			Natural:          2,
			Origin:           "N/A",
			Rare:             2,
			Rex:              3,
			SheddingLevel:    4,
			ShortLegs:        2,
			SocialNeeds:      0,
			StrangerFriendly: 0,
			SupressedTail:    0,
			Temperament:      "Calm",
			VCAHospitalsURL:  "http://test.com/api/synx",
			VetStreetURL:     "http://test.com/api/synx",
			Vocalisation:     0,
			Weight:           models.Weight{
				Imperial: "3 foot",
				Metric: "0.5 meters",
			},
			WikipediaURL:     "http://test.com/api/synx",
		},
		ExpectedError: nil,}

	repo.EXPECT().GetBreedByName(test.Input).Return(&models.Breed{}, test.ExpectedError)
	gate.EXPECT().GetBreedByName(test.Input).Return(test.ExpectedResult, test.ExpectedError)

	ret, err := service.GetBreedByName(test.Input)

	if assert.Equal(t, test.ExpectedError, err, "errors we expected: %s, but not nil has return: %s", test.ExpectedError, err) {
		assert.Equal(t, test.ExpectedResult, ret, "We expected return: %#v but got: %#v", test.ExpectedResult, ret)
	}
}