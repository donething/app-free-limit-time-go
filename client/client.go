package client

import (
	"app-free-limit-time-go/entities"
	"github.com/donething/utils-go/dohttp"
	"github.com/donething/utils-go/dowxpush"
	"time"
)

var (
	// 网络请求
	Request = dohttp.New(30*time.Second, false, false)
	// 微信推送
	WXPush *dowxpush.Sandbox = nil
)

func PushWX(info *entities.ExtraInfo, data *map[string]interface{}, url string) error {
	// 如果微信推送没有被实例化，则先实例化
	if WXPush == nil {
		WXPush = dowxpush.NewSandbox(info.Wxpush.Appid, info.Wxpush.Secret)
	}
	err := WXPush.PushTpl(info.Wxpush.Touid, info.Wxpush.Tplid, data, url)
	return err
}
