BIN= $(CURDIR)/bin
FUNCTION=insert-contact-aws-lambda

.PHONY: build

build: clean fmt
	@echo "Building $(FUNCTION)..."
	@env GOOS=linux go build -ldflags="-s -w" -o bin/$(FUNCTION) cmd/main.go

lint:
	@go vet ./...

fmt:
	@go fmt ./...

clean:
	@rm -rf $(BIN)

zip: build
	@zip -j $(FUNCTION).zip bin/$(FUNCTION)

.PHONY: test
test:
	@go test -v -cover ./...