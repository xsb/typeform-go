package tfcom

type Summary struct {
	Showing   int `json:"showing"`
	Total     int `json:"total"`
	Completed int `json:"completed"`
}

type Stats struct {
	Responses Summary `json:"responses"`
}

type Output struct {
	HttpStatus int           `json:"http_status"`
	Stats      Stats         `json:"stats"`
	Questions  []interface{} `json:"questions"`
	Responses  []interface{} `json:"responses"`
}
