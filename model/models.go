package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Email      string `gorm:"unique"`
	Address    Address
	Orders     []Order
}

type Address struct {
	gorm.Model `swaggerignore:"true"`
	CustomerID uint
	Address    string
}

type Product struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
}

type Order struct {
	gorm.Model   `swaggerignore:"true"`
	CustomerID   uint
	OrderDate    time.Time
	TotalPrice   float64
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	gorm.Model `swaggerignore:"true"`
	OrderID    uint
	ProductID  uint
	Quantity   int
	Price      float64
}

type Category struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Product    []Product `gorm:"many2many:product_categories;"`
}

type ProductCategory struct {
	ProductID  uint
	CategoryID uint
}
