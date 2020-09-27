package main

import (
	. "app-free-limit-time-go/entities"
	"encoding/json"
	"fmt"
	"github.com/donething/utils-go/dostr"
	"github.com/donething/utils-go/dowxpush"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"time"
)

var wxPush *dowxpush.Sandbox = nil

// 定时触发器的入参，需要传递的数据在 Message 中
type TimerEvent struct {
	Message     string
	Time        string
	TriggerName string
	Type        string
}

// 入口函数
func Exec(event TimerEvent) (string, error) {
	//fmt.Printf("收到的额外信息：%s\n", event.Message)
	fmt.Printf("开始检测应用的价格\n")
	extraInfo, err := unmarshal(event.Message)
	if err != nil {
		fmt.Printf("解析额外信息时出错，无法解析文本：%s\n", err)
		return "解析额外信息时出错，无法解析文本", err
	}
	total, free, failed := checkPrice(extraInfo)

	return fmt.Sprintf("共 %d 个应用，%d 个限免，%d 个检测失败", total, free, failed), nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Exec)
}

// 解析额外的数据
func unmarshal(message string) (ExtraInfo, error) {
	var apps ExtraInfo
	err := json.Unmarshal([]byte(message), &apps)
	return apps, err
}

// 检测应用的价格
func checkPrice(info ExtraInfo) (total int, free int, failed int) {
	// 获取各平台的应用 id
	as, ps := info.IDs.As, info.IDs.Ps
	total = len(as) + len(ps)
	//  检测 appstore 商店应用的价格
	for _, id := range as {
		app := AppAS{TrackId: id}
		err := app.Fill()
		// 出错
		if err != nil {
			failed++
			fmt.Printf("填充 appstore 应用(%d)的信息时出错：%v\n", app.TrackId, err)
			continue
		}
		// 发现限免应用
		if app.Price == 0 {
			free++
			fmt.Printf("AppStore 中“%s”(id %d)已限免，将发送消息通知", app.Name, app.TrackId)
			// 如果微信推送没有被实例化，则先实例化
			if wxPush == nil {
				wxPush = dowxpush.NewSandbox(info.Wxpush.Appid, info.Wxpush.Secret)
			}
			// 推送消息
			data := wxPush.GenGeneralTpl("关注的应用已限免",
				fmt.Sprintf("AppStore 中“%s”已限免，点击跳转到 AppStore 商店下载", app.Name),
				dostr.FormatDate(dostr.BeiJingTime(time.Now()), dostr.TimeFormatDefault))
			err := wxPush.PushTpl(info.Wxpush.Touid, info.Wxpush.Tplid, data, app.TrackViewUrl)
			if err != nil {
				fmt.Printf("推送 AppStore 应用(id %d)限免的微信消息时出错：%s\n", app.TrackId, err)
				continue
			}
		}
	}

	//  检测 playstore 商店应用的价格
	return
}
