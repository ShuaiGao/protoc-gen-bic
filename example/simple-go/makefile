# protobuf 文件生成
PROTOC_DIR := $(shell pwd)/script/

pb:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go  --go_out=. proto/*.proto --proto_path=proto
	@$(PROTOC_DIR)protoc-go-inject-tag -input="gen/api/*.pb.go"
	@$(PROTOC_DIR)swag fmt
	@$(PROTOC_DIR)swag init

# swag 文档生成
swag:
	@$(PROTOC_DIR)swag init

ts:
	@pbjs -t static-module -w es6 -o ../client/pb/proto.js ./proto/user.proto --keep-case --no-service --no-typeurl --no-verify --no-delimited --no-encode --no-decode --no-convert --force-number
	@pbts -o ../client/pb/proto.d.ts ../client/pb/proto.js
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --bic_out=ts_dir=../client/pb:. proto/*.proto --proto_path=proto

