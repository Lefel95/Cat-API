package repository

import (
	"cat-api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	var err error
	db, mock, err = sqlmock.New()

	if err != nil {
		fmt.Println("expected no error, but got:", err)
		return
	}

	code := m.Run()

	db.Close()
	os.Exit(code)
}
func TestGetBreedByName(t *testing.T) {
	repo, err := NewRepository(db)

	if err != nil {
		assert.FailNowf(t, "error creating repo: %s", err.Error())
	}

	type test struct {
		InputName string
		ExpectedOutput *models.Breed
		ExpectedError error
	}

	var tests = make(map[string]test)

	tests["ReturnValue"] = test{
		InputName: "synx",
		ExpectedOutput: &models.Breed{
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
		ExpectedError: nil,
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			marsh, err := json.Marshal(test.ExpectedOutput)

			if err != nil {
				assert.NoErrorf(t, err, "Error marshaling json: %s", err)
			}

			rows := sqlmock.NewRows([]string{"attributes"}).AddRow(marsh)
			mock.ExpectPrepare("SELECT attributes FROM cats.breeds").ExpectQuery().WithArgs(test.InputName).WillReturnRows(rows)

			ret, err := repo.GetBreedByName(test.InputName)

			if assert.Equal(t, test.ExpectedError, err, "errors we expected: %s, but not nil has return: %s", test.ExpectedError, err) {
				assert.Equal(t, test.ExpectedOutput, ret, "We expected return: %#v but got: %#v", test.ExpectedOutput, ret)
			}
		})
	}
}
