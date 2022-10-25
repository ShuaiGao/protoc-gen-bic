package gen

type Response struct {
	Code   int         `json:"code"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}
