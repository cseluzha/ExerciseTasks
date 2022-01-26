package models

type ResponseStatus struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Status `json:"data"`
	Errors  []string `json:"errors"`
}

type ResponseTasks struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Task   `json:"data"`
	Errors  []string `json:"errors"`
}
