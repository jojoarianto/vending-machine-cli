

build: ## build binary file
	go build -o main.go vending-machine

run: ## run on development mode
	go run main.go

coverage: ## Generate global code coverage report
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

test: ## run go test
	go test ./... -v

help: ## display help command
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'