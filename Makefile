.PHONY:  tools lint


# Tools required
TOOLS = github.com/golangci/golangci-lint/cmd/golangci-lint

# Lint
lint: tools
	PATH=$(PATH) golangci-lint run

# test
test:
	PATH=$(PATH) go test ./...

# run
run:
	PATH=$(PATH) go run cmd/simplecomputation-service/main.go


# run client
run-client:
	PATH=$(PATH) go run client/client.go





