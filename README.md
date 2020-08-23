# golang-mongodb-restapi
Rest API using GoLang and MongoDB

# How to run the application in local

install golang

install mongodb

create db go-mongo

create collection articles

git clone the repo to the installed go dir (Eg: /c/Users/jaison/go/src)

cd to the root dir (Eg: /c/Users/jaison/go/src)

Run the go command >> go build

Run the go command >> go run main.go

# API details

Port: 8081

http://localhost:8081/api/articles [GET]
For listing all the articles

http://localhost:8081/api/articles [POST]
For creation of articles

Validations: check for empty fields

http://localhost:8081/api/articles/{id} [GET]
For getting single article

Validations:
Check for id and throw error for invalid id

# Unit Testing
 
 cd to controllers dir
 run go test -v

 Unit Test created for:
 
 CreateArtice TestCreateArticle

 GetArticle TestGetArticle
 
 GetArticles TestGetArticles

 # Docker

Docker and docker-compose.yaml in the root of the project