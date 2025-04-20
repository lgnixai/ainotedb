# 测试数据

## 空间数据

```json
{
  "id": "space_1",
  "name": "测试空间",
  "description": "用于测试的空间",
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

## 表数据

### 用户表

```json
{
  "id": "table_1",
  "name": "用户表",
  "description": "存储用户信息",
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

### 订单表

```json
{
  "id": "table_2",
  "name": "订单表",
  "description": "存储订单信息",
  "created_at": "2024-04-20T08:00:00Z",
  "updated_at": "2024-04-20T08:00:00Z"
}
```

## 字段数据

### 用户表字段

```json
[
  {
    "id": "field_1",
    "name": "用户名",
    "type": "text",
    "description": "用户登录名",
    "required": true,
    "unique": true,
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  },
  {
    "id": "field_2",
    "name": "邮箱",
    "type": "text",
    "description": "用户邮箱",
    "required": true,
    "unique": true,
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  },
  {
    "id": "field_3",
    "name": "年龄",
    "type": "number",
    "description": "用户年龄",
    "required": false,
    "unique": false,
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  }
]
```

### 订单表字段

```json
[
  {
    "id": "field_4",
    "name": "订单号",
    "type": "text",
    "description": "订单编号",
    "required": true,
    "unique": true,
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  },
  {
    "id": "field_5",
    "name": "金额",
    "type": "number",
    "description": "订单金额",
    "required": true,
    "unique": false,
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  },
  {
    "id": "field_6",
    "name": "状态",
    "type": "select",
    "description": "订单状态",
    "required": true,
    "unique": false,
    "options": "待支付,已支付,已发货,已完成,已取消",
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  }
]
```

## 记录数据

### 用户记录

```json
[
  {
    "id": "record_1",
    "fields": {
      "username": "john_doe",
      "email": "john@example.com",
      "age": 30
    },
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  },
  {
    "id": "record_2",
    "fields": {
      "username": "jane_smith",
      "email": "jane@example.com",
      "age": 25
    },
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  }
]
```

### 订单记录

```json
[
  {
    "id": "record_3",
    "fields": {
      "order_number": "ORD-2024-001",
      "amount": 99.99,
      "status": "待支付"
    },
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  },
  {
    "id": "record_4",
    "fields": {
      "order_number": "ORD-2024-002",
      "amount": 199.99,
      "status": "已完成"
    },
    "created_at": "2024-04-20T08:00:00Z",
    "updated_at": "2024-04-20T08:00:00Z"
  }
]
```

## 测试用例

### 创建空间

```bash
curl -X POST http://localhost:8080/api/v1/spaces \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "name": "测试空间",
    "description": "用于测试的空间"
  }'
```

### 创建表

```bash
curl -X POST http://localhost:8080/api/v1/spaces/space_1/tables \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "name": "用户表",
    "description": "存储用户信息"
  }'
```

### 创建字段

```bash
curl -X POST http://localhost:8080/api/v1/tables/table_1/fields \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "name": "用户名",
    "type": "text",
    "description": "用户登录名",
    "required": true,
    "unique": true
  }'
```

### 创建记录

```bash
curl -X POST http://localhost:8080/api/v1/tables/table_1/records \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "fields": {
      "username": "john_doe",
      "email": "john@example.com",
      "age": 30
    }
  }'
``` 