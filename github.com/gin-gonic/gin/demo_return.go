package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("text", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%s", "world")
	})
	router.GET("json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "55555555555",
			"tag":  "9999999",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
