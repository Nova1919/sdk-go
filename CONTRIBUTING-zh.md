# Scrapeless SDK Go 贡献与开发指南

感谢您关注并参与 Scrapeless Actor SDK Go 的开源贡献！本指南同时涵盖贡献流程和开发规范。

## 如何参与贡献

### 1. 提交 Issue

- 在新建 issue 前，请先搜索 [已有 issue](https://github.com/scrapeless-ai/sdk-go/issues)。
- 请提供清晰的标题和详细描述。
- 如涉及 bug，请附上复现步骤、期望行为和相关环境信息。

### 2. 提交 Pull Request (PR)

- Fork 本仓库，并从 `main` 分支创建新分支。
- 遵循 [Conventional Commits](https://www.conventionalcommits.org/) 提交规范。
- 确保代码通过测试。
- 如有需要，请补充或更新文档和测试。
- 提交 PR 时请详细描述变更内容，并关联相关 issue。
- 积极响应评审意见并及时更新 PR。

### 3. 代码规范

- 使用 Go 语言，遵循现有代码风格，可使用 `go fmt` 格式化代码。
- 注释和文档应清晰、易懂。
- 新功能或修复请添加相应测试。

### 4. 分支管理

- 建议使用功能分支（如 `feature/xxx` 或 `fix/xxx`）。
- 保持 PR 聚焦单一主题，避免混合无关更改。

## 本地开发流程

### 1. 克隆与安装依赖

```bash
# 克隆仓库
git clone https://github.com/scrapeless-ai/sdk-go.git
cd scrapeless-actor-sdk-go

# 安装依赖
go mod tidy
```
### 2. 常用开发命令
```bash
# 运行测试
go run ./...

# 代码格式化
go fmt ./...

# 构建项目
go build
```

### 3. 项目结构

```text
env/
├── config.go
├── env.go
└── env_test.go
example/
├── actor/
│   └── actor.go
├── actor_service/
│   └── actor_service.go
├── browser/
│   └── browser.go
├── captcha/
│   └── captcha.go
├── deepserp/
│   └── deepserp.go
├── httpserver/
│   └── httpserver.go
├── proxy/
│   └── proxy.go
├── router/
│   └── router.go
├── scraping/
│   └── scraping.go
├── storage_dataset/
│   └── storage_dataset.go
├── storage_kv/
│   └── storage_kv.go
├── storage_object/
│   └── storage_object.go
├── storage_queue/
│   └── storage_queue.go
└── universal/
    └── universal.go
internal/
├── code/
│   └── code.go
├── helper/
│   ├── context_util.go
│   ├── env.go
│   ├── grpc.go
│   ├── gzip.go
│   ├── redis_extends.go
│   └── utils.go
└── remote/
    ├── actor/
    ├── browser/
    ├── captcha/
    ├── deepserp/
    ├── proxy/
    ├── request/
    ├── router/
    ├── scraping/
    ├── storage/
    └── universal/
scrapeless/
├── actor/
│   ├── actor.go
│   └── actor_test.go
├── client.go
├── log/
│   ├── api.go
│   ├── api_test.go
│   └── log.go
└── services/
    ├── actor/
    ├── actor_test.go
    ├── browser/
    ├── captcha/
    ├── deepserp/
    ├── httpserver/
    ├── proxies/
    ├── router/
    ├── scraping/
    ├── storage/
    └── universal/
```

## 代码质量工具

### Go fmt

- 代码格式化：`go fmt ./...`

## 最佳实践

- API 密钥和敏感信息请使用环境变量配置。
- 所有 API 调用建议处理错误。
- 用完资源（如网络连接等）请及时关闭和清理。
- 注意 API 速率限制，合理设置超时时间。
- 保持提交原子性，信息清晰。
- PR 前认真自查代码质量。

### 提交信息规范

本项目使用 Conventional Commits 规范，支持以下类型：

- `feat`: ✨ 新功能
- `fix`: 🐛 错误修复
- `docs`: 📚 文档更新
- `style`: 💎 代码格式（不影响功能的更改）
- `refactor`: 📦 代码重构
- `perf`: 🚀 性能优化
- `test`: 🚨 测试相关
- `build`: 🛠 构建系统或外部依赖
- `ci`: ⚙️ CI 配置文件和脚本
- `chore`: ♻️ 其他更改（不会出现在 CHANGELOG 中）
- `revert`: 🗑 回滚提交

#### 提交信息格式

```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### 示例

```bash
feat: 添加新的 API 端点
fix(auth): 修复登录验证问题
docs: 更新 API 文档
```

### 版本号规则

项目遵循 [Semantic Versioning](https://semver.org/) 规范：

- **MAJOR**: 不兼容的 API 更改
- **MINOR**: 向后兼容的功能添加
- **PATCH**: 向后兼容的错误修复

### CHANGELOG

所有版本更改都会自动记录在 [CHANGELOG.md](./CHANGELOG.md) 文件中，包括：

- 新功能
- 错误修复
- 重大更改
- 性能改进
- 其他重要更新

### 注意事项

1. 确保所有提交都遵循 Conventional Commits 规范
2. 重大更改需要在提交信息中添加 `BREAKING CHANGE:` 标记
3. 手动发版后需要推送 tags：`git push --follow-tags origin main`

## 行为准则

请保持尊重与包容。我们遵循 [Contributor Covenant](https://www.contributor-covenant.org/) 行为准则。

## 联系方式

如有疑问或需支持，请提交 issue 或发送邮件至 [support@scrapeless.com](mailto:support@scrapeless.com)。

---

感谢您的贡献，让 Scrapeless Actor SDK Go 更加完善！ 

        