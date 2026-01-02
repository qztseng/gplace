.PHONY: lint test coverage

lint:
	golangci-lint run ./...

test:
	go test ./...

coverage:
	./scripts/check-coverage.sh
