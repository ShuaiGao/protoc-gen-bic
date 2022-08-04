GOPATHS=`go env GOPATH`

windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "protoc-gen-bic" main.go
	@cp  protoc-gen-bic $(GOPATHS)/bin
	@rm protoc-gen-bic
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o "protoc-gen-bic" main.go
	@cp  protoc-gen-bic $(GOPATHS)/bin
	@rm protoc-gen-bic
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -o "protoc-gen-bic" main.go
	@cp  protoc-gen-bic $(GOPATHS)/bin
	@rm  protoc-gen-bic

install:
	@echo $(GOPATHS)
