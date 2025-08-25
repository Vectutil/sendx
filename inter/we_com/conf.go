package we_com

import "sendx/inter"

type WeComConfig struct {
	webhookURL string
	key        string
}

func NewWeComConfig(key string) *WeComConfig {
	return &WeComConfig{
		webhookURL: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
		key:        key,
	}
}

type MessageConfig struct {
	inter.DefaultSendConf
	MsgType             string            `json:"msgtype"`                         // 消息类型
	Content             map[string]string `json:"text,omitempty"`                  // 消息内容 必须是utf8编码
	MentionedList       []string          `json:"mentioned_list,omitempty"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
	MentionedMobileList []string          `json:"mentioned_mobile_list,omitempty"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
	Base64              string            `json:"base64,omitempty"`                // 图片内容的base64编码 注：图片（base64编码前）最大不能超过2M，支持JPG,PNG格式
	Md5                 string            `json:"md5,omitempty"`                   // 图片内容（base64编码前）的md5值
	Articles            string            `json:"articles,omitempty"`              // 图文消息，一个图文消息支持1到8条图文
	Title               string            `json:"title,omitempty"`                 // 标题，不超过128个字节，超过会自动截断
	Description         string            `json:"description,omitempty"`           // 描述，不超过512个字节，超过会自动截断
	Url                 string            `json:"url,omitempty"`                   // 点击后跳转的链接。
	MediaID             string            `json:"media_id,omitempty"`              // 文件id，通过下文的文件上传接口获取
	PicUrl              string            `json:"picUrl,omitempty"`                // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
	IsAtAll             bool              `json:"isAtAll,omitempty"`
	TemplateCard        *TextNoticeCard   `json:"template_card,omitempty"` // 具体的模版卡片参数
	FileType            string            `json:"type,omitempty"`          // 文件类型，分别有语音(voice)和普通文件(file)
	Key                 string            `json:"key,omitempty"`
}

// TextNoticeCard 文本通知模版卡片
type TextNoticeCard struct {
	CardType              string              `json:"card_type"`                         // 是 模版卡片的模版类型，文本通知模版卡片的类型为text_notice
	Source                *CardSource         `json:"source,omitempty"`                  // 否 卡片来源样式信息，不需要来源样式可不填写
	MainTitle             CardMainTitle       `json:"main_title"`                        // 是 模版卡片的主要内容，包括一级标题和标题辅助信息
	EmphasisContent       *EmphasisContent    `json:"emphasis_content,omitempty"`        // 否 关键数据样式
	QuoteArea             *QuoteArea          `json:"quote_area,omitempty"`              // 否 引用文献样式，建议不与关键数据共用
	SubTitleText          string              `json:"sub_title_text,omitempty"`          // 否 二级普通文本，建议不超过112个字
	HorizontalContentList []HorizontalContent `json:"horizontal_content_list,omitempty"` // 否 二级标题+文本列表，列表长度不超过6
	JumpList              []JumpItem          `json:"jump_list,omitempty"`               // 否 跳转指引样式的列表，列表长度不超过3
	CardAction            CardAction          `json:"card_action"`                       // 是 整体卡片的点击跳转事件
}

// CardSource 卡片来源样式信息
type CardSource struct {
	IconURL   string `json:"icon_url,omitempty"`   // 否 来源图片的url
	Desc      string `json:"desc,omitempty"`       // 否 来源图片的描述，建议不超过13个字
	DescColor int    `json:"desc_color,omitempty"` // 否 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
}

// CardMainTitle 模版卡片的主要内容
type CardMainTitle struct {
	Title string `json:"title,omitempty"` // 否 一级标题，建议不超过26个字
	Desc  string `json:"desc,omitempty"`  // 否 标题辅助信息，建议不超过30个字
}

// EmphasisContent 关键数据样式
type EmphasisContent struct {
	Title string `json:"title,omitempty"` // 否 关键数据样式的数据内容，建议不超过10个字
	Desc  string `json:"desc,omitempty"`  // 否 关键数据样式的数据描述内容，建议不超过15个字
}

