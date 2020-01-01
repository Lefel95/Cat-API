package http

import "github.com/gin-gonic/gin"

//AssignRoute register the routes needed to get the breeds
func (h *handler) AssignRoute(r *gin.Engine) {
	breed := r.Group("/breeds")
	breed.Use(h.AuthMiddleware())
	breed.GET("/", h.GetBreedByName)
}

//AuthMiddleware is a middleware to identify if the user is logged
func (h *handler) AuthMiddleware() gin.HandlerFunc{
	return func(context *gin.Context) {
		
	}
}