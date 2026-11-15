package realtime

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

// Hub WebSocket dengan room (contoh room: "global", "user:{id}", "event:{id}").
type Hub struct {
	// TODO: mu sync.RWMutex; rooms map[string]map[*client]bool
}

func NewHub() *Hub { return &Hub{} }

// Broadcast kirim pesan ke semua client di room. Panggil SETELAH transaksi commit.
func (h *Hub) Broadcast(room string, msg Message) {
	// TODO
}

// LeaveRoom cabut user dari room (saat dikeluarkan / role dicabut).
func (h *Hub) LeaveRoom(userID uint, room string) {
	// TODO
}

// ServeWS route: GET /api/v1/ws?token=JWT (lewat middleware Auth dulu).
func (h *Hub) ServeWS(c *gin.Context) {
	// TODO: tentuin room yang boleh di-join (cek DB), upgrade pakai gorilla/websocket,
	// jalanin readPump + writePump (ping 30s).
	c.JSON(http.StatusNotImplemented, gin.H{"success": false, "message": "websocket not implemented"})
}
