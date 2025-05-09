package proto

//go:generate protoc -I sports/proto --go_out=sports/proto --go_opt=paths=source_relative --go-grpc_out=sports/proto --go-grpc_opt=paths=source_relative sports/proto/sports/sports.proto
