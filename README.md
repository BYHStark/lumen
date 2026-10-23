# Lumen

一个用来把 Go 后端练到能独立交付的学徒项目：多租户 API 计量、账本，以及 Webhook 可靠投递。

业务代码还没开始。先看 [docs/curriculum.md](docs/curriculum.md)，当前进度在 [docs/progress.md](docs/progress.md)。

## 在 Codex 里继续

用 Codex 打开这个文件夹，不要在别的空目录里接着做。仓库技能在 `.agents/skills/lumen-mentor/`。新对话可以说：

```text
Use $lumen-mentor to recover Lumen progress and assign the next learning task.
```

进度写在仓库里，不依赖某一次聊天。

## 以后上传 GitHub

远程建好之后：

```bash
git remote add origin <你的仓库地址>
git push -u origin main
```

模块路径现在是 `lumen`。如果要把这个模块公开给别的项目引用，再把 `go.mod` 第一行改成 GitHub 路径。只作为自己的服务运行时，可以先不改。

## 本机要求

系统默认 `go` 保持 Homebrew 的 1.24，避免影响其他项目。本仓库在 `go.mod` 里指定 `toolchain go1.27.2`，只有进入这个目录后 `go` 才会切到 1.27。不要把 `/usr/local/go/bin` 加进全局 PATH。
