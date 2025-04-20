# API 文档

## 认证

所有 API 请求都需要在 Header 中包含 `Authorization` 字段：

```
Authorization: Bearer <token>
```

## 空间管理

### 创建空间

```http
POST /api/v1/spaces
```

请求体：

```json
{
  "name": "我的空间",
  "description": "这是一个示例空间"
}
```

响应：

```json
{
  "id": "space_123",
  "name": "我的空间",
  "description": "这是一个示例空间",
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

### 获取空间列表

```http
GET /api/v1/spaces
```

响应：

```json
{
  "spaces": [
    {
      "id": "space_123",
      "name": "我的空间",
      "description": "这是一个示例空间",
      "created_at": "2024-04-20T08:00:00Z",
      "updated_at": "2024-04-20T08:00:00Z"
    }
  ]
}
```

## 表管理

### 创建表

```http
POST /api/v1/spaces/{space_id}/tables
```

请求体：

```json
{
  "name": "用户表",
  "description": "存储用户信息"
}
```

响应：

```json
{
  "id": "table_123",
  "name": "用户表",
  "description": "存储用户信息",
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

### 获取表列表

```http
GET /api/v1/spaces/{space_id}/tables
```

响应：

```json
{
  "tables": [
    {
      "id": "table_123",
      "name": "用户表",
      "description": "存储用户信息",
      "created_at": "2024-04-20T08:00:00Z",
      "updated_at": "2024-04-20T08:00:00Z"
    }
  ]
}
```

## 字段管理

### 创建字段

```http
POST /api/v1/tables/{table_id}/fields
```

请求体：

```json
{
  "name": "用户名",
  "type": "text",
  "description": "用户名称",
  "required": true,
  "unique": true
}
```

响应：

```json
{
  "id": "field_123",
  "name": "用户名",
  "type": "text",
  "description": "用户名称",
  "required": true,
  "unique": true,
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

### 获取字段列表

```http
GET /api/v1/tables/{table_id}/fields
```

响应：

```json
{
  "fields": [
    {
      "id": "field_123",
      "name": "用户名",
      "type": "text",
      "description": "用户名称",
      "required": true,
      "unique": true,
      "created_at": "2024-04-20T08:00:00Z",
      "updated_at": "2024-04-20T08:00:00Z"
    }
  ]
}
```

## 记录管理

### 创建记录

```http
POST /api/v1/tables/{table_id}/records
```

请求体：

```json
{
  "fields": {
    "username": "john_doe",
    "email": "john@example.com",
    "age": 30
  }
}
```

响应：

```json
{
  "id": "record_123",
  "fields": {
    "username": "john_doe",
    "email": "john@example.com",
    "age": 30
  },
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

### 获取记录列表

```http
GET /api/v1/tables/{table_id}/records
```

响应：

```json
{
  "records": [
    {
      "id": "record_123",
      "fields": {
        "username": "john_doe",
        "email": "john@example.com",
        "age": 30
      },
      "created_at": "2024-04-20T08:00:00Z",
      "updated_at": "2024-04-20T08:00:00Z"
    }
  ]
}
```

## 错误响应

所有错误响应都遵循以下格式：

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "错误描述",
    "details": {
      "field": "具体错误信息"
    }
  }
}
```

常见错误码：

- `INVALID_REQUEST`: 请求参数错误
- `UNAUTHORIZED`: 未授权
- `FORBIDDEN`: 禁止访问
- `NOT_FOUND`: 资源不存在
- `INTERNAL_ERROR`: 服务器内部错误 