package controller

import (
	"assignment2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ordersInput struct {
	OrderID      uint      `gorm:"primary_key" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}

// GetAllOrders godoc
// @Summary Get all Orders.
// @Description Get a list of Order.
// @Tags Order
// @Produce json
// @Success 200 {object} []models.Orders
// @Router /orders [get]
func GetAllOrder(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var orders []models.Orders
	db.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// CreateOrder godoc
// @Summary Create New Order.
// @Description Creating a new Order.
// @Tags Order
// @Param Body body ordersInput true "the body to create a new Order"
// @Produce json
// @Success 200 {object} models.Orders
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	// Validate input
	var input ordersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	order := models.Orders{CustomerName: input.CustomerName, OrderedAt: input.OrderedAt}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// UpdateOrder godoc
// @Summary Update Order.
// @Description Update Order by id.
// @Tags Order
// @Produce json
// @Param id path string true "Order id"
// @Param Body body ordersInput true "the body to update order"
// @Success 200 {object} models.Orders
// @Router /orders/{id} [patch]
func UpdateOrder(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var order models.Orders
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ordersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Orders
	updatedInput.CustomerName = input.CustomerName
	updatedInput.OrderedAt = input.OrderedAt
	updatedInput.UpdatedAt = time.Now()

	db.Model(&order).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder godoc
// @Summary Delete one Order.
// @Description Delete a Order by id.
// @Tags Order
// @Produce json
// @Param id path string true "Order id"
// @Success 200 {object} map[string]boolean
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var order models.Orders
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
