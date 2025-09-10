# 介绍

Vectutil/sendx 是一个用于向各种平台发送消息的 Go 工具库。它提供了统一的接口来发送消息到不同的平台，包括钉钉、飞书、企业微信和 Telegram。

## 特性

- **统一接口**: 所有平台都使用相似的 API 接口
- **易于使用**: 简单的配置和消息发送方式
- **多平台支持**: 支持主流的即时通讯平台
- **灵活配置**: 支持多种消息类型和配置选项

## 支持的平台

- [钉钉](../api/dingding.md) - 企业级通讯平台
- [飞书](../api/lark.md) - 字节跳动推出的协作平台
- [企业微信](../api/wechat-work.md) - 腾讯企业级通讯工具
- [Telegram](../api/telegram.md) - 跨平台即时通讯软件

## 安装

使用 Go modules 安装：

```bash
go get github.com/Vectutil/sendx
```

## 快速预览

### 钉钉消息发送示例

```go
import "github.com/Vectutil/sendx/inter/dingding"

// 创建钉钉配置
config := dingding.NewDingDingConfig("your_access_token", "your_secret")

// 发送文本消息
message := dingding.TextMessage("Hello, World!")
err := config.SendMsg(context.Background(), message)
```

### 企业微信消息发送示例

```go
import "github.com/Vectutil/sendx/inter/we_com"

// 创建企业微信配置
config := we_com.NewWeComConfig("your_webhook_key")

// 发送 Markdown 消息
message := we_com.MarkdownMessage("## Hello, World!")
err := config.SendMsg(context.Background(), message)
```

### 飞书消息发送示例

```go
import "github.com/Vectutil/sendx/inter/feishu"

// 创建飞书配置
config := feishu.NewFeishuConfig("your_webhook_key")

// 发送文本消息
message := feishu.MessageConfig{
    MsgType: feishu.MsgTypeText,
    Content: feishu.FeishuWebhookRequestContent{
        Text: "Hello, World!",
    },
}
err := config.SendMsg(context.Background(), message)
```

## 核心概念

所有平台的消息发送都遵循相似的模式：

1. **创建配置**: 每个平台都需要特定的配置信息，如 webhook URL 或访问令牌
2. **创建消息**: 根据需要创建不同类型的消息对象
3. **发送消息**: 使用配置对象的 `SendMsg` 方法发送消息

接下来，请查看[快速开始](./getting-started.md)指南了解更多详细信息。