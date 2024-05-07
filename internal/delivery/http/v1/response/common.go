package response

type Common struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
