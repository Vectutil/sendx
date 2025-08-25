package dingding

import "sendx/inter"

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
	IsAtAll bool `json:"isAtAll,omitempty"`
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
