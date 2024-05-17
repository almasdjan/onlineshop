package handler

import (
	"onlineshop/pkg/service"

	"github.com/gin-gonic/gin"

	_ "onlineshop/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signup)
		auth.POST("/login", h.login)

	}
	api := router.Group("/api", h.userIdentity)
	{

		admin := api.Group("/admin", h.checkAdmin)
		{
			products := admin.Group("/products")
			{
				products.POST("/", h.createProduct)
				products.DELETE("/:id", h.deleteProduct)
				products.PUT("/:id", h.updateProduct)
			}
			orders := admin.Group("/orders")
			{
				orders.GET("/", h.getAllOrders)
				orders.PUT("/:id", h.updateOrderStatus)
			}
		}

		products := api.Group("/products")
		{
			products.POST("/:id", h.addToCart)
			products.GET("/", h.getAlProdacts)
			products.GET("/:id", h.getProdactById)

		}
		cart := api.Group("/cart")
		{
			cart.PUT("/minus/:id", h.minus)
			cart.PUT("/plus/:id", h.plus)
			cart.DELETE("/", h.deleteAllFromCart)
			cart.GET("/", h.getAllFromCart)
			cart.POST("/order", h.createOrder)
			cart.GET("/orders", h.getOrderForUser)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router

}
