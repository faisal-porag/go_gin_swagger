package router_manager

import "github.com/gin-gonic/gin"

// MountCustomerRoutes mounts the "/customer" routes
func MountCustomerRoutes(router *gin.RouterGroup) {
	router.GET("/", homeHandler)
	router.POST("/sign-in", customerSignInHandler)
	router.GET("/profile", customerProfileHandler)
}

// @Summary Get home page.
// @Description Get home page details.
// @ID get-home-page
// @Produce json
// @Success 200
// @Router /customer [get]
func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home endpoint",
	})
}

// @Summary Customer sign in process.
// @Description Customer sign in.
// @ID customer-sign-in
// @Produce json
// @Success 200
// @Router /customer/sign-in [post]
func customerSignInHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer sign-in endpoint",
	})
}

// @Summary Get information about a customer.
// @Description Get details about a customer.
// @ID customer-profile
// @Produce json
// @Success 200
// @Router /customer/profile [get]
// @Security BearerToken
func customerProfileHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer profile endpoint",
	})
}
