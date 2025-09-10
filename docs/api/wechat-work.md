# 企业微信 API

## 配置

```go
import "github.com/Vectutil/sendx/inter/we_com"

config := we_com.NewWeComConfig(key)
```

### 参数

- `key`: 企业微信群机器人的 webhook key

## 消息类型

### 文本消息

```go
message := we_com.TextMessage("Hello, World!")
```

### Markdown 消息

```go
message := we_com.MarkdownMessage("## Markdown 内容")
```

### Markdown V2 消息

```go
message := we_com.MarkdownV2Message("**Markdown V2 内容**")
```

### 图片消息

```go
message := we_com.ImageMessage(base64, md5)
```

### 图文消息

```go
articles := []we_com.Articles{
    {
        Title:       "标题",
        URL:         "https://example.com",
        Description: "描述",
        PicURL:      "https://example.com/image.png",
    },
}
message := we_com.NewsMessage(articles)
```

### 文件消息

```go
message := we_com.FileMessage(mediaID)
```

### 语音消息

```go
message := we_com.VoiceMessage(mediaID)
```

### 模板卡片消息

```go
card := we_com.TextNoticeCard{
    CardType: "text_notice",
    MainTitle: we_com.CardMainTitle{
        Title: "主标题",
        Desc:  "描述",
    },
    // 其他字段...
}
message := we_com.TemplateCardMessage(card)
```

## 消息选项

### @指定成员

```go
message := we_com.TextMessage("Hello, World!")
we_com.WithAtMentionedList([]string{"user123"})
```

### @指定手机号

```go
message := we_com.TextMessage("Hello, World!")
we_com.WithAtMentionedMobileList([]string{"13800138000"})
```

### @所有人

```go
message := we_com.TextMessage("Hello, World!")
we_com.WithAtAll()
```

## 发送消息

```go
err := config.SendMsg(context.Background(), message, options...)
```

### 参数

- `ctx`: 上下文
- `sendInter`: 消息对象
- `opts`: 消息选项（可选）

## 上传媒体文件

```go
mediaID, err := config.UploadMedia(filePath, fileType)
```

### 参数

- `filePath`: 文件路径
- `fileType`: 文件类型（"file" 或 "voice"）

## 完整示例

```go
package main

import (
    "context"
    "fmt"
    "github.com/Vectutil/sendx/inter/we_com"
)

func main() {
    // 创建配置
    config := we_com.NewWeComConfig("your_webhook_key")
    
    // 创建文本消息
    message := we_com.TextMessage("Hello, World!")
    
    // 发送消息并@所有人
    err := config.SendMsg(
        context.Background(), 
        message, 
        we_com.WithAtAll(),
    )
    
    if err != nil {
        fmt.Printf("发送失败: %v\n", err)
    } else {
        fmt.Println("发送成功")
    }
}
```