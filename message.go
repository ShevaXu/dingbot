package dingbot

// exported message types
const (
	MsgTypeText       = "text"
	MsgTypeLink       = "link"
	MsgTypeMd         = "markdown"
	MsgTypeActionCard = "actionCard"
	MsgTypeFeedCard   = "feedCard"
)

type TextMsg struct {
	Content string `json:"content"`
}

type LinkMsg struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl,omitempty"`
	MessageURL string `json:"messageUrl"`
}

type MdMsg struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type ActionCardMsg struct {
	Text           string `json:"text"`
	Title          string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`     // "0" or "1"
	BtnOrientation string `json:"btnOrientation"` // "0" or "1"
	// single card
	SingleTitle string `json:"singleTitle"`
	SingleURL   string `json:"singleURL"`
	// independent cards
	Btns []struct {
		Title     string `json:"title"`
		ActionURL string `json:"actionURL"`
	} `json:"btns"`
}

type FeedLink struct {
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}

type FeedCardMsg struct {
	Links []FeedLink `json:"links"`
}

type AtOption struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type DingMessage struct {
	Msgtype    string        `json:"msgtype"`
	Text       TextMsg       `json:"text,omitempty"`
	Link       LinkMsg       `json:"link,omitempty"`
	Markdown   MdMsg         `json:"markdown,omitempty"`
	ActionCard ActionCardMsg `json:"actionCard,omitempty"`
	FeedCard   FeedCardMsg   `json:"feedCard,omitempty"`
	At         AtOption      `json:"at,omitempty"`
}

type DingResponse struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
}
