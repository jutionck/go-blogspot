## Go-Chi & Gorm

### Tool
1. **Go-Chi** https://go-chi.io/#/
2. **Gorm** https://gorm.io


### Database
- Sqlite


### Run
```bash
PORT=8888 go run cmd/main.go
```

### BlogSpot

#### Create Blog
`POST` `/posts`
- request
```json
{
    "title": "Belajar Chii",
    "content": "Body content"
}
```
- response
```json
{
    "id": 1,
    "title": "Belajar Chii",
    "content": "Body content"
}
```


#### List Blog
`GET` `/posts`
- response
```json
[
    {
        "id": 1,
        "title": "Belajar chi",
        "content": "Body content"
    },
    {
        "id": 2,
        "title": "Belajar Gorm",
        "content": "Body content"
    }
]
```