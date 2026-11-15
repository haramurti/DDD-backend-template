package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	g := api.Group("/feature-b", auth)
	g.GET("", h.List)
	g.POST("", h.Create)
	g.GET("/:id", h.GetByID)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	// TODO: route khusus, contoh:
	// g.PATCH("/:id/status", middleware.RequireRole("manager"), h.ChangeStatus)
}
