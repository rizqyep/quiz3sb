package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HitungMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("hitung:%s", c.Query("hitung"))

		if c.Query("hitung") != "keliling" && c.Query("hitung") != "luas" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Parameter hitung harus bernilai keliling atau luas",
			})
			return
		}
		c.Next()
	}
}
