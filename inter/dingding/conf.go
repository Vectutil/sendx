package dingding

import (
	"encoding/json"
	"github.com/Vectutil/sendx/inter"
)

type DingDingConfig struct {
	webhookURL  string
	accessToken string
	secret      string
}

func NewDingDingConfig(accessToken, secret string) *DingDingConfig {
	return &DingDingConfig{
		webhookURL:  "https://oapi.dingtalk.com/robot/send",
		accessToken: accessToken,
		secret:      secret,
	}
}

type MessageConfig struct {
	inter.DefaultSendConf
	MsgType        string `json:"msgtype"`
	Content        string `json:"content,omitempty"`
	Title          string `json:"title,omitempty"`
	Text           string `json:"text,omitempty"`
	MessageUrl     string `json:"messageUrl,omitempty"`
	SingleTitle    string `json:"singleTitle,omitempty"`
	SingleURL      string `json:"singleURL,omitempty"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	PicUrl         string `json:"picUrl,omitempty"`
	At             struct {
		AtMobiles []string `json:"atMobiles,omitempty"`
		AtUserIds []string `json:"atUserIds,omitempty"`
	} `json:"at,omitempty"`
	IsAtAll   bool               `json:"isAtAll,omitempty"`
	Btns      []ActionCardButton `json:"-"`
	FeedLinks []FeedLink         `json:"-"`
}

type ActionCardButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type FeedLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL,omitempty"`
}

func (m MessageConfig) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"msgtype": m.MsgType,
	}

	switch m.MsgType {
	case "text":
		data["text"] = map[string]any{
			"content": m.Content,
		}
	case "markdown":
		data["markdown"] = map[string]any{
			"title": m.Title,
			"text":  m.Text,
		}
	case "link":
		data["link"] = map[string]any{
			"title":      m.Title,
			"text":       m.Text,
			"messageUrl": m.MessageUrl,
			"picUrl":     m.PicUrl,
		}
	case "actionCard":
		actionCard := map[string]any{
			"title":          m.Title,
			"text":           m.Text,
			"btnOrientation": m.BtnOrientation,
		}
		if len(m.Btns) > 0 {
			actionCard["btns"] = m.Btns
		} else {
			actionCard["singleTitle"] = m.SingleTitle
			actionCard["singleURL"] = m.SingleURL
		}
		data["actionCard"] = actionCard
	case "feedCard":
		data["feedCard"] = map[string]any{
			"links": m.FeedLinks,
		}
	}

	if len(m.At.AtMobiles) > 0 || len(m.At.AtUserIds) > 0 || m.IsAtAll {
		data["at"] = map[string]any{
			"atMobiles": m.At.AtMobiles,
			"atUserIds": m.At.AtUserIds,
			"isAtAll":   m.IsAtAll,
		}
	}

	return json.Marshal(data)
}

type MessageOption func(*MessageConfig)

func WithAtMobiles(mobiles []string) MessageOption {
	return func(c *MessageConfig) {
		c.At.AtMobiles = mobiles
	}
}

func WithAtUserIds(userIds []string) MessageOption {
	return func(c *MessageConfig) {
		c.At.AtUserIds = userIds
	}
}

func WithAtAll() MessageOption {
	return func(c *MessageConfig) {
		c.IsAtAll = true
	}
}

func TextMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: "text",
		Content: content,
	}
}

func MarkdownMessage(title, text string) MessageConfig {
	return MessageConfig{
		MsgType: "markdown",
		Title:   title,
		Text:    text,
	}
}

func LinkMessage(title, text, messageUrl, picUrl string) MessageConfig {
	return MessageConfig{
		MsgType:    "link",
		Title:      title,
		Text:       text,
		MessageUrl: messageUrl,
		PicUrl:     picUrl,
	}
}

func ActionCardMessage(title, text, singleTitle, singleURL, btnOrientation string) MessageConfig {
	return MessageConfig{
		MsgType:        "actionCard",
		Title:          title,
		Text:           text,
		SingleTitle:    singleTitle,
		SingleURL:      singleURL,
		BtnOrientation: btnOrientation,
	}
}

func MultiActionCardMessage(title, text, btnOrientation string, btns []ActionCardButton) MessageConfig {
	return MessageConfig{
		MsgType:        "actionCard",
		Title:          title,
		Text:           text,
		BtnOrientation: btnOrientation,
		Btns:           btns,
	}
}

func FeedCardMessage(links []FeedLink) MessageConfig {
	return MessageConfig{
		MsgType:   "feedCard",
		FeedLinks: links,
	}
}
