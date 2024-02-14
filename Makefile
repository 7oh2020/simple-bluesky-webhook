
.Phony: all
all: build

.Phony: depend
depend:
	@go mod tidy

.Phony: fmt
fmt:
	@go fmt ./...

.Phony: run
run: depend
	@go run .

.Phony: build
build: depend
	@go build .

.Phony: test
test: depend 
	@go test -short ./...

.Phony: integration
integration: depend 
	@go test ./...
