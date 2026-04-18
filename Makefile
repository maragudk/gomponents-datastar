.PHONY: deps
deps:
	curl -Lf -o docs/datastar.js https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0/bundles/datastar.js

.PHONY: benchmark
benchmark:
	go test -bench . ./...

.PHONY: cover
cover:
	go tool cover -html cover.out

.PHONY: demo
demo:
	cd demo && go run .

.PHONY: fmt
fmt:
	goimports -w -local `head -n 1 go.mod | sed 's/^module //'` .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -coverprofile cover.out -shuffle on ./...
