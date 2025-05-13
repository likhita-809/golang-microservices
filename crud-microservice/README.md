This is  minimal microservices-style CRUD implementation in Golang without using any database.

We'll just keep in-memory storage (using a map) and expose REST APIs for Create, Read, Update, Delete.

We'll use only standard library (no frameworks like Gin), so we clearly see the structure of how a microservice would be organized.


**Key Microservice patterns shown here:**
*Handler layer* -> for request/response logic
*Store layer* -> for persistent abstraction (in-memory here)
*Model layer* -> for data types
*Router* -> endpoint routing


**Compiling proto files**
```bashrc
protoc --go_out=. --go-grpc_out=. proto/user.proto
```


NOTE: 
- proto folder is for gRPC version
- FROM gcr.io/distroless/base in a Dockerfile sets the base image for your containerization process.
