package entities

// SuccessResponse represents a success response
// @Description Standard success response format
type SuccessResponse struct {
	Status string `json:"status" example:"success"`
}

// ErrorResponse represents an error response
// @Description Standard error response format
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}
