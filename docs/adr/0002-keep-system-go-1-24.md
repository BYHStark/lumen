# ADR 0002：系统默认 Go 保持 1.24，只在本仓库使用 1.27

- 状态：已接受
- 日期：2026-10-15

## 背景

官方安装包已经把 Go 1.27.2 装到 `/usr/local/go/bin/go`。shell 的 `PATH` 里 Homebrew 更靠前，`/opt/homebrew/bin/go` 仍指向 `go@1.24`，所以终端里的 `go version` 继续显示 1.24.13。

用户还有公司项目。把系统默认 `go` 换成 1.27 后，那些项目里的 `go test`、`gofmt` 和 `go mod tidy` 也会跟着变。Go 1 兼容性承诺通常不会让旧代码编译失败，但 1.27 的格式化空白、新增 vet，以及 `encoding/json` 的错误文本都可能和公司 CI 上的 1.24 不一致。

## 决定

不修改 `/opt/homebrew/bin/go`，也不把 `/usr/local/go/bin` 加进全局 `PATH`。

本仓库的 `go.mod` 写 `go 1.27.2` 和 `toolchain go1.27.2`。在这个模块目录里运行 `go` 时，工具链切到 1.27.2。离开这个目录后，仍使用 Homebrew 的 1.24.13。

## 后果

Lumen 可以使用泛型方法和 `encoding/json/v2`。公司项目的默认工具链不变。不要为了本仓库执行 `brew unlink go@1.24` 或改 shell 配置。
