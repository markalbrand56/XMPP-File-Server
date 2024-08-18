package responses

type StandardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UploadSuccessResponse struct {
	StandardResponse
	Paths []string `json:"paths"`
}
