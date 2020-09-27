package entities

// appstore 上查询应用返回的 json 格式数据
type AppASJson struct {
	ResultCount int     `json:"resultCount"`
	Results     []AppAS `json:"results"`
}
