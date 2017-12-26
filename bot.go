package dingbot

import (
	"encoding/json"
	"fmt"
	"net/http"

	utils "github.com/ShevaXu/web-utils"
	"github.com/pkg/errors"
)

// internal constants
const (
	dingBotURL   = "https://oapi.dingtalk.com/robot/send"
	paramToken   = "access_token"
	dingMaxTries = 5
	dingCodeOk   = 0
)

// SenderBot sends various message types to
// the Dingtalk webhook.
type SenderBot interface {
	Send(interface{}) error
}

// dingBot is the default implementation of SenderBot.
type dingBot struct {
	token string
	cl    utils.HTTPClient
}

func (b *dingBot) Send(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "JSON marshal content")
	}

	_, _, body, err := b.cl.DoRequest("POST", fmt.Sprintf("%s?%s=%s", dingBotURL, paramToken, b.token), data, dingMaxTries, func(req *http.Request) {
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
	})
	if err != nil {
		return errors.Wrap(err, "dingbot request error")
	}

	var resp DingResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return errors.Wrap(err, "JSON unmarshal response")
	}
	if resp.Errcode != dingCodeOk {
		return errors.New("dingbot send error: " + resp.Errmsg)
	}

	return nil
}

// BotOption follows the functional option pattern.
type BotOption func(*dingBot)

// WithClient makes the bot use a custom http client.
func WithClient(cl utils.HTTPClient) BotOption {
	return func(b *dingBot) {
		if cl != nil {
			b.cl = cl
		}
	}
}

// NewDingBot returns a functional SenderBot with the provided token.
func NewDingBot(token string, ops ...BotOption) SenderBot {
	bot := &dingBot{
		token: token,
		cl:    utils.StdClient(),
	}
	for _, op := range ops {
		op(bot)
	}
	return bot
}
