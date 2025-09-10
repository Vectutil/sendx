# 快速开始

本指南将帮助您快速开始使用 Vectutil/sendx 库。

## 环境要求

- Go 1.16 或更高版本
- 使用 Go modules 管理依赖

## 安装

首先，确保您的项目使用 Go modules：

```bash
go mod init your-project-name
```

然后安装 Vectutil/sendx：

```bash
go get github.com/Vectutil/sendx
```

## 基本用法

### 1. 导入包

根据您要使用的平台导入相应的包：

```go
import (
    "context"
    "github.com/Vectutil/sendx/inter/dingding"  // 钉钉
    "github.com/Vectutil/sendx/inter/feishu"    // 飞书
    "github.com/Vectutil/sendx/inter/we_com"    // 企业微信
    "github.com/Vectutil/sendx/inter/telegram"  // Telegram
)
```

### 2. 创建配置

每个平台都需要相应的配置：

```go
// 钉钉配置
dingConfig := dingding.NewDingDingConfig("your_access_token", "your_secret")

// 飞书配置
feishuConfig := feishu.NewFeishuConfig("your_webhook_key")

// 企业微信配置
weComConfig := we_com.NewWeComConfig("your_webhook_key")

// Telegram 配置（开发中）
// telegramConfig := telegram.NewTelegramConfig("your_bot_token", "your_chat_id")
```

### 3. 创建消息

根据不同平台创建不同类型的消息：

```go
// 钉钉文本消息
dingMessage := dingding.TextMessage("Hello, World!")

// 企业微信 Markdown 消息
weComMessage := we_com.MarkdownMessage("## Hello, World!")

// 飞书富文本消息
feishuMessage := feishu.MessageConfig{
    MsgType: feishu.MsgTypePost,
    Content: feishu.FeishuWebhookRequestContent{
        Post: feishu.FeishuWebhookRequestContentPost{
            ZhCn: feishu.ZhCn{
                Title: "标题",
                Content: [][]feishu.ZhCnContent{
                    {
                        {
                            Tag:  "text",
                            Text: "内容",
                        },
                    },
                },
            },
        },
    },
}
```

### 4. 发送消息

```go
// 发送钉钉消息
err := dingConfig.SendMsg(context.Background(), dingMessage)
if err != nil {
    // 处理错误
    panic(err)
}

// 发送企业微信消息
err = weComConfig.SendMsg(context.Background(), weComMessage)
if err != nil {
    // 处理错误
    panic(err)
}

// 发送飞书消息
err = feishuConfig.SendMsg(context.Background(), feishuMessage)
if err != nil {
    // 处理错误
    panic(err)
}
```

## 平台特定配置

### 钉钉

钉钉机器人需要 `access_token` 和 `secret`：

```go
config := dingding.NewDingDingConfig("access_token", "secret")
```

获取方式：
1. 在钉钉群中添加机器人
2. 选择"自定义"机器人类型
3. 记录 webhook URL 中的 access_token
4. 如果启用了安全设置，记录密钥(secret)

### 飞书

飞书机器人需要 webhook key：

```go
config := feishu.NewFeishuConfig("webhook_key")
```

获取方式：
1. 在飞书群中添加机器人
2. 选择"自定义机器人"
3. 记录 webhook URL 中的 key

### 企业微信

企业微信机器人需要 webhook key：

```go
config := we_com.NewWeComConfig("webhook_key")
```

获取方式：
1. 在企业微信群中添加机器人
2. 选择"群机器人"
3. 记录 webhook URL 中的 key

### Telegram

Telegram 机器人需要 bot token 和 chat ID：

```go
// 开发中，API 可能会变化
config := telegram.NewTelegramConfig("bot_token", "chat_id")
```

获取方式：
1. 在 Telegram 中创建 Bot，获取 bot token
2. 将 Bot 添加到群组或发送消息给 Bot 获取 chat ID

## 消息类型

### 文本消息

所有平台都支持简单的文本消息：

```go
// 钉钉
message := dingding.TextMessage("Hello, World!")

// 企业微信
message := we_com.TextMessage("Hello, World!")

// 飞书
message := feishu.MessageConfig{
    MsgType: feishu.MsgTypeText,
    Content: feishu.FeishuWebhookRequestContent{
        Text: "Hello, World!",
    },
}
```

### Markdown 消息

```go
// 企业微信
message := we_com.MarkdownMessage("**Bold** and *italic* text")

// 钉钉
message := dingding.MarkdownMessage("标题", "**Bold** and *italic* text")
```

### 高级消息类型

不同平台支持不同的高级消息类型：

- **钉钉**: Link 消息、ActionCard 消息
- **企业微信**: 图文消息、文件消息、模板卡片消息
- **飞书**: 富文本消息、图片消息、卡片消息

## 错误处理

所有发送方法都会返回 error，建议进行适当的错误处理：

```go
err := config.SendMsg(context.Background(), message)
if err != nil {
    // 记录日志
    log.Printf("发送消息失败: %v", err)
    
    // 根据错误类型进行处理
    // ...
}
```

## 完整示例

```go
package main

import (
    "context"
    "log"
    "github.com/Vectutil/sendx/inter/dingding"
)

func main() {
    // 创建配置
    config := dingding.NewDingDingConfig("your_access_token", "your_secret")
    
    // 创建消息
    message := dingding.TextMessage("Hello, World!")
    
    // 发送消息
    err := config.SendMsg(context.Background(), message)
    if err != nil {
        log.Fatalf("发送消息失败: %v", err)
    }
    
    log.Println("消息发送成功")
}
```

现在您已经了解了基本用法，请查看各平台的详细 API 文档了解更多信息。