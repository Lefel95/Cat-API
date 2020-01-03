package http

import "github.com/gin-gonic/gin"

//AssignRoute register the routes needed to get the breeds
func (h *handler) AssignRoute(r *gin.Engine) {
	r.POST("/login", h.Login)
}
