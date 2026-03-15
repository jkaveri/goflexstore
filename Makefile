lint:
	golangci-lint run
.PHONY: lint

lint-fix:
	golangci-lint run --fix
.PHONY: lint-fix

test:
	./scripts/test.sh
.PHONY: test

mock:
	rm -rf ./mocks && mockery
.PHONY: mock