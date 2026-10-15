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

当前仓库按已安装的 Go 1.24 初始化。课程目标是 Go 1.27；升级是阶段 0 的第一步，升级前不要使用 1.27 才有的语言和标准库功能。
