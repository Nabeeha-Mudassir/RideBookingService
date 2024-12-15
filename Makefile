run-bookings:
	@go run services/bookings/*.go
run-users:
	@go run services/users/*.go
run-rides:
	@go run services/rides/*.go

run-userclient:
	@go run client/users_client/*.go
run-assignment:
	@go run services/assignment/*.go

gen-users:
	@protoc \
		--proto_path=protobuf "protobuf/users.proto" \
		--go_out=services/common/genproto/users --go_opt=paths=source_relative \
  	--go-grpc_out=services/common/genproto/users --go-grpc_opt=paths=source_relative

gen-bookings:
	@protoc \
		--proto_path=protobuf "protobuf/bookings.proto" \
		--go_out=services/common/genproto/bookings --go_opt=paths=source_relative \
  		--go-grpc_out=services/common/genproto/bookings --go-grpc_opt=paths=source_relative

gen-rides:
	@protoc \
		--proto_path=protobuf "protobuf/rides.proto" \
		--go_out=services/common/genproto/rides --go_opt=paths=source_relative \
  		--go-grpc_out=services/common/genproto/rides --go-grpc_opt=paths=source_relative