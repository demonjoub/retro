package schema

type Response struct {
	Code int         `json:"code" xml:"code"`
	Body interface{} `json:"data" xml:"data"`
}

type Body struct {
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}
