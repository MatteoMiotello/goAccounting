package api

import "github.com/gin-gonic/gin"

func ResFromData(data any) map[string]any {
	return gin.H{
		"data": data,
	}
}

func ResFromError(err error) map[string]any {
	return gin.H{
		"error": err.Error(),
	}
}

func ResFromMessage(message string) map[string]any {
	return gin.H{
		"error": message,
	}
}
