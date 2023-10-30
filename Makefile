init:
	go mod download


compile:
	rm -rf ./pb
	protoc -I. --proto_path=./proto --go_out=. ./proto/*.proto