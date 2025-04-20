# API 文档

本文件整理了主要 RESTful API，包括接口说明、请求参数、示例请求、响应 JSON 结构和测试数据。

## 认证
所有 API 请求需在 Header 中包含：
```
Authorization: Bearer <token>
```

---

## 空间管理

### 创建空间
- **URL**: `POST /api/v1/spaces`
- **请求参数**：
```json
{
  "name": "团队空间名称",
  "owner_id": "用户ID",
  "visibility":"public"
}
```
- **返回示例**：
```json
{
  "id": "spc_1234567890",
  "name": "团队空间名称",
  "owner_id": "用户ID",
  "created_at": "2025-04-21T00:00:00Z"
}
```

### 获取空间列表
- **URL**: `GET /api/v1/spaces`
- **返回示例**：
```json
[
  {
    "id": "spc_1234567890",
    "name": "团队空间名称",
    "owner_id": "用户ID",
    "created_at": "2025-04-21T00:00:00Z"
  }
]
```

---

## 用户管理

### 用户注册
- **URL**: `POST /api/v1/users/register`
- **请求参数**：
```json
{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com",
  "name": "测试用户"
}
```
- **返回示例**：
```json
{
  "id": "usr_001",
  "username": "testuser"
}
```

### 用户登录
- **URL**: `POST /api/v1/users/login`
- **请求参数**：
```json
{
  "email": "test@example.com",
  "password": "123456"
}
```
- **返回示例**：
```json
{
  "id": "usr_001",
  "email": "test@example.com",
  "name": "测试用户"
}
```
- **失败示例**：
```json
{
  "error": "invalid email or password"
}
```

---

## 表管理

### 创建表
- **URL**: `POST /api/v1/tables`
- **请求参数**：
```json
{
  "name": "客户表",
  "space_id": "spc_1234567890"
}
```
- **返回示例**：
```json
{
  "id": "tbl_001",
  "name": "客户表",
  "space_id": "spc_1234567890",
  "created_at": "2025-04-21T00:00:00Z"
}
```

### 获取表信息
- **URL**: `GET /api/v1/tables/:id`
- **返回示例**：
```json
{
  "id": "tbl_001",
  "name": "客户表",
  "space_id": "spc_1234567890",
  "created_at": "2025-04-21T00:00:00Z"
}
```

---

## 字段管理

### 创建字段
- **URL**: `POST /api/v1/fields`
- **请求参数**：
```json
{
  "name": "客户名称",
  "table_id": "tbl_001",
  "type": "string"
}
```
- **返回示例**：
```json
{
  "id": "fld_001",
  "name": "客户名称",
  "table_id": "tbl_001",
  "type": "string"
}
```

---

## 记录管理

### 新增记录
- **URL**: `POST /api/v1/records`
- **请求参数**：
```json
{
  "table_id": "tbl_001",
  "data": {
    "客户名称": "张三"
  }
}
```
- **返回示例**：
```json
{
  "id": "rec_001",
  "table_id": "tbl_001",
  "data": {
    "客户名称": "张三"
  },
  "created_at": "2025-04-21T00:00:00Z"
}
```

### 获取记录
- **URL**: `GET /api/v1/records/:id`
- **返回示例**：
```json
{
  "id": "rec_001",
  "table_id": "tbl_001",
  "data": {
    "客户名称": "张三"
  },
  "created_at": "2025-04-21T00:00:00Z"
}
```

---

## 附件上传

### 上传文件
- **URL**: `POST /api/v1/files`
- **请求参数**（multipart/form-data）：
  - file: 文件
  - record_id: 记录ID
  - table_id: 表ID
- **返回示例**：
```json
{
  "id": "file_001",
  "name": "test.png",
  "size": 12345,
  "mime_type": "image/png",
  "url": "/api/files/file_001"
}
```

---

（更多接口可根据 handler/router/service 代码自动补全，支持增删改查、批量、配置等）

> 若需补充具体接口或返回结构，请告知接口名。
