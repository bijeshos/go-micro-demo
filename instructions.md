
- to generate .pb.go
    -`protoc --proto_path=. --micro_out=. --go_out=. greeter.proto`

- to run the server
    - `go run main.go`
    
- to run the client
    - `go run client.go`
