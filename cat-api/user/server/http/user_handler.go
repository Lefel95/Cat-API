package http

import (
	"cat-api/models"
	"cat-api/user"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service user.Service
}

//NewHandler instantiate a handler and fix it with the router
func NewHandler(s user.Service, r *gin.Engine) {
	h := handler{
		service: s,
	}

	h.AssignRoute(r)
}

//Login get user credentials and if exists returns the token for client authentication
func (h *handler) Login(c *gin.Context) {
	req	, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("could not get credentials")})
		return
	}

	var login models.UserLogin
	err = json.Unmarshal(req, &login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("could not get credentials")})
		return
	}

	tokenString, exists, err := h.service.Login(login)

	if err != nil {
		if exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return
}