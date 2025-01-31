run-orders:
	@go run orders/*.go

run-kitchen:
	@go run kitchen/*.go

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=common/genproto/orders --go_opt=paths=source_relative \
  	--go-grpc_out=common/genproto/orders --go-grpc_opt=paths=source_relative
