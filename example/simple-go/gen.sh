#!/usr/bin/env sh
protoc --go_out=simple-go proto/*.proto --proto_path=proto --bic_out=simple-go

# 格式化生成的swag
# swag fmt
	
pbjs -t static-module -w es6 -o ../client/pb/proto.js ./proto/user.proto --keep-case --no-service --no-typeurl --no-verify --no-delimited --no-encode --no-decode --no-convert --force-number
pbts -o ../client/pb/proto.d.ts ../client/pb/proto.js
protoc --plugin=protoc-gen-bic=/Users/ziga/go/bin/protoc-gen-bic --plugin=protoc-gen-go=/Users/ziga/go/bin/protoc-gen-go --bic_out=ts_dir=../client/pb:. ./proto/*.proto --proto_path=./proto
