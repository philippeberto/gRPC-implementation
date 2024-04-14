# Basic gRPC implementation

## How to run

To install gRPC
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28  
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```
To install evans
```
brew tap ktr0731/evans                                         
brew install evans
```
To start services
````
docker-compose up -d
````
To create table in database
````
sqlite3 db.sqlite  
create table categories (id string, name string, description string);
````
To generate protocol buffer results
````
protoc --go_out=. --go-grpc_out=. ./proto/course_category.proto
````
To acess evans
````
evans -r repl
````
To acess resources and call service inside evans
````
package pb
service CategoryService
  call CreateCategory
  call ListCategories
````
