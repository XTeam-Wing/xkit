package push

type RawMessage struct {
	Content any    `json:"content"`
	Type    string `json:"type"`
}
