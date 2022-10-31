go-build: go-install-vendor go-generate # Install vendor and generate mock-easyjson files

go-generate: go-mocks

go-mocks: ## Generate mocks for testing
	@echo "generating mock files ..."
	find $(CURDIR) -name "mock_*.go" -not -path "$(CURDIR)/vendor/*" -delete
	go generate -run="mockgen" ./...
	@echo "... done"

up: ## Create docker's
	docker-compose up -d


go-test: go-mocks ##Test
	go test -v --tags=unit,e2e ./...

go-test-coverage:
	go test -v --tags=unit ./... -covermode=count -coverpkg=./... -coverprofile infrastructure/coverage/coverage.out
	go tool cover -html infrastructure/coverage/coverage.out -o infrastructure/coverage/coverage.html

go-install-vendor: ## Install dependencies
	go mod vendor


go-update-vendor: ## Updates dependencies
	go mod tidy && go mod vendor

go-test-unit:
	go test -v -tags=unit ./...


