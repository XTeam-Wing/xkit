package push

import (
	"github.com/pkg/errors"
	wxworkbot "github.com/vimsucks/wxwork-bot-go"
)

var _ = TextPusher(&WechatWork{})

const TypeWechatWork = "wechatwork"

type WechatWorkConfig struct {
	Type string `json:"type" yaml:"type"`
	Key  string `yaml:"key" json:"key"`
}

type WechatWork struct {
	client *wxworkbot.WxWorkBot
}

func NewWechatWork(config *WechatWorkConfig) TextPusher {
	return &WechatWork{
		client: wxworkbot.New(config.Key),
	}
}

func (d *WechatWork) PushText(s string) error {
	// fixme: wxworkbot 不支持 text 类型
	msg := wxworkbot.Markdown{Content: s}
	err := d.client.Send(msg)
	if err != nil {
		return errors.Wrap(err, "wechat-work")
	}
	return nil
}

func (d *WechatWork) PushMarkdown(title, content string) error {
	msg := wxworkbot.Markdown{Content: content}
	err := d.client.Send(msg)
	if err != nil {
		return errors.Wrap(err, "wechat-work")
	}
	return nil
}
