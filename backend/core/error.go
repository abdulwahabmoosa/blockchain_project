package core

type APIError struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
