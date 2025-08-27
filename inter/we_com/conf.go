package we_com

import (
	"encoding/json"
	"sendx/inter"
)

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
	MsgType             string                 `json:"msgtype"` // 消息类型
	Content             map[string]interface{} `json:"-"`
	MentionedList       []string               `json:"-"` // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
	MentionedMobileList []string               `json:"-"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
	IsAtAll             bool                   `json:"-"`
	FileType            string                 `json:"-"` // 文件类型，分别有语音(voice)和普通文件(file)
}

func (m MessageConfig) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"msgtype": m.MsgType,
	}
	data[m.MsgType] = m.Content
	return json.Marshal(data)
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
		c.IsAtAll = true
	}
}

func TextMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: "text",
		Content: map[string]interface{}{
			"content": content,
		},
	}
}

func MarkdownMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: "markdown",
		Content: map[string]interface{}{
			"content": content,
		},
	}
}

func MarkdownV2Message(content string) MessageConfig {
	return MessageConfig{
		MsgType: "markdown_v2",
		Content: map[string]interface{}{
			"markdown_v2": content,
		},
	}
}

func ImageMessage(base64, md5 string) MessageConfig {
	return MessageConfig{
		MsgType: "image",
		Content: map[string]interface{}{
			"base64": base64,
			"md5":    md5,
		},
	}
}

type Articles struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	PicURL      string `json:"picurl"`
}

func NewsMessage(articles []Articles) MessageConfig {
	return MessageConfig{
		MsgType: "news",
		Content: map[string]interface{}{
			"articles": func() []map[string]interface{} {
				var _articles = make([]map[string]interface{}, 0)
				for _, item := range articles {
					_articles = append(_articles, map[string]interface{}{
						"title":       item.Title,
						"url":         item.URL,
						"description": item.Description,
						"picurl":      item.PicURL,
					})
				}
				return _articles
			}(),
		},
	}
}

func FileMessage(mediaID string) MessageConfig {
	return MessageConfig{
		MsgType: "file",
		Content: map[string]interface{}{
			"media_id": mediaID,
		},
	}
}

func VoiceMessage(mediaID string) MessageConfig {
	return MessageConfig{
		MsgType: "voice",
		Content: map[string]interface{}{
			"media_id": mediaID,
		},
	}
}

func TemplateCardMessage(card TextNoticeCard) MessageConfig {
	return MessageConfig{
		MsgType: "template_card",
		Content: map[string]interface{}{
			"msgtype":   "template_card",
			"card_type": card.CardType,
			"source": map[string]interface{}{
				"icon_url":   card.Source.IconURL,
				"desc":       card.Source.Desc,
				"desc_color": card.Source.DescColor,
			},
			"main_title": map[string]interface{}{
				"title": card.MainTitle.Title,
				"desc":  card.MainTitle.Desc,
			},
			"emphasis_content": map[string]interface{}{
				"title": card.EmphasisContent.Title,
				"desc":  card.EmphasisContent.Desc,
			},
			"quote_area": map[string]interface{}{
				"type":       card.QuoteArea.Type,
				"url":        card.QuoteArea.URL,
				"appid":      card.QuoteArea.AppID,
				"pagepath":   card.QuoteArea.PagePath,
				"title":      card.QuoteArea.Title,
				"quote_text": card.QuoteArea.QuoteText,
			},
			"sub_title_text": card.SubTitleText,
			"horizontal_content_list": func() []map[string]interface{} {
				var items []map[string]interface{}
				for _, h := range card.HorizontalContentList {
					item := map[string]interface{}{
						"keyname": h.KeyName,
						"value":   h.Value,
					}
					if h.Type != 0 {
						item["type"] = h.Type
					}
					if h.URL != "" {
						item["url"] = h.URL
					}
					if h.MediaID != "" {
						item["media_id"] = h.MediaID
					}
					items = append(items, item)
				}
				return items
			}(),
			"jump_list": func() []map[string]interface{} {
				var jumps []map[string]interface{}
				for _, j := range card.JumpList {
					jump := map[string]interface{}{
						"type":  j.Type,
						"title": j.Title,
					}
					if j.URL != "" {
						jump["url"] = j.URL
					}
					if j.AppID != "" {
						jump["appid"] = j.AppID
					}
					if j.PagePath != "" {
						jump["pagepath"] = j.PagePath
					}
					jumps = append(jumps, jump)
				}
				return jumps
			}(),
			"card_action": map[string]interface{}{
				"type":     card.CardAction.Type,
				"url":      card.CardAction.URL,
				"appid":    card.CardAction.AppID,
				"pagepath": card.CardAction.PagePath,
			},
		},
	}
}

func FilePush(fType string) MessageConfig {
	return MessageConfig{
		MsgType:  "file_upload",
		FileType: fType,
	}
}
