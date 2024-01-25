# protobuf 文件生成
PROTOC_DIR := $(shell pwd)/script/

go:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --bic_out=. --go_out=. proto/*.proto --proto_path=proto
	@$(PROTOC_DIR)gomodifytags -file gen/api/user.pb.go -all -add-tags  form -w -quiet
	@$(PROTOC_DIR)protoc-go-inject-tag -input="gen/api/*.pb.go"
	@$(PROTOC_DIR)swag fmt
	@$(PROTOC_DIR)swag init

# swag 文档生成
swag:
	@$(PROTOC_DIR)swag init

ts:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --bic_out=ts_dir=../client/api_ts:. proto/*.proto --proto_path=proto

js:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --bic_out=js_dir=../client/api_js:. proto/*.proto --proto_path=proto

