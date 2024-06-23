generate:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		api/go_poll/.proto