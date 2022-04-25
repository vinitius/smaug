help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

mock-generate: ## Generate mocks

tools: ## Basic helper tools
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest

fmt: ## Format files
	goimports -l -w .
	gofumpt -l -w .

test: ## Run unit tests
	touch count.out
	go test -covermode=count -coverprofile=count.out -v ./...
	$(MAKE) coverage

coverage: ## Unit tests coverage
	go tool cover -func=count.out

lint: ## Run linter
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.44.0 golangci-lint run -v

create-env: ## Create sample env file locally
	cp .env.docker.sample .env # Create .env file

.PHONY: test
