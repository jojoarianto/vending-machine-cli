
build: ## build binary file
	go build main.go

dev: ## run on development mode
	go run main.go

test: ## run go test
	go test ./... -v

help: ## display help command
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'