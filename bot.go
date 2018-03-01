package dingbot

import (
	"encoding/json"
	"fmt"

	"github.com/ShevaXu/golang/web"
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
	cl    web.Client
}

func (b *dingBot) Send(v interface{}) error {
	req, err := web.NewJSONPost(fmt.Sprintf("%s?%s=%s", dingBotURL, paramToken, b.token), v)
	if err != nil {
		return errors.Wrap(err, "make request")
	}

	_, _, body, err := b.cl.Do(req, dingMaxTries)
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
func WithClient(cl web.Client) BotOption {
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
		cl:    web.NewClient(), // use default setting
	}
	for _, op := range ops {
		op(bot)
	}
	return bot
}
