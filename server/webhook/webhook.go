package webhook

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	log.Printf("Raw Data: %+v\n", payload)
	c.Status(http.StatusOK)
}
