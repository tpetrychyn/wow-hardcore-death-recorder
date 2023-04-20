.PHONY: unit-test
unit-test:
	@go test --failfast ./...

.PHONY: run-flights-api
run-flights-api:
	@export DOT_ENV_PATH=cmd/player_deaths/api/example.env && \
	go run cmd/player_deaths/api/main.go

