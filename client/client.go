package client

import (
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
