init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1


compile:
	rm -rf ./pb
	protoc -I. --go_out=. proto/*.proto