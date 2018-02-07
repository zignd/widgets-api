package services

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func ConnStrMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ConnStr", viper.GetString("postgresql_connection_string"))
	}
}

func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := uuid.Must(uuid.NewV4()).String()
		c.Request.Header.Set("Correlation-ID", correlationID)
		c.Header("Correlation-ID", correlationID)
	}
}
