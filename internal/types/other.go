package types

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Source struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Key  string `json:"key"`
}
