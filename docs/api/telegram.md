# Telegram API

::: warning 注意
Telegram 功能目前还在开发中，API 可能会在未来发生变化。
:::

## 配置

```go
import "github.com/Vectutil/sendx/inter/telegram"

config := telegram.NewTelegramConfig(botToken, chatID)
```

### 参数

- `botToken`: Telegram 机器人的访问令牌
- `chatID`: 要发送消息的聊天 ID

## 消息类型

### 文本消息

```go
message := telegram.TextMessage("Hello, World!")
```

### Markdown 消息

```go
message := telegram.MarkdownMessage("*Markdown* _message_")
```

### 图片消息

```go
message := telegram.ImageMessage(photoURL, caption)
```

## 发送消息

```go
err := config.SendMsg(context.Background(), message)
```

### 参数

- `ctx`: 上下文
- `sendInter`: 消息对象

## 完整示例

```go
package main

import (
    "context"
    "fmt"
    "github.com/Vectutil/sendx/inter/telegram"
)

func main() {
    // 创建配置
    config := telegram.NewTelegramConfig("your_bot_token", "your_chat_id")
    
    // 创建文本消息
    message := telegram.TextMessage("Hello, World!")
    
    // 发送消息
    err := config.SendMsg(context.Background(), message)
    
    if err != nil {
        fmt.Printf("发送失败: %v\n", err)
    } else {
        fmt.Println("发送成功")
    }
}
```