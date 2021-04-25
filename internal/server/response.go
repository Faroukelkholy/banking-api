package server

type HTTPResponse struct {
	Data interface{} `json:"data"`
	Err  HTTPError   `json:"error"`
}

type HTTPError struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
