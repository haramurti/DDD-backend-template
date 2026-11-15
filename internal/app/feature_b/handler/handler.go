package handler

import (
	"net/http"
	"strconv"

	"hackathon-be/internal/app/feature_b/contract"
	"hackathon-be/internal/app/feature_b/dto"
	"hackathon-be/pkg/response"

	"github.com/gin-gonic/gin"
)

// Handler cuma: bind/validasi input -> panggil service -> tulis response.
type Handler struct {
	svc contract.FeatureBService
}

func New(svc contract.FeatureBService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateFeatureBRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.svc.Create(c.Request.Context(), c.GetUint("uid"), req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusCreated, res)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	res, err := h.svc.GetByID(c.Request.Context(), c.GetUint("uid"), id)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) List(c *gin.Context) {
	var q dto.ListFeatureBQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.ValidationError(c, err)
		return
	}
	items, total, err := h.svc.List(c.Request.Context(), c.GetUint("uid"), q)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Paginated(c, items, response.PageMeta{Page: q.Page, Size: q.Size, Total: total})
}

func (h *Handler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req dto.UpdateFeatureBRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.svc.Update(c.Request.Context(), c.GetUint("uid"), id, req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.Delete(c.Request.Context(), c.GetUint("uid"), id); err != nil {
		response.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{Success: false, Message: "invalid id"})
		return 0, false
	}
	return uint(id), true
}
