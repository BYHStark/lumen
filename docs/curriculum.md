# Lumen 课程

兼职每周约 12 小时，大约 5 个月。全职大约 10 到 12 周。

做完不能自动变成资深。毕业标准是：不看实现，也能独立处理超时、双重关账、租户越权和一次慢查询，并讲清投递语义。

每一阶段结束要有测试、一篇 ADR，以及一个能演示的失败场景。只有成功路径不算完成。

## 不变的保证

- 同一租户、同一幂等键，用量只记一次。
- 投递是 at-least-once。客户按事件 ID 去重。
- 两个关账进程同时跑，只生成一张发票。
- 租户 A 不能读、改或投递租户 B 的数据。
- 已关账金额只追加反向分录，不原地修改。

## 阶段 0：工程底座

安装 Go 1.27 并把 `go.mod` 改到该版本。然后做模块、Compose 里的 PostgreSQL 和 Redis、配置、`log/slog`、lint、`go test -race`。

验收：`make up && make test && make lint` 能重复跑。

在升级完成前，不使用泛型方法、`encoding/json/v2`、`testing/synctest` 或 `goroutineleak`。前两个是 Go 1.27 的内容，后两个从 1.25 起可用；这个仓库现在按 1.24 初始化。

## 阶段 1：身份与多租户

租户、用户、成员角色、API Key 的生成和撤销。密码用 argon2id 或 bcrypt。Key 只存哈希和可展示前缀，比较用常量时间。会话先用 opaque token，不上 JWT。

验收：租户隔离的 repository 测试。错误密码和错误 Key 不泄露多余的存在性信息。

## 阶段 2：控制面 HTTP

chi、中间件、统一错误、分页、写接口的 `Idempotency-Key`。先手写 handler，后补 OpenAPI。`recover` 只放在请求边界。

验收：未登录、过期、越权、重复提交都有测试。

## 阶段 3：计量网关

另一组路由表示客户在调用出售的 API。校验 Key。自己实现 Redis 令牌桶，不要只调一个限流库。先同步写用量，再把用量和出箱放进同一事务。

验收：相同幂等键打两次，只记一次。Redis 不可用时的行为写进 ADR。

## 阶段 4：可靠投递

这是核心阶段。

Worker 在事务里用 `FOR UPDATE SKIP LOCKED` 领取，设置 `locked_until` 租约，提交后再发 HTTP。成功则完成；失败则指数退避加抖动，超过次数进死信。进程死在发送和落库之间时，租约到期后可以被别人接管。

出站请求必须有超时。签名用 HMAC。Webhook URL 拒绝私网、链路本地和云元数据地址。慢客户端不能把 worker 打满。优雅退出停止领取，并处理正在飞的请求。

Go 1.25 及以上用 `testing/synctest` 测退避，测试里不要真睡。

验收：

- 客户前两次返回 500，第三次成功；最终一次成功，用量不重复。
- 杀掉 worker 后重启，未完成任务会被再投递。
- 客户一直超时，worker 和 goroutine 有上限。
- 关闭进程时不再领取新任务。

## 阶段 5：账本与账单

按小时汇总。关账用 PostgreSQL advisory lock。发票在一个事务里生成。时区写进套餐，不靠服务器本地时区。

验收：两个 scheduler 同时关账，只有一张发票。少记用反向分录修正。

## 阶段 6：调度器、实时、CLI

scheduler 单独进程，用 advisory lock 做单主。WebSocket 推送投递尝试。慢消费者策略必须写下来：丢弃、断开，或有界缓冲。CLI 先用标准库 `flag`。

验收：两个 scheduler 只有一个在关账。慢 WebSocket 客户端不会拖死 hub。

## 阶段 7：可观测与性能

日志带 request id 和 trace id。指标至少有投递延迟、队列深度、失败率和限流次数。trace 穿过 API、数据库、worker 和出站 HTTP。pprof 不能裸露在公网。Go 1.27 再启用 `goroutineleak`。

故意留一个性能问题，用 benchmark、`EXPLAIN` 和 pprof 找到它。没有测量，不加 `sync.Pool`。

验收：一页纸性能笔记，有改动前后的数字。

## 阶段 8：故障与安全

演练 worker 中途被杀、Postgres 重启、Redis 清空、客户变慢。测试参数化 SQL、租户越权、Key 不进日志、Webhook 地址拒绝私网。给解析或签名规范化做一次 fuzz。

验收：一份事故复盘，包含时间线、影响、根因、哪条测试没挡住、接下来改什么。

## 阶段 9：加宽

选做，不要全做，也不要推倒重来：一个小 gRPC、一个查询包迁到 sqlc、OpenAPI 生成、Redis Stream 或 NATS、一页 `html/template` 管理页、一个泛型方法助手。

`unsafe`、cgo 和汇编只写 ADR 说明为什么产品不用。

## 阶段 10：交付

陌生人 30 分钟内能跑起来的 README、至少 5 篇 ADR、三份 runbook：投递卡住、账单不平、Key 泄露。

## 毕业

不看实现，独立完成：

1. 客户 60% 超时，系统不失控，恢复后用量不重复。
2. 两个关账进程只有一张发票。
3. 自己写的租户越权测试能拦住。
4. 一个慢查询能指出原因。
5. 讲清 at-least-once、at-most-once、exactly-once，以及 Lumen 保证哪一种。
