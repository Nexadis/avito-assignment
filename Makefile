GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

help:
	@echo "This is a helper makefile for Avito banner service"
	@echo ""
	@echo "Targets:"
	@echo "	generate:	regenerate all generated files"
	@echo "	test:		run all tests"
	@echo "	lint:		lint the project"
	@echo "	fmt:		all project"

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.57.2

$(GOBIN)/oapi-codegen:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: tools
tools: $(GOBIN)/golangci-lint $(GOBIN)/oapi-codegen


lint: tools
	$(GOBIN)/golangci-lint run ./...

lint-ci: tools
	$(GOBIN)/golangci-lint run ./... -out-format=github-actions --timeout=5m

generate:
	go generate ./...

test:
	go test -cover ./...

tidy:
	go mod tidy

tidy-ci:
	tidied -verbose
