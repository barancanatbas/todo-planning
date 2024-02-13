package dto

type FetchingResponse struct {
	Message string `json:"message"`
	Data    interface{}
}
