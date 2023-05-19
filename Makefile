gen-proto-linux:
	mkdir -p pkg/services/restaurant
	protoc	\
		-I api/services/restaurant \
		-I third_party \
		--go_out=./pkg/services/restaurant --go-grpc_out=./pkg/services/restaurant \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/services/restaurant --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/services/restaurant \
		api/services/restaurant/*.proto 

	mkdir -p pkg/services/statistics
	protoc	\
		-I api/services/statistics \
		-I third_party \
		--go_out=./pkg/services/statistics --go-grpc_out=./pkg/services/statistics \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/services/statistics --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/services/statistics \
		api/services/statistics/*.proto 
	
	mkdir -p pkg/services/customer
	protoc	\
		-I api/services/customer \
		-I third_party \
		--go_out=./pkg/services/customer --go-grpc_out=./pkg/services/customer \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/services/customer --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/services/customer \
		api/services/customer/*.proto 

gen-proto-win:
	if not exist "pkg\services\restaurant" mkdir pkg\services\restaurant
	protoc	\
		-I api/services/restaurant \
		-I third_party \
		--go_out=./pkg/services/restaurant --go-grpc_out=./pkg/services/restaurant \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/services/restaurant --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/services/restaurant \
		api/services/restaurant/*.proto 
	
	if not exist "pkg\services\statistics" mkdir pkg\services\statistics
	protoc	\
		-I api/services/statistics \
		-I third_party \
		--go_out=./pkg/services/statistics --go-grpc_out=./pkg/services/statistics \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/services/statistics --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/services/statistics \
		api/services/statistics/*.proto
	
	if not exist "pkg\services\customer" mkdir pkg\services\customer
	protoc	\
		-I api/services/customer \
		-I third_party \
		--go_out=./pkg/services/customer --go-grpc_out=./pkg/services/customer \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/services/customer --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/services/customer \
		api/services/customer/*.proto 