# 钉钉 API

## 配置

```go
import "github.com/Vectutil/sendx/inter/dingding"

config := dingding.NewDingDingConfig(accessToken, secret)
```

### 参数

- `accessToken`: 钉钉机器人的访问令牌
- `secret`: 钉钉机器人的签名密钥

## 消息类型

### 文本消息

```go
message := dingding.TextMessage("Hello, World!")
```

### Markdown 消息

```go
message := dingding.MarkdownMessage("标题", "## Markdown 内容")
```

### 链接消息

```go
message := dingding.LinkMessage(
    "标题",
    "内容描述",
    "https://example.com",  // 点击链接地址
    "https://example.com/image.png"  // 图片 URL
)
```

### ActionCard 消息

```go
message := dingding.ActionCardMessage(
    "标题",
    "内容描述",
    "按钮文字",
    "https://example.com",  // 按钮链接
    "0"  // 按钮排列方式，0-按钮竖直排列，1-按钮横向排列
)
```

## 消息选项

### @指定手机号

```go
message := dingding.TextMessage("Hello, World!")
dingding.WithAtMobiles([]string{"13800138000"})
```

### @指定用户ID

```go
message := dingding.TextMessage("Hello, World!")
dingding.WithAtUserIds([]string{"user123"})
```

### @所有人

```go
message := dingding.TextMessage("Hello, World!")
dingding.WithAtAll()
```

## 发送消息

```go
err := config.SendMsg(context.Background(), message, options...)
```

### 参数

- `ctx`: 上下文
- `sendInter`: 消息对象
- `opts`: 消息选项（可选）

## 完整示例

```go
package main

import (
    "context"
    "fmt"
    "github.com/Vectutil/sendx/inter/dingding"
)

func main() {
    // 创建配置
    config := dingding.NewDingDingConfig("your_access_token", "your_secret")
    
    // 创建文本消息
    message := dingding.TextMessage("Hello, World!")
    
    // 发送消息并@所有人
    err := config.SendMsg(
        context.Background(), 
        message, 
        dingding.WithAtAll(),
    )
    
    if err != nil {
        fmt.Printf("发送失败: %v\n", err)
    } else {
        fmt.Println("发送成功")
    }
}
```