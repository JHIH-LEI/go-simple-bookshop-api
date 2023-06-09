## init

go mod init <module-name>

## 安裝dependency
go get -u github.com/gin-gonic/gin

## start
go run main.go

## test

### GET /books

curl localhost:8080/books
```
[
    {
        "id": 1,
        "title": "前端三十",
        "author": "朱信穎",
        "Quantity": 10
    },
    {
        "id": 2,
        "title": "精通Python",
        "author": "Bill Lubanovic",
        "Quantity": 1
    }
]
```

### POST /books

curl localhost:8080/books --include --header "Content-Type: application/json" -d @create-book.json --request "POST"

response:
```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 02 Apr 2023 17:04:15 GMT
Content-Length: 81

{
    "id": 3,
    "title": "新書",
    "author": "Alicia",
    "Quantity": 1
}%         
```

### GET /books/:id
成功狀況：
curl localhost:8080/books/1

response:
```
{
    "id": "1",
    "title": "前端三十",
    "author": "朱信穎",
    "Quantity": 10
}
```
失敗狀況：
```
curl localhost:8080/books/11
{
    "message": "Book not found."
}
```