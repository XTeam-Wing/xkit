package tests

import (
	timeutils "github.com/XTeam-Wing/xkit/kits/time"
	"testing"
	"time"
)

func TestTimeDiff(t *testing.T) {
	startTime := time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(3 * time.Second)
	endTime := time.Now().Format("2006-01-02 15:04:05")
	t.Log("Time diff:", timeutils.GetTimeDiff(startTime, endTime))
}
