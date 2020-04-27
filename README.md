# Intro to Rest with Golang

This repository is documenting my learning progress in building REST api with golang. The course is available in Udemy with this link :

https://www.udemy.com/course/build-a-restful-api-with-golang-go-programming-language/



# Running

go run main.go


# Testing

curl http://localhost:8080/books

curl http://localhost:8080/books/1

curl -H "Content-Type: application/json" -X POST -d '{"id":5,"title":"Title 5", "author":"Author 5", "year":"2025"}' http://localhost:8080/books
