package models

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Total   uint        `json:"total"`
}
