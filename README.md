# Sendx 消息推送库

![Go Version](https://img.shields.io/badge/go-1.24-blue)
![License](https://img.shields.io/badge/license-MIT-green)

企业微信/钉钉消息推送SDK，提供多种消息类型支持，包含文本、Markdown、图文、文件等格式。

## 功能特性

✅ 多平台支持：企业微信/钉钉/飞书/飞机  
✅ 多种消息类型：文本/卡片/文件/图文  
✅ 安全认证：支持签名校验  
✅ 媒体管理：文件上传API  
✅ 消息模板：预置常用消息模板  

## 安装

```sh
  go get -u github.com/Vectutil/sendx
```

## 企业微信示例

```go
    wc := we_com.NewWeComConfig("YOUR_API_KEY")
    
    // 发送文本消息
    err := wc.SendMsg(context.Background(),
    we_com.TextMessage("系统告警通知"),
    we_com.WithAtAll(),
    )
```

支持的消息类型：
- `TextNoticeCard` 通知卡片
- `MarkdownV2Message` Markdownv2格式
- `FileMessage` 文件消息
- `TemplateCardMessage` 模板卡片
- `...` 

## 钉钉示例

```go
// 初始化配置
dd := dingding.NewDingDingConfig(
    "YOUR_ACCESS_TOKEN",
    "YOUR_SECRET",
)

// 发送Markdown消息
err := dd.SendMsg(context.Background(),
    dingding.MarkdownMessage("服务器状态", "**CPU使用率**: 98%")
)
```

## 贡献指南

1. Fork仓库并创建分支
2. 提交测试用例
3. 发起Pull Request

## 许可证
[MIT License](/LICENSE)