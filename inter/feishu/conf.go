package feishu

import "github.com/Vectutil/sendx/inter"

const (
	MsgTypeText        = "text"        // 文本
	MsgTypePost        = "post"        // 富文本
	MsgTypeImage       = "image"       // 图片
	MsgTypeFile        = "file"        // 文件
	MsgTypeAudio       = "audio"       // 音频
	MsgTypeMedia       = "media"       // 视频
	MsgTypesTicker     = "sticker"     // 表情
	MsgTypeInteractive = "interactive" // 卡片
	MsgTypeShareChat   = "share_chat"  // 群聊
	MsgTypeShareUser   = "share_user"  // 个人
	MsgTypeSystem      = "system"      // 系统
)

type FeishuConfig struct {
	webhookURL string
	key        string
}

func NewFeishuConfig(key string) *FeishuConfig {
	return &FeishuConfig{
		webhookURL: "https://open.feishu.cn/open-apis/bot/v2/hook/",
		key:        key,
	}
}

type MessageOption func(*MessageConfig)

type MessageConfig struct {
	inter.DefaultSendConf
	MsgType  string                       `json:"msg_type"`
	Content  FeishuWebhookRequestContent  `json:"content"`
	Interact FeishuWebhookRequestInteract `json:"card"`
}

type FeishuWebhookRequestInteract struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Tag  string       `json:"tag"`
	Text ElementsText `json:"text"`
}

type ElementsText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type FeishuWebhookRequestContent struct {
	Text string                          `json:"text"`
	Post FeishuWebhookRequestContentPost `json:"post"`
}

type FeishuWebhookRequestContentPost struct {
	ZhCn ZhCn `json:"zh_cn"`
}

type ZhCn struct {
	Title   string          `json:"title"`
	Content [][]ZhCnContent `json:"content"`
}

type ZhCnContent struct {
	Tag      string `json:"tag"`
	Text     string `json:"text,omitempty"`
	Href     string `json:"href,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	ImageKey string `json:"image_key,omitempty"`
}
