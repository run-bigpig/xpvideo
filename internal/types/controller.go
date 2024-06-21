package types

type ListRequest struct {
	Class    int `json:"class"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type DetailRequest struct {
	Id int `json:"id"`
}

type SearchRequest struct {
	Keyword  string `json:"keyword"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type ClassResponse struct {
	Id   int    `json:"id"`
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}

type ListItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Actor    string `json:"actor"`
	Director string `json:"director"`
	Img      string `json:"img"`
}

type ListResponse struct {
	List  []*ListItem `json:"list"`
	Total int         `json:"total"`
}

type DetailResponse struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Img      string  `json:"img"`
	Actor    string  `json:"actor"`
	Director string  `json:"director"`
	Desc     string  `json:"desc"`
	Play     []*Play `json:"play"`
}

type Play struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SourceResponse struct {
	List    []*Source `json:"list"`
	Default *Source   `json:"default"`
}

type SettingRequest struct {
	Source string `json:"source"`
}
