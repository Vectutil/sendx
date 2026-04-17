# Sendx 消息推送库

![Go Version](https://img.shields.io/badge/go-1.24-blue)
![License](https://img.shields.io/badge/license-MIT-green)

企业微信/钉钉/飞书消息推送 SDK，提供多种消息类型支持，包含文本、 Markdown、图文、文件等格式。

## 功能特性

✅ 多平台支持：企业微信/钉钉/飞书
✅ 多种消息类型：文本/卡片/文件/图文  
✅ 安全认证：支持签名校验  
✅ 媒体管理：文件上传API  
✅ 消息模板：预置常用消息模板  

## 安装

```sh
  go get -u github.com/Vectutil/sendx
```

## 快速开始

### 企业微信示例

```go
import (
    "context"
    "github.com/Vectutil/sendx/inter/we_com"
)

wc := we_com.NewWeComConfig("YOUR_API_KEY")

// 发送文本消息
err := wc.SendMsg(context.Background(),
    we_com.TextMessage("系统告警通知"),
    we_com.WithAtAll(),
)
```

### 钉钉示例

```go
import (
    "context"
    "github.com/Vectutil/sendx/inter/dingding"
)

dd := dingding.NewDingDingConfig(
    "YOUR_ACCESS_TOKEN",
    "YOUR_SECRET",
)

// 发送Markdown消息
err := dd.SendMsg(context.Background(),
    dingding.MarkdownMessage("服务器状态", "**CPU使用率**: 98%")
)
```

### 飞书示例

```go
import (
    "context"
    "github.com/Vectutil/sendx/inter/feishu"
)

fs := feishu.NewFeishuConfig("YOUR_WEBHOOK_KEY")

err := fs.SendMsg(context.Background(),
    feishu.TextMessage("飞书推送测试")
)
```

## 支持的消息类型

- `TextMessage` 文本消息
- `MarkdownMessage` Markdown 格式
- `FileMessage` 文件消息
- `TemplateCardMessage` 模板卡片
- `MarkdownV2Message` MarkdownV2 格式

## 测试

```sh
  go test ./...
```

## 贡献指南

1. Fork 仓库并创建分支
2. 编写并提交测试用例
3. 发起 Pull Request

## 许可证

[MIT License](/LICENSE)
