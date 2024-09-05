package push

import "testing"

func TestDingTalk(t *testing.T) {
	dingPusher := NewDingDing(&DingDingConfig{
		Type:        "dingding",
		AccessToken: "xx",
		SignSecret:  "xx",
	})
	dingPusher.PushMarkdown("a", "a")
}
