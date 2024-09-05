package push

import (
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
)

var _ = RawPusher(&Webhook{})

const TypeWebhook = "webhook"

type WebhookConfig struct {
	Type string `json:"type" yaml:"type"`
	URL  string `yaml:"url" json:"url"`
}

type Webhook struct {
	url    string
	client *req.Client
}

func NewWebhook(config *WebhookConfig) RawPusher {
	return &Webhook{
		url:    config.URL,
		client: NewHttpClient(),
	}
}

func (m *Webhook) PushRaw(r *RawMessage) error {
	_, err := m.client.R().SetBodyJsonMarshal(r).Post(m.url)
	if err != nil {
		return errors.Wrap(err, "webhook")
	}
	return nil
}
