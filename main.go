package main

import (
	. "app-free-limit-time-go/entities"
	"app-free-limit-time-go/utils"
	"encoding/json"
	"fmt"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"os"
	"strings"
)

// 入口函数
func Exec() (string, error) {
	// 获取、解析 数据
	var data = os.Getenv("DATA")
	if strings.TrimSpace(data) == "" {
		return "传入的信息为空，退出执行", nil
	}

	var extraInfo ExtraInfo
	err := json.Unmarshal([]byte(data), &extraInfo)
	if err != nil {
		fmt.Printf("解析额外信息时出错，无法解析文本：%s\n", err)
		return "解析额外信息时出错，无法解析文本", err
	}

	// 应用价格
	fmt.Printf("开始检测 关注应用的价格\n")
	total := len(extraInfo.Apps)
	free, failed := utils.CheckPrice(&extraInfo)

	fmt.Printf("关注的应用 检测完成\n")
	result := fmt.Sprintf("共 %d 个应用，%d 个限免，%d 个检测失败", total, free, failed)
	return result, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Exec)
}
