package router_manager

import (
	"go_gin_swagger/params"
	"go_gin_swagger/response_structs"

	"github.com/gin-gonic/gin"
)

// MountCustomerRoutes mounts the "/customer" routes
func MountCustomerRoutes(router *gin.RouterGroup) {
	router.GET("/", homeHandler)
	router.POST("/signup", createCustomer)
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

// @Summary Create a new customer.
// @Description Create a new customer.
// @ID create-customer
// @Accept json
// @Produce json
// @Param input body params.CustomerInput true "Customer information to create"
// @Success 201 {object} response_structs.CustomerResponse
// @Router /customer/signup [post]
func createCustomer(c *gin.Context) {
	var input params.CustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Your logic to create a new customer with the provided input
	// ...

	// Return the created customer
	c.JSON(201, response_structs.CustomerResponse{
		ID:   3, // Replace with the actual ID of the created customer
		Name: input.Name,
		// Include other fields as needed
	})
}
