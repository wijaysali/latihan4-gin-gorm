package handlers

import (
	"errors"
	"latihan4-gin-gorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCustomerHandler retrieves a customer by ID
// @Summary Retrieve customer
// @Description Get Customer details by ID including address and orders
// @ID get-customer
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} model.Customer
// @Router /customers/{id} [get]
func GetCustomerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customer model.Customer
		id := c.Param("id")

		if err := db.Preload("Address").Preload("Orders.OrderDetails").First(&customer, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}

func CreateCustomerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCustomer model.Customer
		if err := c.ShouldBindJSON(&newCustomer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if result := db.Create(&newCustomer); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, newCustomer)
	}
}
