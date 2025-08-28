package feishu

func TextMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: "text",
		Content: FeishuWebhookRequestContent{
			Text: content,
		},
	}
}
func InteractiveMessage(content []Element) MessageConfig {
	return MessageConfig{
		MsgType: MsgTypeInteractive,
		Interact: FeishuWebhookRequestInteract{
			Elements: content,
		},
	}
}
