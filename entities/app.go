package entities

import (
	"app-free-limit-time-go/utils"
	"encoding/json"
	"fmt"
)

// 没有查找到应用
var ErrNotFound = fmt.Errorf("没有查找到应用")

// appstore 上的应用
// 根据 TrackId 获取应用信息：
// app := AppAS{TrackId: 1261944766}
// err := app.Fill()
type AppAS struct {
	TrackId        int64   `json:"trackId"`
	Name           string  `json:"trackName"`
	Price          float32 `json:"price"`
	Currency       string  `json:"currency"`
	FormattedPrice string  `json:"formattedPrice"`
	BundleId       string  `json:"bundleId"`
	TrackViewUrl   string  `json:"trackViewUrl"`
}

// 填充应用信息
func (app *AppAS) Fill() error {
	// 获取应用信息
	queryUrl := fmt.Sprintf("http://itunes.apple.com/lookup?country=cn&id=%d", app.TrackId)
	bs, err := utils.Client.Get(queryUrl, nil)
	if err != nil {
		return err
	}
	// 解析 json
	var payload AppASJson
	err = json.Unmarshal(bs, &payload)
	if err != nil {
		return err
	}

	// 没有查找到应用
	if payload.ResultCount == 0 {
		return ErrNotFound
	}
	*app = payload.Results[0]
	return nil
}

// playstore 上的应用
type AppPS struct {
	ID string
}

func (app AppPS) Fill() error {
	return nil
}