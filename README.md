# go-grpc-boilerplate

## Recreation

- Fork/copy from [`go-gin-boilerplate`](https://github.com/madevara24/go-gin-boilerplate)
- Install protobuff
  - On mac run `brew install protobuf`
  - TODO: Find out how to on windows
- Verify Installation
  - On mac run `protoc --version`
  - TODO: Find out how to on windows
- Install go plugin for protobuf
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- Make sure your `$GOPATH/bin` is in your `$PATH`
  - On mac run `export PATH="$PATH:$(go env GOPATH)/bin"`
- Install grpc dependencies
  - `go get google.golang.org/grpc`
  - `go get google.golang.org/protobuf/reflect/protoreflect`
  - `go get google.golang.org/protobuf/runtime/protoimpl`
- Create folder/directory named `proto` for proto schema on project root
- Create `proto/domain_name/domain_name.proto` file split per domain
- Generate go proto files
  - `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/auth/auth.proto`
