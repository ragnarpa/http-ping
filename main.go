package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		logrus.
			WithFields(
				logrus.Fields{
					"agent": c.Request.UserAgent(),
				},
			).
			Info("/ping called")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
