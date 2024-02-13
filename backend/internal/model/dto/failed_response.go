package dto

type FailedResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
