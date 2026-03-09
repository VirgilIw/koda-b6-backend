package dto

type Response struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Error   string  `json:"error,omitempty"`
	Result  []Users `json:"result"`
}

type ResponseOneData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  Users  `json:"result"`
}

type ResponseToken struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Token   string `json:"token,omitempty"`
}

type ResponseError struct {
	Status  string `json:"status" example:"Error"`
	Message string `json:"message" example:"Failed get data"`
	Error   string `json:"errors,omitempty" example:"failed get data"`
}
