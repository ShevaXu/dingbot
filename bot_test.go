package dingbot

import (
	"testing"
)

const testToken = "YOUR_TOKEN_HERE"

func TestDingBot_Send(t *testing.T) {
	bot := NewDingBot(testToken)
	msg := DingMessage{
		Msgtype: MsgTypeText,
		Text: TextMsg{
			Content: "hello world 3",
		},
	}
	t.Log(bot.Send(msg))
}
