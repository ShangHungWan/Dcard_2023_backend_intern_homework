package helper

import (
	"key-value-system/db"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func ClearDB() error {
	_, err := db.DB.Exec(db.GetSql(db.DELETE_ALL_SQL))
	return err
}
