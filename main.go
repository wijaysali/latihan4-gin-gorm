package main

import (
	_ "latihan4-gin-gorm/docs"
	"latihan4-gin-gorm/handlers"
	"latihan4-gin-gorm/model"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	dsn := "host=localhost user=postgres dbname=tokoonline password= postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&model.Customer{}, &model.Address{}, &model.Product{}, &model.Order{}, &model.OrderDetail{}, &model.Category{}); err != nil {
		log.Fatal("failed to auto migrate:", err)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/customers/:id", handlers.GetCustomerHandler(db))
	r.POST("/customers", handlers.CreateCustomerHandler(db))

	r.GET("/products/:id", handlers.GetProductHandler(db))
	r.GET("/orders/:id", handlers.GetOrderHandler(db))

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
