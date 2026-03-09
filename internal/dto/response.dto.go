package dto

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Error   string  `json:"error,omitempty"`
	Result  []Users `json:"result"`
}

type ResponseToken struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Token   string `json:"token,omitempty"`
}
