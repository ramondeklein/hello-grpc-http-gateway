package helloworld

// Regenerating all files requires the installation of:
//
//	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
//	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
//
//go:generate protoc -I=. -I=./googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto
//go:generate protoc -I=. -I=./googleapis --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true helloworld.proto
//go:generate protoc -I=. -I=./googleapis --openapiv2_out=. helloworld.proto
