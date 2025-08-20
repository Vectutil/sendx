package dingding

type DingDingConfig struct {
	webhookURL  string
	accessToken string
	secret      string
}

func NewDingDingConfig(webhookURL, accessToken, secret string) *DingDingConfig {
	return &DingDingConfig{
		webhookURL:  webhookURL,
		accessToken: accessToken,
		secret:      secret,
	}
}

type messageConfig struct {
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

type MessageOption func(*messageConfig)

func WithAtMobiles(mobiles []string) MessageOption {
	return func(c *messageConfig) {
		c.At.AtMobiles = mobiles
	}
}

func WithAtUserIds(userIds []string) MessageOption {
	return func(c *messageConfig) {
		c.At.AtUserIds = userIds
	}
}

func WithAtAll() MessageOption {
	return func(c *messageConfig) {
		c.IsAtAll = true
	}
}

func WithTextMessage(content string) MessageOption {
	return func(c *messageConfig) {
		c.MsgType = "text"
		c.Content = content
	}
}

func WithMarkdownMessage(title, text string) MessageOption {
	return func(c *messageConfig) {
		c.MsgType = "markdown"
		c.Title = title
		c.Text = text
	}
}

func WithLinkMessage(title, text, messageUrl, picUrl string) MessageOption {
	return func(c *messageConfig) {
		c.MsgType = "link"
		c.Title = title
		c.Text = text
		c.MessageUrl = messageUrl
		c.PicUrl = picUrl
	}
}

func WithActionCardMessage(title, text, singleTitle, singleURL, btnOrientation string) MessageOption {
	return func(c *messageConfig) {
		c.MsgType = "actionCard"
		c.Title = title
		c.Text = text
		c.SingleTitle = singleTitle
		c.SingleURL = singleURL
		c.BtnOrientation = btnOrientation
	}
}
