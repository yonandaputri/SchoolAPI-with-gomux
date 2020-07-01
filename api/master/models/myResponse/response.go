package myResponse

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
