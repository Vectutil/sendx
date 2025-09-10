# 飞书 API

## 配置

```go
import "github.com/Vectutil/sendx/inter/feishu"

config := feishu.NewFeishuConfig(key)
```

### 参数

- `key`: 飞书机器人 webhook 的 key

## 消息类型

### 文本消息

```go
message := feishu.MessageConfig{
    MsgType: feishu.MsgTypeText,
    Content: feishu.FeishuWebhookRequestContent{
        Text: "Hello, World!",
    },
}
```

### 富文本消息

```go
message := feishu.MessageConfig{
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

### 图片消息

```go
// TODO: 添加图片消息示例
```

### 文件消息

```go
// TODO: 添加文件消息示例
```

### 卡片消息

```go
// TODO: 添加卡片消息示例
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
    "github.com/Vectutil/sendx/inter/feishu"
)

func main() {
    // 创建配置
    config := feishu.NewFeishuConfig("your_webhook_key")
    
    // 创建文本消息
    message := feishu.MessageConfig{
        MsgType: feishu.MsgTypeText,
        Content: feishu.FeishuWebhookRequestContent{
            Text: "Hello, World!",
        },
    }
    
    // 发送消息
    err := config.SendMsg(context.Background(), message)
    
    if err != nil {
        fmt.Printf("发送失败: %v\n", err)
    } else {
        fmt.Println("发送成功")
    }
}
```