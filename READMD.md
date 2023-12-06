##### 编译

```bash
# macos
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o isx main.go

# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o isx.exe main.go

# linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o isx main.go
```