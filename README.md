# 电商独立站后端

Go 语言实现的跨境电商独立站后端。通用电商模型，MVP 目标 = 最小闭环 + 运营后台；兼作求职作品集与副业尝试。

## 技术栈

| 层 | 选型 |
| --- | --- |
| Web 框架 | Gin |
| 数据访问 | sqlx（手写 SQL） |
| 存储 | SQLite |
| 定时任务 | robfig/cron |
| 鉴权 | golang-jwt + bcrypt |

## 项目结构

```
ecommerce-backend/
├── cmd/server/main.go      # 入口：配置、连库、启 cron、启 HTTP
├── internal/
│   ├── config/             # 配置加载
│   ├── model/              # 结构体
│   ├── repository/         # 数据访问（sqlx）
│   ├── service/            # 业务逻辑
│   ├── handler/            # HTTP 层
│   ├── middleware/         # JWT、日志、错误恢复、CORS
│   ├── router/             # 路由
│   └── cron/               # 定时任务
├── migrations/             # 建表 SQL
└── go.mod
```

分层原则：handler 只做 HTTP，service 管业务，repository 管数据。

## 开发路线

- **Phase 0 · 骨架**：项目结构 + 配置 + 日志 + 统一响应/错误 + `/healthz` —— 学 Gin 路由、分层、错误处理
- **Phase 1 · 用户域**：注册/登录、JWT 签发、鉴权中间件 —— 学 bcrypt、JWT、middleware、sqlx 起步
- **Phase 2 · 商品域**：分类/商品，公开列表+详情 + 后台 CRUD —— 学建表、手写 SQL、分页
- **Phase 3 · 购物车**：加购/改数量/删/列表（需登录）—— 学登录态、关联查询
- **Phase 4 · 订单 + mock 支付 + cron**：下单事务、模拟支付回调、cron 超时取消 —— 学事务、状态机、cron
- **Phase 5 · 运营后台**：admin 角色、商品/订单管理 —— 学简化 RBAC
- **Phase 6 · 加分项（可选）**：优惠券、多币种、单元测试、docker、架构图

## 核心数据模型

- `users`（含 role）
- `addresses`
- `categories` / `products`
- `cart_items`
- `orders` / `order_items`
- `payments`（mock）

## 版本控制

每完成一个 Phase（或可 demo 的一步）就 `git commit` 并 `git push`。
