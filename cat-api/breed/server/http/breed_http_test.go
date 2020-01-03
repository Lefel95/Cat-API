package http

import (
	"cat-api/breed/mocks"
	"cat-api/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type test struct {
	Input string
	StatusCode int
	Breed *models.Breed
	Error error
	ExpectedOutput []byte
}

func TestGetBreedByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	r := gin.New()
	secret := []byte("some secret test")
	NewHandler(service, r, secret)
	tokenstring := makeTestToken(secret)

	w := httptest.NewRecorder()

	theTest := test{
		Input:          "synx",
		Breed:          &models.Breed{
			Adaptability:     0,
			AffectionLevel:   0,
			AltNames:         "",
			CFAURL:           "",
			ChildFriendly:    0,
			CountryCode:      "BR",
			CountryCodes:     "BR",
			Description:      "",
			DogFriendly:      0,
			EnergyLevel:      0,
			Experimental:     0,
			Grooming:         0,
			Hairless:         0,
			HealthIssues:     0,
			HypoAllergenic:   0,
			ID:               "synx",
			Indoor:           0,
			Intelligence:     0,
			Lap:              0,
			LifeSpan:         "1-6 years",
			Name:             "Synx Tests",
			Natural:          0,
			Origin:           "",
			Rare:             0,
			Rex:              0,
			SheddingLevel:    0,
			ShortLegs:        0,
			SocialNeeds:      0,
			StrangerFriendly: 0,
			SupressedTail:    0,
			Temperament:      "",
			VCAHospitalsURL:  "",
			VetStreetURL:     "",
			Vocalisation:     0,
			Weight:           models.Weight{
				Imperial: "1-6 meters",
				Metric:"1-5 meters",
			},
			WikipediaURL:     "",
		},
		Error:          nil,
		StatusCode: http.StatusOK,
	}

	service.EXPECT().GetBreedByName(theTest.Input).Return(theTest.Breed, theTest.Error)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/breeds?name=%s", theTest.Input), nil)
	req.Header.Add("Bearer-Token", tokenstring)

	r.ServeHTTP(w, req)

	assert.Equal(t, theTest.StatusCode, w.Code)
}

func TestGetBreedByNameEmptyError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	r := gin.New()
	secret := []byte("some secret test")
	NewHandler(service, r, secret)
	tokenstring := makeTestToken(secret)

	w := httptest.NewRecorder()

	theTest := test{
		Input:          "",
		StatusCode:     http.StatusBadRequest,
		Breed:          nil,
		Error:          fmt.Errorf("invalid Breed name"),
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/breeds?name=%s", theTest.Input), nil)
	req.Header.Add("Bearer-Token", tokenstring)

	r.ServeHTTP(w, req)

	assert.Equal(t, theTest.StatusCode, w.Code)
}

func TestGetBreedByNameInternalError (t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockService(ctrl)
	r := gin.New()
	secret := []byte("some secret test")
	NewHandler(service, r, secret)
	tokenstring := makeTestToken(secret)

	w := httptest.NewRecorder()

	theTest := test{
		Input:          "asudhas",
		StatusCode:     http.StatusInternalServerError,
		Breed:          nil,
		Error:          fmt.Errorf("breed Not Found"),
	}

	service.EXPECT().GetBreedByName(theTest.Input).Return(theTest.Breed, theTest.Error)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/breeds?name=%s", theTest.Input), nil)
	req.Header.Add("Bearer-Token", tokenstring)

	r.ServeHTTP(w, req)

	assert.Equal(t, theTest.StatusCode, w.Code)
}

func makeTestToken(secret []byte) string {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &struct{
		Username string `json:"username"`
		jwt.StandardClaims
	}{
		Username: "John Doe",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return ""
	}

	return tokenString
}