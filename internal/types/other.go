package types

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Source struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Default bool   `json:"default,omitempty"`
}
