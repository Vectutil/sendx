package we_com

import (
	"context"
	"testing"
)

func TestTextSendMsg(t *testing.T) {
	ddc := NewWeComConfig("a1bb7d9d-8380-4182-8e26-6670d2ef60e6")
	err := ddc.SendMsg(context.Background(), TextMessage("hong i love you"), WithAtAll())
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMsg(t *testing.T) {
	ddc := NewWeComConfig("a1bb7d9d-8380-4182-8e26-6670d2ef60e6")
	err := ddc.SendMsg(context.Background(), MarkdownMessage("# 一、标题\n## 二级标题\n### 三级标题\n# 二、字体\n*斜体*\n\n**加粗**\n# 三、列表 \n- 无序列表 1 \n- 无序列表 2\n  - 无序列表 2.1\n  - 无序列表 2.2\n1. 有序列表 1\n2. 有序列表 2\n# 四、引用\n> 一级引用\n>>二级引用\n>>>三级引用\n# 五、链接\n[这是一个链接](https:work.weixin.qq.com/api/doc)\n![](https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png)\n# 六、分割线\n\n---\n# 七、代码\n`这是行内代码`\n```\n这是独立代码块\n```\n\n# 八、表格\n| 姓名 | 文化衫尺寸 | 收货地址 |\n| :----- | :----: | -------: |\n| 张三 | S | 广州 |\n| 李四 | L | 深圳 |\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCardSendMsg(t *testing.T) {
	ddc := NewWeComConfig("a1bb7d9d-8380-4182-8e26-6670d2ef60e6")
	err := ddc.SendMsg(context.Background(), TemplateCardMessage(TextNoticeCard{
		CardType: "text_notice",
		Source: &CardSource{
			IconURL:   "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
			Desc:      "企业微信",
			DescColor: 0,
		},
		MainTitle: CardMainTitle{
			Title: "欢迎使用企业微信",
			Desc:  "您的好友正在邀请您加入企业微信",
		},
		EmphasisContent: &EmphasisContent{
			Title: "100",
			Desc:  "数据含义",
		},
		QuoteArea: &QuoteArea{
			Type:      1,
			URL:       "https://work.weixin.qq.com/?from=openApi",
			AppID:     "APPID",
			PagePath:  "PAGEPATH",
			Title:     "引用文本标题",
			QuoteText: "Jack：企业微信真的很好用~\nBalian：超级好的一款软件！",
		},
		SubTitleText: "下载企业微信还能抢红包！",
		HorizontalContentList: []HorizontalContent{
			{
				KeyName: "邀请人",
				Value:   "张三",
			},
			{
				KeyName: "企微官网",
				Value:   "点击访问",
				Type:    1,
				URL:     "https://work.weixin.qq.com/?from=openApi",
			},
		},
		JumpList: []JumpItem{
			{
				Type:  1,
				URL:   "https://work.weixin.qq.com/?from=openApi",
				Title: "企业微信官网",
			},
		},
		CardAction: CardAction{
			Type:     1,
			URL:      "https://work.weixin.qq.com/?from=openApi",
			AppID:    "APPID",
			PagePath: "PAGEPATH",
		},
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewsSendMsg(t *testing.T) {
	ddc := NewWeComConfig("a1bb7d9d-8380-4182-8e26-6670d2ef60e6")
	err := ddc.SendMsg(context.Background(), NewsMessage([]Articles{
		{
			Title:       "标题",
			URL:         "https://www.baidu.com",
			Description: "描述",
			PicURL:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "标题2",
			URL:         "https://www.baidu1.com",
			Description: "描述2",
			PicURL:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
	}))
	if err != nil {
		t.Fatal(err)
	}
}