// QuoteArea 引用文献样式
type QuoteArea struct {
	Type      int    `json:"type,omitempty"`       // 否 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	URL       string `json:"url,omitempty"`        // 否 点击跳转的url，quote_area.type是1时必填
	AppID     string `json:"appid,omitempty"`      // 否 点击跳转的小程序的appid，quote_area.type是2时必填
	PagePath  string `json:"pagepath,omitempty"`   // 否 点击跳转的小程序的pagepath，quote_area.type是2时选填
	Title     string `json:"title,omitempty"`      // 否 引用文献样式的标题
	QuoteText string `json:"quote_text,omitempty"` // 否 引用文献样式的引用文案
}

// HorizontalContent 二级标题+文本列表项
type HorizontalContent struct {
	Type    int    `json:"type,omitempty"`     // 否 模版卡片的二级标题信息内容支持的类型，1是url，2是文件附件，3 代表点击跳转成员详情
	KeyName string `json:"keyname"`            // 是 二级标题，建议不超过5个字
	Value   string `json:"value,omitempty"`    // 否 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过26个字
	URL     string `json:"url,omitempty"`      // 否 链接跳转的url，horizontal_content_list.type是1时必填
	MediaID string `json:"media_id,omitempty"` // 否 附件的media_id，horizontal_content_list.type是2时必填
	UserID  string `json:"userid,omitempty"`   // 否 成员详情的userid，horizontal_content_list.type是3时必填
}

// JumpItem 跳转指引样式的列表项
type JumpItem struct {
	Type     int    `json:"type,omitempty"`     // 否 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
	Title    string `json:"title"`              // 是 跳转链接样式的文案内容，建议不超过13个字
	URL      string `json:"url,omitempty"`      // 否 跳转链接的url，jump_list.type是1时必填
	AppID    string `json:"appid,omitempty"`    // 否 跳转链接的小程序的appid，jump_list.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 否 跳转链接的小程序的pagepath，jump_list.type是2时选填
}

// CardAction 整体卡片的点击跳转事件
type CardAction struct {
	Type     int    `json:"type"`               // 是 卡片跳转类型，1 代表跳转url，2 代表打开小程序
	URL      string `json:"url,omitempty"`      // 否 跳转事件的url，card_action.type是1时必填
	AppID    string `json:"appid,omitempty"`    // 否 跳转事件的小程序的appid，card_action.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 否 跳转事件的小程序的pagepath，card_action.type是2时选填
}
type MessageOption func(*MessageConfig)

func WithAtMentionedList(mentionedList []string) MessageOption {
	return func(c *MessageConfig) {
		c.MentionedList = mentionedList
	}
}

func WithAtMentionedMobileList(mentionedMobileList []string) MessageOption {
	return func(c *MessageConfig) {
		c.MentionedMobileList = mentionedMobileList
	}
}

func WithAtAll() MessageOption {
	return func(c *MessageConfig) {
		c.MentionedMobileList = append(c.MentionedMobileList, "@all")
		c.MentionedList = append(c.MentionedList, "@all")
		c.IsAtAll = true
	}
}

func TextMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: "text",
		Content: map[string]string{
			"content": content,
		},
	}
}

func MarkdownMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: "markdown",
		Content: map[string]string{
			"markdown": content,
		},
	}
}

func MarkdownV2Message(content string) MessageConfig {
	return MessageConfig{
		MsgType: "markdown_v2",
		Content: map[string]string{
			"markdown_v2": content,
		},
	}
}

func ImageMessage(base64, md5 string) MessageConfig {
	return MessageConfig{
		MsgType: "image",
		Base64:  base64,
		Md5:     md5,
	}
}

func NewsMessage(articles, title, url, description, picUrl string) MessageConfig {
	return MessageConfig{
		MsgType:     "news",
		Articles:    articles,
		Title:       title,
		Url:         url,
		Description: description,
		PicUrl:      picUrl,
	}
}

func FileMessage(mediaID string) MessageConfig {
	return MessageConfig{
		MsgType: "file",
		MediaID: mediaID,
	}
}

func VoiceMessage(mediaID string) MessageConfig {
	return MessageConfig{
		MsgType: "voice",
		MediaID: mediaID,
	}
}

func TemplateCardMessage(templateCard *TextNoticeCard) MessageConfig {
	return MessageConfig{
		MsgType:      "template_card",
		TemplateCard: templateCard,
	}
}

func FilePush(fType string) MessageConfig {
	return MessageConfig{
		MsgType:  "file_upload",
		FileType: fType,
	}
}
