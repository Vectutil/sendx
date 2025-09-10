import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Vectutil/sendx',
  description: '一个用于发送消息到各种平台的工具库',
  
  head: [
    ['link', { rel: 'icon', href: 'https://avatars.githubusercontent.com/u/224289848?s=48&v=4' }]
  ],
  
  // 配置中文语言支持
  locales: {
    root: {
      label: '简体中文',
      lang: 'zh-CN',
      themeConfig: {
        nav: [
          { text: '首页', link: '/' },
          { text: '指南', link: '/guide/' },
          { text: 'API 参考', link: '/api/' },
          { text: 'GitHub', items: [
            { text: '组织', link: 'https://github.com/Vectutil' },
            { text: '仓库', link: 'https://github.com/Vectutil/sendx' }
          ] }
        ],
        
        sidebar: [
          {
            text: '指南',
            items: [
              { text: '介绍', link: '/guide/' },
              { text: '快速开始', link: '/guide/getting-started' }
            ]
          },
          {
            text: 'API 参考',
            items: [
              { text: '概览', link: '/api/' },
              { text: '飞书', link: '/api/lark' },
              { text: '钉钉', link: '/api/dingding' },
              { text: '企业微信', link: '/api/wechat-work' },
              { text: 'Telegram', link: '/api/telegram' }
            ]
          }
        ],
        
        footer: {
          message: 'Released under the MIT License.',
          copyright: 'Copyright © 2025 Vectutil'
        },
        
        docFooter: {
          prev: '上一页',
          next: '下一页'
        },
        
        outline: {
          label: '页面导航'
        },
        
        lastUpdated: {
          text: '最后更新于',
          formatOptions: {
            dateStyle: 'short',
            timeStyle: 'medium'
          }
        },
        
        langMenuLabel: '多语言',
        returnToTopLabel: '回到顶部',
        sidebarMenuLabel: '菜单',
        darkModeSwitchLabel: '主题',
        lightModeSwitchTitle: '切换到浅色模式',
        darkModeSwitchTitle: '切换到深色模式'
      }
    }
  }
})