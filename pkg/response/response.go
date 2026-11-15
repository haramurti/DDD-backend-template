package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Body struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

type PageMeta struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

func OK(c *gin.Context, status int, data any) {
	c.JSON(status, Body{Success: true, Data: data})
}

func Paginated(c *gin.Context, data any, meta PageMeta) {
	c.JSON(http.StatusOK, Body{Success: true, Data: data, Meta: meta})
}

// Error mapping error domain -> HTTP status. Selalu `return` setelah manggil ini.
func Error(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, ErrBadRequest):
		status = http.StatusBadRequest
	case errors.Is(err, ErrUnauthorized):
		status = http.StatusUnauthorized
	case errors.Is(err, ErrForbidden):
		status = http.StatusForbidden
	case errors.Is(err, ErrNotFound), errors.Is(err, gorm.ErrRecordNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrConflict), errors.Is(err, gorm.ErrDuplicatedKey):
		status = http.StatusConflict
	case errors.Is(err, ErrNotImplemented):
		status = http.StatusNotImplemented
	}
	msg := err.Error()
	if status == http.StatusInternalServerError {
		_ = c.Error(err) // detail masuk log, client cukup pesan umum
		msg = "internal server error"
	}
	c.JSON(status, Body{Success: false, Message: msg})
}

// ValidationError buat error dari ShouldBindJSON / ShouldBindQuery.
func ValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Body{Success: false, Message: err.Error()})
}
