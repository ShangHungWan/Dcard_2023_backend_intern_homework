package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsStatusOK(status int) bool {
	okStatus := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNonAuthoritativeInfo,
		http.StatusNoContent,
		http.StatusResetContent,
		http.StatusPartialContent,
		http.StatusMultiStatus,
		http.StatusAlreadyReported,
		http.StatusIMUsed,
	}

	for _, value := range okStatus {
		if value == status {
			return true
		}
	}

	return false
}

func APIResponse(c *gin.Context, status int, obj any) {
	if IsStatusOK(status) {
		c.IndentedJSON(status, obj)
	} else {
		c.IndentedJSON(status, gin.H{"error": obj})
	}
}

func ToInterface(slice []string) []interface{} {
	s := make([]interface{}, len(slice))
	for i, v := range slice {
		s[i] = v
	}
	return s
}
