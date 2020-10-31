package utils

import (
	. "app-free-limit-time-go/client"
	. "app-free-limit-time-go/entities"
	"fmt"
	"github.com/donething/utils-go/dostr"
	"strconv"
	"time"
)

// 检测应用的价格
func CheckPrice(info *ExtraInfo) (free int, failed int) {
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
				// 推送消息
				data := WXPush.GenGeneralTpl("已限免！！"+appInfo.Name,
					fmt.Sprintf("AppStore 上“%s”已限免，点击去下载", appInfo.Name),
					dostr.FormatDate(dostr.BeiJingTime(time.Now()), dostr.TimeFormatDefault))
				err = PushWX(info, data, appInfo.TrackViewUrl)
				if err != nil {
					fmt.Printf("推送 AppStore 应用(%d)限免的微信消息时出错：%s\n", appInfo.TrackId, err)
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
