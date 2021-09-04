package main

import (
	"food-delivery/component"
	"food-delivery/modules/restaurants/transport"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading env file")
	}

	dbUri := os.Getenv("DB_URI")
	db, err := gorm.Open(mysql.Open(dbUri))

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)
	restaurants := r.Group("/restaurants")
	{
		restaurants.GET("/:id", transport.GetRestaurant(appCtx))
		restaurants.GET("", transport.ListRestaurant(appCtx))
		restaurants.POST("", transport.CreateRestaurant(appCtx))
	}
	return r.Run()
}
