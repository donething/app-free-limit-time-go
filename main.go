package main

import (
	. "app-free-limit-time-go/entities"
	"encoding/json"
	"fmt"
	"github.com/donething/utils-go/dostr"
	"github.com/donething/utils-go/dowxpush"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"strconv"
	"time"
)

// 微信推送
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
	fmt.Printf("开始检测 关注应用的价格\n")
	extraInfo, err := unmarshal(event.Message)
	if err != nil {
		fmt.Printf("解析额外信息时出错，无法解析文本：%s\n", err)
		return "解析额外信息时出错，无法解析文本", err
	}
	total := len(extraInfo.Apps)
	free, failed := checkPrice(extraInfo)

	fmt.Printf("关注的应用 检测完成\n")
	result := fmt.Sprintf("共 %d 个应用，%d 个限免，%d 个检测失败", total, free, failed)
	return result, nil
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
func checkPrice(info ExtraInfo) (free int, failed int) {
	//  检测 appstore 商店应用的价格
	for _, app := range info.Apps {
		if app.Plat == PlatAppstore {
			// 转化字符串 id 为数字
			id, err := strconv.ParseInt(app.ID, 10, 64)
			if err != nil {
				failed++
				fmt.Printf("填写的 appstore 应用的id(%s)无法转为数字：%v\n", app.ID, err)
				continue
			}

			// 填充应用信息
			appInfo := AppAS{TrackId: id, Area: app.Area}
			err = appInfo.Fill()
			if err != nil {
				failed++
				fmt.Printf("填充 appstore 应用(id %d)的信息时出错：%v\n", appInfo.TrackId, err)
				continue
			}

			// 发现限免应用
			if appInfo.Price == 0 {
				free++
				fmt.Printf("AppStore 上“%s”(id %d)已限免，将发送消息通知\n", appInfo.Name, appInfo.TrackId)
				// 如果微信推送没有被实例化，则先实例化
				if wxPush == nil {
					wxPush = dowxpush.NewSandbox(info.Wxpush.Appid, info.Wxpush.Secret)
				}
				// 推送消息
				data := wxPush.GenGeneralTpl("已限免！！"+appInfo.Name,
					fmt.Sprintf("AppStore 上“%s”已限免，点击去下载", appInfo.Name),
					dostr.FormatDate(dostr.BeiJingTime(time.Now()), dostr.TimeFormatDefault))
				err := wxPush.PushTpl(info.Wxpush.Touid, info.Wxpush.Tplid, data, appInfo.TrackViewUrl)
				if err != nil {
					fmt.Printf("推送 AppStore 应用(id %d)限免的微信消息时出错：%s\n", appInfo.TrackId, err)
					continue
				}
			}
		} else if app.Plat == PlatPlaystore {
			//  检测 playstore 商店应用的价格
		} else {
			failed++
			fmt.Printf("未知的平台：%+v\n", app)
		}
	}
	return
}
