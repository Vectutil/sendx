# API 参考

Vectutil/sendx 提供了统一的接口来向不同平台发送消息。

## 核心概念

所有平台都遵循相似的模式：

1. **创建平台配置**: 每个平台都需要特定的配置信息
2. **创建消息对象**: 根据需要创建不同类型的消息
3. **发送消息**: 使用配置对象的 `SendMsg` 方法发送消息

## 通用接口

所有平台都实现了 `SendInter` 接口：

```go
type SendInter interface {
    SendHttpRequest(url string, param any) (map[string]interface{}, error)
}
```

## 平台 API

### [钉钉 API](./dingding.md)

钉钉是阿里巴巴集团推出的即时通讯和协作平台，广泛用于企业内部沟通。

主要特性：
- 支持多种消息类型（文本、Markdown、链接、ActionCard）
- 支持 @ 功能
- 安全签名验证

### [飞书 API](./lark.md)

飞书是字节跳动推出的企业协作平台，集即时通讯、日历、云文档等功能于一体。

主要特性：
- 支持富文本消息
- 支持卡片消息
- 支持多种媒体类型

### [企业微信 API](./wechat-work.md)

企业微信是腾讯推出的办公通讯工具，与微信互通，方便企业内外部沟通。

主要特性：
- 支持多种消息类型
- 支持 @ 功能
- 支持文件上传
- 支持模板卡片消息

### [Telegram API](./telegram.md)

Telegram 是一个跨平台的即时通讯软件，注重安全性和速度。

主要特性：
- 开发中，API 可能会变化
- 支持文本和媒体消息
- 机器人支持

## 公共方法

所有平台的配置对象都实现了 `SendMsg` 方法：

```go
func (config *Config) SendMsg(ctx context.Context, sendInter SendInter, opts ...MessageOption) error
```

### 参数说明

- `ctx`: Go 的上下文，用于控制请求的生命周期
- `sendInter`: 消息对象，包含要发送的内容
- `opts`: 可选的消息选项，如 @ 功能等

## 消息选项

不同平台支持不同的消息选项：

### @ 功能

大多数平台都支持 @ 功能，用于提醒特定用户：

- 钉钉：`WithAtMobiles`, `WithAtUserIds`, `WithAtAll`
- 企业微信：`WithAtMentionedList`, `WithAtMentionedMobileList`, `WithAtAll`

### 其他选项

根据不同平台的特性，还支持其他选项，如文件上传、按钮配置等。

## 错误处理

所有发送方法都会返回标准的 Go error，可以根据错误类型进行相应的处理：

```go
if err != nil {
    if strings.Contains(err.Error(), "errcode") {
        // 处理平台特定错误
    } else {
        // 处理网络或其他错误
    }
}
```

请查看各平台的详细文档了解具体的 API 使用方法和示例。