package push

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
)

var _ = TextPusher(&ServerChan{})

const TypeServerChan = "serverchan"

type ServerChanConfig struct {
	Type string `json:"type" yaml:"type"`
	Key  string `yaml:"key" json:"key"`
}

type ServerChan struct {
	pushUrl string
	client  *req.Client
}

func NewServerChan(config *ServerChanConfig) TextPusher {
	return &ServerChan{
		pushUrl: fmt.Sprintf("https://sctapi.ftqq.com/%s.send", config.Key),
		client:  NewHttpClient(),
	}
}

func (d *ServerChan) PushText(s string) error {
	err := d.send("", s)
	if err != nil {
		return errors.Wrap(err, "server-chan")
	}
	return nil
}

func (d *ServerChan) PushMarkdown(title, content string) error {
	err := d.send(title, content)
	if err != nil {
		return errors.Wrap(err, "server-chan")
	}
	return nil
}

func (d *ServerChan) send(text string, desp string) error {
	body := map[string]string{
		"text": text,
		"desp": desp,
	}
	_, err := d.client.R().SetBodyJsonMarshal(body).Post(d.pushUrl)
	return err
}
