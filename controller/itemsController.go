package controller

import (
	"assignment2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type itemsInput struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}

// GetAllItems godoc
// @Summary Get all Items.
// @Description Get a list of Items.
// @Tags Items
// @Produce json
// @Success 200 {object} []models.Items
// @Router /items [get]
func GetAllItems(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var ratings []models.Items
	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// CreateItems godoc
// @Summary Create New Items.
// @Description Creating a new Items.
// @Tags Items
// @Param Body body itemsInput true "the body to create a new Items"
// @Produce json
// @Success 200 {object} models.Items
// @Router /items [post]
func CreateItems(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input itemsInput
	var order models.Orders
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.OrderID).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "orderID not found!"})
		return
	}

	// Create Item
	item := models.Items{ItemCode: input.ItemCode, Description: input.Description, Quantity: input.Quantity, OrderID: input.OrderID}
	db.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// UpdateItems godoc
// @Summary Update Items.
// @Description Update Items by id.
// @Tags Items
// @Produce json
// @Param id path string true "Items id"
// @Param Body body itemsInput true "the body to update age rating category"
// @Success 200 {object} models.Items
// @Router /items/{id} [patch]
func UpdateItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var item models.Items
	if err := db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input itemsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Items
	updatedInput.ItemCode = input.ItemCode
	updatedInput.Description = input.Description
	updatedInput.Quantity = input.Quantity
	updatedInput.UpdatedAt = time.Now()

	db.Model(&item).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// DeleteItems godoc
// @Summary Delete one Items.
// @Description Delete a Items by id.
// @Tags Items
// @Produce json
// @Param id path string true "Items id"
// @Success 200 {object} map[string]boolean
// @Router /items/{id} [delete]
func DeleteItem(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var item models.Items
	if err := db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
