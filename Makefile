BIN= $(CURDIR)/bin
Q= $(if $(filter 1,$V),,@)
M= $(shell printf "\033[34;1m▶\033[0m")

.PHONY: clean lint fmt

test:
	@go test -v cover ./...

lint: ; $(info $(M) running linter...) @ ## Run go vet on all source files
								$Q $(MAKE) -C insert-contact-aws-lambda

fmt: ; $(info $(M) running gofmt...) @ ## Run gofmt on all source files
								$Q $(MAKE) -C insert-contact-aws-lambda

clean:
	@rm -rf $(BIN)