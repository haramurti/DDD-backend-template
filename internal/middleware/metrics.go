package middleware

import "github.com/gin-gonic/gin"

// Metrics TODO: Prometheus counter (method, route, status) + histogram latency.
// Lalu di bootstrap: r.GET("/metrics", gin.WrapH(promhttp.Handler()))
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
