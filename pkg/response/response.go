package response

type Response struct {
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Ok(data any, message string) *Response {
	return &Response{Data: data, Msg: message}
}

func Err(err error) *Response {
	return &Response{Data: nil, Msg: err.Error()}
}
