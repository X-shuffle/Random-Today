# Random Today API 文档

## 概述

Random Today API 提供个人偏好管理功能，支持添加、查询、删除和随机选择偏好。

## 基础信息

- **基础URL**: `http://localhost:8080/api`
- **内容类型**: `application/json`
- **字符编码**: `UTF-8`

## 偏好类型

系统支持以下偏好类型：

- `main_food` - 主食
- `snack` - 零食
- `scenery` - 风景
- `milk_tea` - 奶茶

## API 端点

### 1. 创建偏好

**POST** `/preferences`

创建新的偏好记录。

#### 请求体

```json
{
  "name": "宫保鸡丁",
  "type": "main_food",
  "description": "经典川菜，口感麻辣鲜香"
}
```

#### 响应

```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "宫保鸡丁",
  "type": "main_food",
  "description": "经典川菜，口感麻辣鲜香",
  "created_at": "2024-01-01T12:00:00Z"
}
```

### 2. 获取随机偏好

**GET** `/preferences/random?type=main_food`

根据类型随机获取一个偏好。

#### 查询参数

- `type` (必需): 偏好类型

#### 响应

```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "宫保鸡丁",
  "type": "main_food",
  "description": "经典川菜，口感麻辣鲜香",
  "created_at": "2024-01-01T12:00:00Z"
}
```

### 3. 根据类型获取偏好列表

**GET** `/preferences/type?type=main_food`

获取指定类型的所有偏好。

#### 查询参数

- `type` (必需): 偏好类型

#### 响应

```json
{
  "preferences": [
    {
      "id": "507f1f77bcf86cd799439011",
      "name": "宫保鸡丁",
      "type": "main_food",
      "description": "经典川菜，口感麻辣鲜香",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 4. 获取所有偏好

**GET** `/preferences`

获取所有偏好记录。

#### 响应

```json
{
  "preferences": [
    {
      "id": "507f1f77bcf86cd799439011",
      "name": "宫保鸡丁",
      "type": "main_food",
      "description": "经典川菜，口感麻辣鲜香",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 5. 删除偏好

**DELETE** `/preferences/delete?id=507f1f77bcf86cd799439011`

删除指定的偏好记录。

#### 查询参数

- `id` (必需): 偏好ID

#### 响应

```json
{
  "success": true,
  "message": "删除成功"
}
```

## 错误响应

当发生错误时，API 会返回相应的 HTTP 状态码和错误信息：

```json
{
  "error": "错误描述"
}
```

### 常见状态码

- `200` - 成功
- `201` - 创建成功
- `400` - 请求参数错误
- `404` - 资源不存在
- `500` - 服务器内部错误

## 使用示例

### cURL 示例

```bash
# 创建偏好
curl -X POST http://localhost:8080/api/preferences \
  -H "Content-Type: application/json" \
  -d '{"name":"宫保鸡丁","type":"main_food","description":"经典川菜"}'

# 获取随机主食
curl "http://localhost:8080/api/preferences/random?type=main_food"

# 获取所有主食
curl "http://localhost:8080/api/preferences/type?type=main_food"

# 删除偏好
curl -X DELETE "http://localhost:8080/api/preferences/delete?id=507f1f77bcf86cd799439011"
```

### JavaScript 示例

```javascript
// 创建偏好
const createPreference = async (name, type, description) => {
  const response = await fetch('/api/preferences', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ name, type, description }),
  });
  return response.json();
};

// 获取随机偏好
const getRandomPreference = async (type) => {
  const response = await fetch(`/api/preferences/random?type=${type}`);
  return response.json();
};
``` 