BIN= $(CURDIR)/bin
Q= $(if $(filter 1,$V),,@)
M= $(shell printf "\033[34;1m▶\033[0m")
PKGS = $(or $(PKG),$(shell GO111MODULE=on go list ./...))

.PHONY: build
build: fmt $(BIN) ; $(info $(M) building executable...) @
								$Q @env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/insert-contact-aws-lambda cmd/main.go

lint: ; $(info $(M) running linter...) @
								$Q @go vet ./...

fmt: ; $(info $(M) running gofmt...) @
								$Q @go fmt ./...

clean:
	@rm -rf $(BIN)