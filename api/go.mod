module git.neds.sh/matty/entain/api

go 1.23.0

toolchain go1.24.1

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	google.golang.org/genproto/googleapis/api v0.0.0-20250404141209-ee84b53bf3d0
	google.golang.org/grpc v1.71.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250404141209-ee84b53bf3d0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/SylvanSol/Entain_Test/sports => ../sports
