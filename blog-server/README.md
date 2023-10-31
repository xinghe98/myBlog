---
title: blog v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# blog

> v1.0.0

Base URLs:

* <a href="127.0.0.1:3001">开发环境: 127.0.0.1:3001</a>

# Authentication

- HTTP Authentication, scheme: bearer

# Default

## GET 测试

GET /test/ping

> Response Examples

> 200 Response

```json
{
  "msg": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» msg|string|true|none||none|

## POST 创建用户

POST /admin/regist

> Body Parameters

```json
{
  "username": "xinghe",
  "password": "9899"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## POST 登录

POST /admin/login

> Body Parameters

```json
{
  "username": "xinghe",
  "password": "9899"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

# 标签

## POST 新建标签

POST /tags/create

> Body Parameters

```json
{
  "name": "python"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 查看所有标签

GET /tags/findall

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## PUT 更新标签

PUT /tags/45

> Body Parameters

```json
{
  "name": "gogo",
  "id": 111
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 404 Response

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|记录不存在|Inline|

### Responses Data Schema

## DELETE 删除标签

DELETE /tags/7

> Body Parameters

```yaml
{}

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 404 Response

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|记录不存在|Inline|

### Responses Data Schema

## GET 查看所有标签及其文章

GET /tags/findwith

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

# 文章

## POST 添加文章

POST /article/create

> Body Parameters

```json
{
  "tags": [
    {
      "name": "go"
    }
  ],
  "title": "分页测试",
  "status": 1,
  "content": "dsafkljwlkf;jdklsa;jfwkl;jfkldsjflasd;jf\n",
  "image": "https://pic.52112.com/2019/07/23/JPG-190723_759/mfvvEHO7sx_small.jpg"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## PUT 更新文章

PUT /article/1

> Body Parameters

```json
{
  "tags": [
    {
      "name": "go"
    }
  ],
  "title": "nice",
  "status": 1,
  "content": "haha",
  "image": "https://aioseo.com/wp-content/uploads/2021/04/how-to-find-and-fix-404-errors-in-wordpress.png"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## DELETE 删除文章

DELETE /article/1

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

# 前前端接口

## GET 查看所有头条文章

GET /headline

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "data": {
    "articles": [
      {
        "ID": 0,
        "CreatedAt": "string",
        "UpdatedAt": "string",
        "DeletedAt": null,
        "Tags": null,
        "title": "string",
        "status": 0,
        "content": "string",
        "image": "string",
        "headimg": "string"
      }
    ]
  },
  "msg": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» articles|[object]|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» Tags|null|true|none||none|
|»»» title|string|true|none||none|
|»»» status|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» image|string|true|none||none|
|»»» headimg|string|true|none||none|
|» msg|string|true|none||none|

## GET 所有文章（无查询参数）

GET /article/findall

该接口可以查询草稿和已发布的所有文章

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

## GET 根据id查找文章

GET /article/30

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "data": {
    "article": {
      "ID": 0,
      "CreatedAt": "string",
      "UpdatedAt": "string",
      "DeletedAt": null,
      "Tags": [
        {
          "ID": 0,
          "name": "string",
          "HasArt": null
        }
      ],
      "title": "string",
      "status": 0,
      "content": "string",
      "image": "string",
      "headimg": "string"
    }
  },
  "msg": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» article|object|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» Tags|[object]|true|none||none|
|»»»» ID|integer|false|none||none|
|»»»» name|string|false|none||none|
|»»»» HasArt|null|false|none||none|
|»»» title|string|true|none||none|
|»»» status|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» image|string|true|none||none|
|»»» headimg|string|true|none||none|
|» msg|string|true|none||none|

## GET 根据tag查找文章

GET /tags/go

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|page|query|string| no |none|
|limit|query|string| no |none|

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "data": {
    "ID": 0,
    "name": "string",
    "HasArt": [
      {
        "ID": 0,
        "CreatedAt": "string",
        "UpdatedAt": "string",
        "DeletedAt": null,
        "Tags": null,
        "title": "string",
        "status": 0,
        "content": "string",
        "image": "string",
        "headimg": "string"
      }
    ]
  },
  "msg": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» ID|integer|true|none||none|
|»» name|string|true|none||none|
|»» HasArt|[object]|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» Tags|null|true|none||none|
|»»» title|string|true|none||none|
|»»» status|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» image|string|true|none||none|
|»»» headimg|string|true|none||none|
|» msg|string|true|none||none|

## GET 根据发布时间查询文章简略信息

GET /article/archive

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "data": {
    "articles": [
      {
        "ID": 0,
        "Title": "string",
        "CreatedAt": "string"
      }
    ]
  },
  "msg": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|object|true|none||none|
|»» articles|[object]|true|none||none|
|»»» ID|integer|true|none||none|
|»»» Title|string|true|none||none|
|»»» CreatedAt|string|true|none||none|
|» msg|string|true|none||none|

# Data Schema

