OPTS?=GO111MODULE=on

build: ## Build enkrypt binary
	${OPTS} go build

clean: ## Clean compiled binary
	rm -rf ./enkrypt

build_proto:
	protoc -I. --go_out=plugins=grpc:. \
	  pkg/proto/enkrypt.proto
