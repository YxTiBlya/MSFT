
gen-proto:
	mkdir -p pkg/services/restaurant
	protoc	\
		-I api/services/restaurant \
		--go_out=./ --go-grpc_out=./ \
		api/services/restaurant/*.proto