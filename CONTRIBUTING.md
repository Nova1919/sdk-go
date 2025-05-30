# Scrapeless Actor SDK Go Contribution & Development Guide

Thank you for your interest in contributing to the Scrapeless Actor SDK Go project! This guide outlines both the contribution process and development conventions.

## How to Contribute

### 1. Submit an Issue

* Before opening a new issue, please search through the [existing issues](https://github.com/scrapeless-ai/sdk-go/issues).
* Provide a clear title and detailed description.
* For bug reports, include steps to reproduce, expected behavior, and environment information.

### 2. Submit a Pull Request (PR)

* Fork the repository and create a new branch from the `main` branch.
* Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages.
* Ensure your code passes tests.
* Update or add documentation and tests as needed.
* Describe your changes clearly in the PR and reference any related issues.
* Be responsive to review feedback and update your PR accordingly.

### 3. Code Style

* Use the Go programming language and follow the existing code style.
* Format code using `go fmt`.
* Write clear and understandable comments and documentation.
* Add appropriate tests for new features or fixes.

### 4. Branch Management

* Use feature branches such as `feature/xxx` or `fix/xxx`.
* Keep PRs focused and avoid mixing unrelated changes.

## Local Development Workflow

### 1. Clone & Install Dependencies

```bash
# Clone the repository
git clone https://github.com/scrapeless-ai/sdk-go.git
cd scrapeless-actor-sdk-go

# Install dependencies
go mod tidy
```

### 2. Common Development Commands

```bash
# Run tests
go run ./...

# Format code
go fmt ./...

# Build the project
go build
```

### 3. Project Structure

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

## Code Quality Tools

### Go fmt

* Format code using: `go fmt ./...`

## Best Practices

* Use environment variables for API keys and sensitive information.
* Always handle errors in API calls.
* Close and clean up resources such as network connections.
* Respect API rate limits and set appropriate timeouts.
* Keep commits atomic and messages clear.
* Review code quality before submitting PRs.

### Commit Message Convention

This project follows the Conventional Commits specification. Supported types:

* `feat`: ✨ New feature
* `fix`: 🐛 Bug fix
* `docs`: 📚 Documentation update
* `style`: 💎 Code formatting (non-functional changes)
* `refactor`: 📦 Code refactoring
* `perf`: 🚀 Performance optimization
* `test`: 🚨 Test-related changes
* `build`: 🛠 Build system or external dependencies
* `ci`: ⚙️ CI configuration files and scripts
* `chore`: ♻️ Miscellaneous changes (not in CHANGELOG)
* `revert`: 🗑 Revert a commit

#### Commit Message Format

```text
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### Examples

```bash
feat: add new API endpoint
fix(auth): fix login validation issue
docs: update API documentation
```

### Versioning

This project uses [Semantic Versioning](https://semver.org/):

* **MAJOR**: incompatible API changes
* **MINOR**: backward-compatible feature additions
* **PATCH**: backward-compatible bug fixes

### CHANGELOG

All version changes are automatically recorded in [CHANGELOG.md](./CHANGELOG.md), including:

* New features
* Bug fixes
* Breaking changes
* Performance improvements
* Other significant updates

### Notes

1. Ensure all commits follow the Conventional Commits specification
2. Add `BREAKING CHANGE:` tag for major changes
3. After manual release, push tags with: `git push --follow-tags origin main`

## Code of Conduct

Please maintain respect and inclusiveness. We follow the [Contributor Covenant](https://www.contributor-covenant.org/) Code of Conduct.

## Contact

If you have questions or need support, please open an issue or email us at [support@scrapeless.com](mailto:support@scrapeless.com).

---

Thank you for your contribution to making Scrapeless Actor SDK Go better!
