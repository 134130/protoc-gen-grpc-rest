build : clean
	go build -o protoc-gen-grpc-rest ./generate/main.go

clean :
	rm -rf examples/pkg
	rm -rf protoc-gen-grpc-rest
