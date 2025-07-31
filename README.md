# Go gRPC Server

Langkah awal setup:

```bash
go mod init grpc-course-protobuf
go get google.golang.org/protobuf
go get google.golang.org/grpc
go run main.go
```

# Go gRPC Client
```bash
go run main.go
```
kemudian jalankan client di terminal baru:
```bash
go run grpcclient/main.go
```

# Client-Streaming
```bash
go run main.go
```
kemudian jalankan client di terminal baru:
```bash
go run grpcclient/main.go
```

# Server-streaming RPC
```bash
go run main.go
```
kemudian jalankan client di terminal baru:
```bash
go run grpcclient/main.go
```

# Bidirectional streaming RPC
```bash
go run main.go
```
kemudian jalankan client di terminal baru:
```bash
go run grpcclient/main.go
```

## Mengaktifkan penjelasan ketika hover (install gopls)
```bash
go install golang.org/x/tools/gopls@latest
```

# gRPC Status Code (Praktik)
```bash
go run main.go
```
kemudian jalankan client di terminal baru:
```bash
go run grpcclient/main.go
```

# gRPC Response Wrapper
setelah mengisi variable di CreateResponse, jalankan
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto     
```
menjalankan dari server:
```bash
go run main.go
```

dan menjalankan dari client:
```bash
go run grpcclient/main.go
```

## membuat base_response.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative common/base_response.proto     
```
dan generate ulang user
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto     
```

kemudian jalankan lagi dari server:
```bash
go run main.go
```

dan juga dari client:
```bash
go run grpcclient/main.go
```

# Protobuf - Timestamp (Praktik)
copy package timestamp, dan pindahkan ke direktori seperti file timestamp.proto:

```bash
https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/timestamp.proto

```

kemudian di file timestamp.proto, pada go_package, copy URL nya dan jalankan
```bash
go get google.golang.org/protobuf/types/known/timestamppb
```

kemudian generete user dengan menjalankan:
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto     
```
ketinggalan, jangan lupa menjalankan run
```bash
go run main.go
```

# gRPC Input Validation (Praktik)
```bash
https://github.com/bufbuild/protovalidate/blob/main/proto/protovalidate/buf/validate/validate.proto

```
dokumentasinya
```bash
https://github.com/bufbuild/protovalidate/tree/main

```
atau lengkapnya
```bash
https://buf.build/bufbuild/protovalidate/docs/main:buf.validate

```
generate ulang user
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto
```

menambah library validate 1 lagi:
```bash
# go get github.com/bufbuild/protovalidate-go

```
jika gagal tambahkan link ini
```bash
go get buf.build/go/protovalidate

```

jalankan main.go 
```bash
go run main.go
```

menambahkan valitation pada balance

kemudian generate ulang:
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user/user.proto
```

jalankan ulang main.go 
```bash
go run main.go
```

setelah mengubah response wrapper, jalankan lagi
```bash
go run main.go
```

# Go gRPC Server Middleware
- ChainUnaryInterceptor: digunakan utk middleware dengan API bersifat Unary
- ChainUnaryInterceptor: digunakan utk middleware dengan API bersifat client Streaming, server streaming dan bidirectional
```bash
https://github.com/bufbuild/protovalidate/blob/main/proto/protovalidate/buf/validate/validate.proto

```