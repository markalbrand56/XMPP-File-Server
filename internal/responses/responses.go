package responses

type StandardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
