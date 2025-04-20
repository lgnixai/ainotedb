# API 文档（自动分析生成）

## 用户相关

### 注册
- URL: `POST /api/users/register`
- 输入：
  - `email` (string, required)
  - `password` (string, required)
  - `name` (string, required)
- 输出：
  - 用户对象

### 登录
- URL: `POST /api/users/login`
- 输入：
  - `email` (string, required)
  - `password` (string, required)
- 输出：
  - `token` (string)
  - `user` (object)

### 获取当前用户
- URL: `GET /api/users/:id`
- 需登录，token 认证
- 输出：
  - 用户对象

### 更新用户
- URL: `PUT /api/users/:id`
- 需登录，token 认证
- 输入：用户对象（JSON）
- 输出：用户对象

### 删除用户
- URL: `DELETE /api/users/:id`
- 需登录，token 认证


## 空间（Space）相关

### 创建空间
- URL: `POST /api/spaces`
- 需登录，token 认证
- 输入：
  - `name` (string, required)
  - `description` (string, 可选)
  - `visibility` (string, 可选，默认 public)
- 输出：
  - 空间对象

### 获取空间列表
- URL: `GET /api/spaces`
- 需登录，token 认证
- 输出：
  - 空间数组

### 获取空间详情
- URL: `GET /api/spaces/:id`
- 需登录，token 认证
- 输出：
  - 空间对象

### 更新空间
- URL: `PUT /api/spaces/:id`
- 需登录，token 认证
- 输入：空间对象
- 输出：空间对象

### 删除空间
- URL: `DELETE /api/spaces/:id`
- 需登录，token 认证

### 空间成员管理
- 添加成员：`POST /api/spaces/:id/members`  输入：{ user_id, role }
- 移除成员：`DELETE /api/spaces/:id/members/:user_id`
- 更新成员角色：`PUT /api/spaces/:id/members/:user_id/role`  输入：{ role }
- 获取成员列表：`GET /api/spaces/:id/members`


## 表（Table）相关

### 创建表
- URL: `POST /api/tables`
- 需登录，token 认证
- 输入：表对象
- 输出：表对象

### 获取表详情
- URL: `GET /api/tables/:id`
- 需登录，token 认证
- 输出：表对象

### 获取空间下所有表
- URL: `GET /api/tables/space/:space_id`
- 需登录，token 认证
- 输出：表数组

### 更新表
- URL: `PUT /api/tables/:id`
- 需登录，token 认证
- 输入：表对象
- 输出：表对象

### 删除表
- URL: `DELETE /api/tables/:id`
- 需登录，token 认证


## 视图（View）相关

### 创建视图
- URL: `POST /api/views`
- 需登录，token 认证
- 输入：视图对象
- 输出：视图对象

### 获取视图详情
- URL: `GET /api/views/:id`
- 需登录，token 认证
- 输出：视图对象

### 获取表下所有视图
- URL: `GET /api/views/table/:tableId`
- 需登录，token 认证
- 输出：视图数组

### 更新视图
- URL: `PUT /api/views/:id`
- 需登录，token 认证
- 输入：视图对象
- 输出：视图对象

### 删除视图
- URL: `DELETE /api/views/:id`
- 需登录，token 认证

### 更新视图配置
- URL: `PUT /api/views/:id/config`
- 需登录，token 认证
- 输入：配置对象
- 输出：视图对象


## 记录（Record）相关

### 创建记录
- URL: `POST /api/records`
- 需登录，token 认证
- 输入：记录对象
- 输出：记录对象

### 获取记录详情
- URL: `GET /api/records/:id`
- 需登录，token 认证
- 输出：记录对象

### 获取表下所有记录
- URL: `GET /api/records/:tableId`
- 需登录，token 认证
- 输出：记录数组

### 更新记录
- URL: `PUT /api/records/:id`
- 需登录，token 认证
- 输入：记录对象
- 输出：记录对象

### 删除记录
- URL: `DELETE /api/records/:id`
- 需登录，token 认证

### 批量创建记录
- URL: `POST /api/records/batch-create`
- 需登录，token 认证
- 输入：{ table_id, records: [...] }
- 输出：创建结果

### 批量更新记录
- URL: `POST /api/records/batch-update`
- 需登录，token 认证
- 输入：{ records: [...] }
- 输出：更新结果

### 批量删除记录
- URL: `POST /api/records/batch-delete`
- 需登录，token 认证
- 输入：{ ids: [...] }
- 输出：删除结果

### 聚合
- URL: `POST /api/records/aggregate`
- 需登录，token 认证
- 输入：聚合请求参数
- 输出：聚合结果

### 透视
- URL: `POST /api/records/pivot`
- 需登录，token 认证
- 输入：透视请求参数
- 输出：透视结果


## 字段（Field）相关

### 创建字段
- URL: `POST /api/fields`
- 需登录，token 认证
- 输入：字段对象
- 输出：字段对象

### 获取字段详情
- URL: `GET /api/fields/:id`
- 需登录，token 认证
- 输出：字段对象

### 获取表下所有字段
- URL: `GET /api/fields/table/:table_id`
- 需登录，token 认证
- 输出：字段数组

### 更新字段
- URL: `PUT /api/fields/:id`
- 需登录，token 认证
- 输入：字段对象
- 输出：字段对象

### 删除字段
- URL: `DELETE /api/fields/:id`
- 需登录，token 认证


## 实时（WebSocket）

### 连接 WebSocket
- URL: `GET /ws`
- 需登录，token 认证（通常通过 Cookie 或 token）
- 用于实时协作/推送

---

**说明：**
- 所有 /api/xxx 路由均需带上 `Authorization: Bearer <token>` 头（除注册、登录外）。
- 输入字段、输出字段如需详细结构，请参考具体 handler/model 定义。
- 如需详细字段说明，请指定具体接口。
