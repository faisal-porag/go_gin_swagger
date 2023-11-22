package router_manager

import "github.com/gin-gonic/gin"

// MountCustomerRoutes mounts the "/customer" routes
func MountCustomerRoutes(router *gin.RouterGroup) {
	router.GET("/", customerHandler)
	router.GET("/sign-in", customerSignInHandler)
}

// @Summary Get information about a customer.
// @Description Get details about a customer.
// @ID get-customer
// @Produce json
// @Success 200
// @Router /customer [get]
func customerHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer endpoint",
	})
}

func customerSignInHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer sign-in endpoint",
	})
}
