BIN= $(CURDIR)/bin
Q= $(if $(filter 1,$V),,@)
M= $(shell printf "\033[34;1m▶\033[0m")

.PHONY: build
build: ; $(info $(M) building executables...) @
		$Q $(MAKE) -C insert-contact-aws-lambda build
		$Q $(MAKE) -C get-contact-aws-lambda build
		$Q $(MAKE) -C process-contact-aws-lambda build

test: ; $(info $(M) running tests...) @
		$Q $(MAKE) -C insert-contact-aws-lambda test
		$Q $(MAKE) -C get-contact-aws-lambda test

zip: ; $(info $(M) generating zip files...) @
		$Q $(MAKE) -C insert-contact-aws-lambda zip
		$Q $(MAKE) -C get-contact-aws-lambda zip
		$Q $(MAKE) -C process-contact-aws-lambda zip

clean: ; $(info $(M) cleaning...) @
		$Q $(MAKE) -C insert-contact-aws-lambda clean
		$Q $(MAKE) -C get-contact-aws-lambda clean
		$Q $(MAKE) -C process-contact-aws-lambda clean