# infrastructure 说明

此目录用于存放所有 Repository 的具体实现、数据库、缓存、文件存储等第三方依赖的实现代码。

- db/: GORM 实现、数据库连接、迁移等
- cache/: Redis 等缓存实现
- file/: 文件存储相关实现

所有实现应依赖 domain 层定义的接口，禁止依赖 handler/service。
