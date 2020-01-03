package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AssignRoute register the routes needed to get the breeds
func (h *handler) AssignRoute(r *gin.Engine, secret []byte) {
	breeds := r.GET("/breeds", h.GetBreedByName)
	breeds.Use(h.AuthMiddleware(secret))
}

//AuthMiddleware is a middleware to identify if the user is logged
func (h *handler) AuthMiddleware(secret []byte) gin.HandlerFunc{
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Bearer-Token")
		claims := &struct {
			Username string `json:"username"`
			jwt.StandardClaims
		}{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "Error making signature"})
				return
			}

			context.JSON(http.StatusBadRequest, gin.H{"error": "Error making signature"})
			return
		}

		if !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "tokens is not valid"})
			return
		}

		context.Next()
	}
}