package main

import (
	"github.com/gin-gonic/gin"
	"gram/db"
	"gram/models"
)

func main() {
	conn := db.DbConn()

	err := models.Migrate(conn)
	if err != nil {
		return
	}

	r := gin.Default()

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
