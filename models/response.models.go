package models

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Total   uint        `json:"total"`
}

type ResponseCart struct {
	Data   Cart `json:"data"`
	Status bool `json:"status"`
}

type ResponseProduct struct {
	Data   Product `json:"data"`
	Status bool    `json:"status"`
}
