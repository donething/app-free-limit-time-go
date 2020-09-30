package entities

// 应用，格式为 id 和区域，如{"123": "cn", "345":"us"}
type App map[string]string

// 云函数 事件传递来的数据，包含微信消息推送的 token、需要检测价格的应用
type ExtraInfo struct {
	Wxpush struct {
		Appid  string `json:"appid"`
		Secret string `json:"secret"`
		Touid  string `json:"touid"`
		Tplid  string `json:"tplid"`
	} `json:"wxpush"`
	Apps App `json:"apps"`
}
