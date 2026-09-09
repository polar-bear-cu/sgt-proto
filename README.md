# Subglutee Project - Proto

gRPC Repo สำหรับ Project Subglutee

### Prerequisite

1. Go 1.26
2. VSCode Extension - Protobuf
3. protoc plugins + buf + lefthook

```terminal
winget install Google.Protobuf
go install github.com/bufbuild/buf/cmd/buf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.12
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.2
go install github.com/evilmartians/lefthook@latest
```

### Setup

```terminal
git clone https://github.com/polar-bear-cu/sgt-proto.git
cd sgt-proto
lefthook install
go mod download
buf generate
go build ./...
```

### Useful Commands

```terminal
buf generate        # regen gen/go
buf lint            # lint check
buf format -w       # format check
```
