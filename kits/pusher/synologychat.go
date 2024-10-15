package push

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
)

var _ = TextPusher(&SynologyChat{})

type SynologyChatConfig struct {
	Type string `json:"type" yaml:"type"`
	URL  string `yaml:"url" json:"url"`
}

type SynologyChat struct {
	url    string
	client *req.Client
}

type SynologyChatData struct {
	Text string `json:"text"`
}

func NewSynologyChat(config *SynologyChatConfig) TextPusher {

	return &SynologyChat{
		url:    config.URL,
		client: NewHttpClient(),
	}
}

func (m *SynologyChat) PushText(s string) error {
	params := &SynologyChatData{
		Text: s,
	}
	payload, err := json.Marshal(params)
	if err != nil {
		return err
	}
	postData := fmt.Sprintf("payload=%s", payload)
	_, err = m.client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded").SetBody(postData).Post(m.url)
	if err != nil {
		return err
	}

	return nil
}

func (m *SynologyChat) PushMarkdown(title, content string) error {
	return m.PushText(content)
}

func (m *SynologyChat) postJSON(url string, params *SynologyChatData) ([]byte, error) {
	resp, err := m.client.R().SetBodyJsonMarshal(params).Post(url)
	if err != nil {
		return nil, err
	}
	return resp.ToBytes()
}
