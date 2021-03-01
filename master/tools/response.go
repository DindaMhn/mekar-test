package tools

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data, omitempty"`
}
