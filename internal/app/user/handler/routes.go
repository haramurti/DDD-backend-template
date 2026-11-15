package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	a := api.Group("/auth")
	a.POST("/register", h.Register)
	a.POST("/login", h.Login)

	api.GET("/me", auth, h.Me)
}
