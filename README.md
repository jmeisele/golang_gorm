# golang_gorm

Go lang CRUD app with sqlite DB

`go run github.com/jmeisele/golang_gorm`

## Try it out

Open up POSTMAN or Thunder client from vs code and hit our endpoint(s)

Get all books

```
curl http://localhost:8081/book/
```

Create a Book entry

```
curl -X POST http://localhost:8081/book/ -d '{"Name": "Lord of the Rings", "Author": "JRR Tolkien", "Publication": "Harper Collins"}'
```

Get book entry by ID

```
curl http://localhost:8081/book/1
```

Update a book entry

```
curl -X PUT http://localhost:8081/book/1 -d '{"Name": "Lord of the Rings", "Author": "JRR Tolkien", "Publication": "Penguin Press"}'
```

Delete a book entry

```
curl -X DELETE http://localhost:8081/book/1
```
