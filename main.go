package main

import (
	. "app-free-limit-time-go/entities"
	"app-free-limit-time-go/utils"
	"encoding/json"
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

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
	// 解析额外的数据
	var extraInfo ExtraInfo
	err := json.Unmarshal([]byte(event.Message), &extraInfo)
	if err != nil {
		fmt.Printf("解析额外信息时出错，无法解析文本：%s\n", err)
		return "解析额外信息时出错，无法解析文本", err
	}
	total := len(extraInfo.Apps)
	free, failed := utils.CheckPrice(extraInfo)

	fmt.Printf("关注的应用 检测完成\n")
	result := fmt.Sprintf("共 %d 个应用，%d 个限免，%d 个检测失败", total, free, failed)
	return result, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Exec)
}
