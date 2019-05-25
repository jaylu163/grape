package main

import (
	"github.com/gin-gonic/gin"
	db "grape/database"
	rt "grape/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	defer db.SqlDB.Close()
	router := rt.InitRouter()
	router.Run(":7621")
}
