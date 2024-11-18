package push

import (
	"context"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/format"
	"maunium.net/go/mautrix/id"
)

var _ = TextPusher(&Matrix{})

type MatrixConfig struct {
	Type        string `json:"type" yaml:"type"`
	URL         string `yaml:"url" json:"url"`
	Username    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
	Room        string `yaml:"room" json:"room"`
	UserID      string `yaml:"user_id" json:"user_id"`
	AccessToken string `yaml:"access_token" json:"access_token"`
}

type Matrix struct {
	url    string
	client *mautrix.Client
	roomID string
}

type MatrixData struct {
	Text string `json:"text"`
}

func NewMatrix(config *MatrixConfig) TextPusher {
	client, err := mautrix.NewClient(config.URL, "", "")
	if err != nil {
		return nil
	}
	if config.AccessToken != "" {
		client.SetCredentials(id.UserID(config.UserID), config.AccessToken)
	} else {
		_, err := client.Login(context.Background(), &mautrix.ReqLogin{
			Type:             "m.login.password",
			StoreCredentials: true,
			Password:         config.Password,
			Identifier: mautrix.UserIdentifier{
				User: config.Username,
				Type: "m.id.user",
			},
		})
		if err != nil {
			return nil
		}
	}

	resp, err := client.JoinRoom(context.Background(), config.Room, "", nil)
	if err != nil {
		return nil
	}
	return &Matrix{
		url:    config.URL,
		client: client,
		roomID: string(resp.RoomID),
	}
}

func (m *Matrix) PushText(s string) error {
	content := format.RenderMarkdown(s, true, false)
	content.MsgType = event.MsgNotice
	_, err := m.client.SendMessageEvent(context.Background(), id.RoomID(m.roomID), event.EventMessage, content)
	if err != nil {
		return err
	}
	return nil
}

func (m *Matrix) PushMarkdown(title, content string) error {
	return m.PushText(content)
}
