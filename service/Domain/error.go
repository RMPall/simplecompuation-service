package domain

// ErrorResponse contains error response attributes including status, code and message.
type ErrorResponse struct {
	Status  int
	Code    int
	Message string
}
