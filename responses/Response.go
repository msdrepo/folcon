package responses

type SuccessResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Token      interface{} `json:"token"`
}

type ErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
