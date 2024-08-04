build-proto:
	protoc \
	--go_out=client_pub_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=client_pub_proto \
	--go-grpc_opt=paths=source_relative \
	client_pub.proto
