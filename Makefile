gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean:
	rm -f pb/*.pb.go

run:
	go run main.go

test:
	go test -race -cover ./...