 5000OPTS?=GO111MODULE=on

build: ## Build enkrypt binary
	${OPTS} go build -o ./enkrypt ./cmd

clean: ## Clean compiled binary
	rm -rf ./enkrypt

build_proto:
	protoc -I. --go_out=plugins=grpc:. \
	  pkg/proto/enkrypt.proto

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -local github.com/Kifen/crypto-watch ./pkg
	goimports -w -local github.com/Kifen/crypto-watch ./cmd

install-linters:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
	golangci-lint --version
