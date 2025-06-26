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

