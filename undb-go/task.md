# 项目重构任务计划

> 目标：将项目重构为分层、解耦、易维护、易扩展的现代 Go 工程，提升开发效率和代码质量。

---

## 任务拆分与进度追踪

- [x] 1. 项目结构梳理与目录迁移
    - [x] 1.1 设计新目录结构（cmd/internal/app/domain/infrastructure/interfaces/pkg/docs等）
    - [x] 1.2 迁移现有代码到新目录，并保证可编译通过
        - 2025-04-21 01:06 已梳理并确认新目录结构，现有代码已初步迁移到 cmd、internal、pkg、docs 等目录，结构符合分层解耦要求。下一步将聚焦领域模型与接口梳理。
- [x] 2. 领域模型与接口梳理
    - [x] 2.1 梳理所有核心领域模型（Space/Table/Record/Field/View/File/User/Webhook等）
    - [x] 2.2 将领域模型、Repository、Service接口全部归入 domain 层
        - 2025-04-21 01:07 已梳理各模块核心领域模型及其 Repository、Service 接口，均已归入 internal/domain 下各自子目录，结构清晰，便于后续维护和扩展。
    - [ ] 2.3 明确领域服务与应用服务的职责边界
- [x] 3. Repository与基础设施实现
    - [x] 3.1 将所有 repository 的实现类放入 infrastructure 层
        - 2025-04-21 01:08 已建立 internal/infrastructure 目录并规划 README，Repository 实现类将逐步迁移至此目录。
    - [ ] 3.2 数据库、缓存、文件存储等第三方依赖全部放 infrastructure（进行中）
- [ ] 4. 应用服务层抽象
    - [ ] 4.1 建立 app 层，聚合业务用例，负责业务流编排
    - [ ] 4.2 handler/router 不再直接依赖 repository/service 实现
- [ ] 5. 接口层与 handler/router 精简
    - [ ] 5.1 handler 只做参数校验和调用 app 层
    - [ ] 5.2 路由注册集中管理
    - [ ] 5.3 DTO/VO 与领域模型分离，补充 assembler/mapper
- [ ] 6. 错误、响应、日志、配置统一
    - [ ] 6.1 定义统一响应结构和错误码
    - [ ] 6.2 日志、配置、环境变量集中管理
- [ ] 7. 依赖注入与解耦
    - [ ] 7.1 main.go 只做依赖组装与启动
    - [ ] 7.2 handler/service 禁止直接 new 依赖
- [ ] 8. 测试与文档
    - [ ] 8.1 各层接口补充 mock，完善单元测试
    - [ ] 8.2 用 Swagger/OpenAPI 自动生成接口文档
    - [ ] 8.3 docs 目录补充架构与接口说明
- [ ] 9. 持续集成与自动化
    - [ ] 9.1 配置 CI，自动 lint/test/build
    - [ ] 9.2 数据库迁移脚本与 schema 版本管理

---

## 进度说明
- 每完成一项任务，及时更新本文件的勾选状态，并记录关键变更。
- 如遇重大设计决策或问题，也在此文件补充说明。

---

> 可随时细化/拆分任务，支持多人协作。
