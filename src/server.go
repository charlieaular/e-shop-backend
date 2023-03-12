package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"e-shop-backend/src/database"
	"e-shop-backend/src/routes"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	mysqlIns := database.MysqlConn

	db := mysqlIns.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisteAuthRoutes(r, db)
	routes.RegisteCategoryRoutes(r, db)
	routes.RegisteProductRoutes(r, db)
	routes.RegisteCartRoutes(r, db)
	routes.RegisterTagRoutes(r, db)
	routes.NewAddressRoutes(r, db)

	r.Run(":9000")
}
