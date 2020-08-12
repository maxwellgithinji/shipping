# shipping

## Prerequisites

- [protoc compiler](https://grpc.io/docs/protoc-installation/)
- [gRPC protobuf](https://grpc.io/docs/languages/go/)
- [Basic knowledge of microservices](https://www.nginx.com/blog/introduction-to-microservices/)
- [Docker](https://hasura.io/blog/the-ultimate-guide-to-writing-dockerfiles-for-go-web-apps-336efad7012c/)
- [Docker Networking](https://docs.docker.com/network/)
- [Docker-compose](https://docs.docker.com/compose/install/)

## Services
The following are the services and instructions on how to run them

### Shipping service consignment

- `cd shipping-service-consignment`
- `docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shipping-service-consignment`

the -e flag is for the environment variable for the microservice port

### Shipping service vessel

- `cd shipping-service-vessel`
- `docker run -p 50052:50052 -e MICRO_SERVER_ADDRESS=:50052 shipping-service-vessel`

### shipping cli consignment

- `cd shipping-cli-consignment`
- `docker run -p 50053:50053 -e MICRO_SERVER_ADDRESS=:50053 shipping-cli-consignment`


## Generating go-micro files
- ensure you have the following installed
  - `go get github.com/micro/micro/v2`
  - `go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master`

- run `protoc --proto_path=. --go_out=. --micro_out=. proto/<service_name>/<service_name.proto>`

- Note: if you get an error at this point, regarding an etcd dependency, try adding the following to your go.mod file
```
...

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
  ...
)
```

## References

[Ewan Valentine's tutorial](https://ewanvalentine.io/microservices-in-golang-part-0/)