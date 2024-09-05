package push

import (
	"github.com/imroc/req/v3"
	"strings"
)

var _ = TextPusher(&Bark{})

const TypeBark = "bark"

type BarkConfig struct {
	Type string `json:"type" yaml:"type"`
	URL  string `yaml:"url" json:"url"`
}

type Bark struct {
	url       string
	deviceKey string
	client    *req.Client
}

type BarkData struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	DeviceKey string `json:"device_key"`
	Badge     int    `json:"badge"`
	Group     string `json:"group"`
	Sound     string `json:"sound"`
	Icon      string `json:"icon"`
	Url       string `json:"url"`
}

func NewBark(config *BarkConfig) TextPusher {
	deviceKeys := strings.Split(config.URL, "/")
	deviceKey := deviceKeys[len(deviceKeys)-1]
	u := strings.Replace(config.URL, deviceKey, "push", -1)

	return &Bark{
		url:       u,
		deviceKey: deviceKey,
		client:    NewHttpClient(),
	}
}

func (m *Bark) PushText(s string) error {
	params := &BarkData{
		Title:     "xkit",
		Body:      s,
		DeviceKey: m.deviceKey,
		Badge:     1,
		Group:     "xkit",
		Sound:     "alert",
		Icon:      "",
		Url:       "",
	}

	_, err := m.postJSON(m.url, params)
	if err != nil {
		return err
	}

	return nil
}

func (m *Bark) PushMarkdown(title, content string) error {
	// m.log.Infof("sending markdown %s", title)

	// params := &Bark{
	// 	Title:   title,
	// 	Content: content,
	// 	Type:    "markdown",
	// }
	// resp, err := m.postJSON(m.url, params)
	// if err != nil {
	// 	return err
	// }
	// m.log.Infof("markdown response from server: %s", string(resp))
	return m.PushText(content)
}

func (m *Bark) postJSON(url string, params *BarkData) ([]byte, error) {
	resp, err := m.client.R().SetBodyJsonMarshal(params).Post(url)
	if err != nil {
		return nil, err
	}
	return resp.ToBytes()
}
