package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"assignment2/controller"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/items", controller.GetAllItems)
	r.POST("/items", controller.CreateItems)
	r.PATCH("/items/:id", controller.UpdateItem)
	r.DELETE("/items/:id", controller.DeleteItem)

	r.GET("/orders", controller.GetAllOrder)
	r.POST("/orders", controller.CreateOrder)
	r.PATCH("/orders/:id", controller.UpdateOrder)
	r.DELETE("/orders/:id", controller.UpdateOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
