package response

type CustomResponse[T any] struct {
	Data    T      `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
