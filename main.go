package main

import (
	db "go-sujor/database"
	rt "go-sujor/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.DebugMode)
	defer db.SqlDB.Close()
	router := rt.InitRouter()
	router.Run(":3000")
}
