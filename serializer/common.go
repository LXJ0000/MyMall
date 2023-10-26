package serializer

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
