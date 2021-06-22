package interfaces

// GetHealthCheckRequest - health check request
type GetHealthCheckRequest struct{}

// GetHealthCheckResponse - health check response
type GetHealthCheckResponse struct {
	Status int `json:"status"` // status code
}
