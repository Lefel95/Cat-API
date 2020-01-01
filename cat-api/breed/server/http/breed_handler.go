package http

import (
	"cat-api/breed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service breed.Service
}

//NewHandler instantiate a handler and fix it with the router
func NewHandler(s breed.Service, r *gin.Engine) {
	h := handler{
		service: s,
	}

	h.AssignRoute(r)
}

//GetBreedByName gets a breed name from the client and tries to return the breed attributes
func (h *handler) GetBreedByName(c *gin.Context) {
	breedName := c.Query("name")

	if breedName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid breed name")})
		return
	}

	breed, err := h.service.GetBreedByName(breedName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"breed": breed})
}
