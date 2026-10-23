# 进度

- 更新：2026-10-20
- 阶段：0
- 状态：等待用户
- 本机 Go：目录外 go1.24.13 darwin/arm64；本仓库 toolchain go1.27.2

## 已完成

- 仓库和导师技能已建立。证据：初始提交 `f3eea96`。
- 官方 Go 1.27.2 已安装在 `/usr/local/go/bin/go`。系统默认仍是 Homebrew 1.24.13。证据：仓库内 `go version go1.27.2 darwin/arm64`。
- 决定不替换系统默认 Go。证据：`docs/adr/0002-keep-system-go-1-24.md`。
- 本地 PostgreSQL 和 Redis 已启动。证据：`deploy/compose.yaml`。
- 用户能区分两类数据：账单进 PostgreSQL，重启后还在；限流计数进 Redis，按当前 Compose 重启后丢失。证据：2026-10-16 的口头回答。
- Makefile 的 `up` 和 `down` 可用，用户理解它是给命令起短名。证据：`make -n up` 与 `make -n down` 的输出；2026-10-16 的口头回答。
- 缺少环境变量的测试已先失败。证据：`TestLoadMissingEnv` 输出 `config_test.go:11: 缺少环境变量时必须返回错误`。当前 `Load` 仍是 `return nil`。

## 现在

用户要改 `Load`，让这个失败的测试通过。

## 下一步

用 `os.Getenv` 读取两个变量，任何一个为空就 `return errors.New(...)`，再运行同一个测试。

## 未决

- 无。

## 跳过的验收

- 无。Compose 这一片是用户要求代写的。
