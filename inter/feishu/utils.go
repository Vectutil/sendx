package feishu

func TextMessage(content string) MessageConfig {
	return MessageConfig{
		MsgType: MsgTypeText,
		Content: FeishuWebhookRequestContent{
			Text: content,
		},
	}
}

func PostMessage(title string, content [][]ZhCnContent) MessageConfig {
	return MessageConfig{
		MsgType: MsgTypePost,
		Content: FeishuWebhookRequestContent{
			Post: FeishuWebhookRequestContentPost{
				ZhCn: ZhCn{
					Title:   title,
					Content: content,
				},
			},
		},
	}
}

func ImageMessage(imageKey string) MessageConfig {
	return MessageConfig{
		MsgType:  MsgTypeImage,
		ImageKey: imageKey,
	}
}

func ShareChatMessage(chatID string) MessageConfig {
	return MessageConfig{
		MsgType:   MsgTypeShareChat,
		ShareChat: chatID,
	}
}

func ShareUserMessage(userID string) MessageConfig {
	return MessageConfig{
		MsgType:   MsgTypeShareUser,
		ShareUser: userID,
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

func InteractiveCardMessage(card map[string]any) MessageConfig {
	return MessageConfig{
		MsgType: MsgTypeInteractive,
		Card:    card,
	}
}
