package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zignd/widgets-api/api"
	"github.com/zignd/widgets-api/services"
)

func init() {
	if err := services.ConfigViper(); err != nil {
		panic(errors.Wrap(err, "failed to configure Viper"))
	}

	log.SetFormatter(&log.JSONFormatter{})
	// TODO: Add a Logrus hook capable of sending our logs somewhere useful, ElasticSearch for example.
	// TODO: Add to the logs the correlation ids the CorrelationIDMiddleware is appending to requests and responses.
	// TODO: Add a Gin middleware capable of logging requests and response (don't forget the most important header, "Correlation-ID").
}

func main() {
	router := gin.Default()
	router.Use(services.ConnStrMiddleware(), services.CorrelationIDMiddleware())

	router.GET("/users/:id", api.GetUserByID)
	router.GET("/users", api.GetUsers)
	router.GET("/widgets/:id", api.GetWidgetByID)
	router.GET("/widgets", api.GetWidgets)

	router.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
}
