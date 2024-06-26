﻿# Final Project Sanbercode Golang
Ini merupakan API untuk aplikasi sosial media sederhana. API ini dibuat untuk memenuhi project akhir pada Bootcamp Golang Sanbercode.


# Base
`https://final-project-sanbercode-golang-production.up.railway.app`

# Authorization
Authorization using JWT token

### Protected API
| Method | Endpoint                 | 
|-----------------|-----------------|
| POST  | `/feeds`                    |
| PUT   | `/feeds/:id`                |
| DELETE| `/feeds/:id`                |
| POST  | `/comments`                 |
| PUT   | `/comments/:id`             |
| DELETE| `/comments/:id`             |
| POST  | `/likes`                    |

# Endpoints
## Users

### Register

- URL
  - `/register`
- Method
  - POST
- Request Body
  - `name` as `string`
  - `email` as `string`
  - `password` as `string`
- Response
``` 
{
    "success": true,
    "message": "Berhasil Terdaftar",
    "data": {}
} 
```

### Login

- URL
  - `/login`
- Method
  - POST
- Request Body
  - `email` as `string`
  - `password` as `string`
- Response
``` 
{
    "success": true,
    "message": "Berhasil Masuk",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsInBhc3N3b3JkIjoiIiwiZXhwIjoxNzE3MTY5MDAyLCJpYXQiOjE3MTcwODI2MDIsImlzcyI6ImdvbGFuZ19zYW5iZXJjb2RlIn0.3v7BoBy1hoCdq_ttOuGyZO-_DOeZxUDgvmsCJkoso8k"
    }
}
```
## Feeds
### Create Feed
- URL
  - `/feeds`
- Method
  - POST
- Request Body
  - `message` as `string`
- Response
``` 
{
    "success": true,
    "message": "Berhasil membuat feed",
    "data": {}
}
```

### Update Feed
- URL
  - `/feeds:id`
- Method
  - PUT
- Request Body
  - `message` as `string`
- Params
  - `id` as `int`
- Response
``` 
{
    "success": true,
    "message": "Berhasil membuat feed",
    "data": {}
}
```
### See All Feed
- URL
  - `/feeds`
- Method
  - GET
- Response
``` 
{
    "success": true,
    "message": "Berhasil menampilkan feed",
    "data": [
        {
            "id": 1,
            "message": "Thread Pertama",
            "created_at": "2024-05-31T00:31:46.543625Z",
            "total_comments": 0,
            "total_likes": 0,
            "user": {
                "id": 1,
                "name": "Lustiyana"
            }
        },
        {
            "id": 2,
            "message": "Thread Ketiga",
            "created_at": "2024-05-31T00:32:04.756731Z",
            "total_comments": 0,
            "total_likes": 0,
            "user": {
                "id": 3,
                "name": "Jane"
            }
        }
    ]
}
```
### See Detail Feed
- URL
  - `/feeds/:id`
- Method
  - GET
- Params
  - `id` as `int`
- Response
``` 
{
    "success": true,
    "message": "Berhasil menampilkan detail feed",
    "data": {
        "id": 1,
        "message": "Thread Pertama",
        "created_at": "2024-05-31T00:31:46.543625Z",
        "total_comments": 1,
        "total_likes": 1,
        "comments": [
            {
                "id": 1,
                "message": "comment 1",
                "created_at": "2024-05-31T00:36:23.810539Z",
                "user": {
                    "id": 2,
                    "name": "JohnDoe"
                }
            }
        ],
        "likes": [
            {
                "id": 1,
                "created_at": "2024-05-31T00:35:51.890493Z",
                "user": {
                    "id": 2,
                    "name": "JohnDoe"
                }
            }
        ],
        "user": {
            "id": 1,
            "name": "Lustiyana"
        }
    }
}
```
### Delete Feed
- URL
  - `/feeds/:id`
- Method
  - DELETE
- Params
  - `id` as `int`
- Response
``` 
{
    "success": true,
    "message": "Berhasil menghapus feed",
    "data": {}
}
```
## Comments
### Create Comment
- URL
  - `/comments`
- Method
  - POST
- Request Body
  - `feed_id` as `int`
- Response
``` 
{
    "success": true,
    "message": "Berhasil membuat komentar",
    "data": {}
}
```

### Update Comment
- URL
  - `/comments/:id`
- Method
  - PUT
- Request Body
  - `message` as `string`
- Params
  - `id` as `int`
- Response
``` 
{
    "success": true,
    "message": "Berhasil memperbarui komentar",
    "data": {}
}
```
### Delete Comment
- URL
  - `/comments/:id`
- Method
  - DELETE
- Params
  - `id` as `int`
- Response
``` 
{
    "success": true,
    "message": "Berhasil menghapus komentar",
    "data": {}
}
```
## Likes
### Likes
### Like and Unlike
- URL
  - `/likes`
- Method
  - POST
- Request Body
  - `feed_id` as `int`
- Response
  - Like
    ``` 
    {
        "success": true,
        "message": "Berhasil menyukai",
        "data": {}
    }
    ```
  - Unlike
    ``` 
    {
        "success": true,
        "message": "Berhasil menyukai",
        "data": {}
    }
    ```

