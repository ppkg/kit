package kit

import (
	"encoding/json"

	"github.com/gogrpc/glog"
	kitHttp "github.com/gogrpc/kit/http"
	"github.com/limitedlee/microservice/common/config"
)

func Alert(content string) {
	wecomAlert(content)
}

type WecomAlert struct {
	MsgType string            `json:"msgtype"`
	Text    WecomAlertContent `json:"text"`
}

type WecomAlertContent struct {
	Content string `json:"content"`
}

type WecomAlertResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func wecomAlert(content string) {
	httpHandle := kitHttp.NewHttpHandle(config.GetString("wecom-webhook"))

	params := JsonEncodeByte(WecomAlert{
		MsgType: "text",
		Text: WecomAlertContent{
			Content: content,
		},
	})
	resByte, err := httpHandle.Post(params)
	if err != nil {
		glog.Error("发送企业微信通知消息失败，", err, "，参数：", string(params))
		return
	}

	res := new(WecomAlertResponse)
	if err = json.Unmarshal(resByte, res); err != nil {
		glog.Error("解析企业微信返回失败，", err, "，参数：", string(params))
		return
	}
	if res.Errcode != 0 {
		glog.Error("解析企业微信返回失败，", string(resByte), "，参数：", string(params))
		return
	}
}
