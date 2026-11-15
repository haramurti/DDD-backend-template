package handler

import (
	"net/http"

	"hackathon-be/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Me(c *gin.Context) {
	res, err := h.svc.Me(c.Request.Context(), c.GetUint("uid"))
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}
