package handlers

import (
	"errors"
	"latihan4-gin-gorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrderHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order model.Order
		id := c.Param("id")

		if result := db.Preload("OrderDetails").First(&order, id); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}
