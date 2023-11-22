package params

type CustomerInput struct {
	Name string `json:"name" binding:"required"`
	// Add other input fields as needed
}
