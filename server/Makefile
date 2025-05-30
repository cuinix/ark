.PHONY: build clean cov help intergrationtest lint run run-neutrino run-neutrino-mutinynet test vet

define setup_env
    $(eval include $(1))
    $(eval export)
endef

## build: build for all platforms
build:
	@echo "Building arkd binary..."
	@bash ./scripts/build

## clean: cleans the binary
clean:
	@echo "Cleaning..."
	@go clean

## cov: generates coverage report
cov:
	@echo "Coverage..."
	@go test -cover ./...

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## intergrationtest: runs integration tests
integrationtest:
	@echo "Running integration tests..."
	@go test -v -count 1 -timeout 500s github.com/ark-network/ark/server/test/e2e

## lint: lint codebase
lint:
	@echo "Linting code..."
	@golangci-lint run --fix

## run: run arkd in dev mode with bitcoind on regtest
run: clean
	@echo "Running arkd with Bitcoin Core in dev mode on regtest ..."
	$(call setup_env, envs/bitcoind.regtest.env)
	go run ./cmd/arkd

## run-neutrino: run arkd in dev mode with neutrino on regtest
run-neutrino: clean
	@echo "Running arkd with Neutrino in dev mode on regtest ..."
	$(call setup_env, envs/neutrino.regtest.env)
	go run ./cmd/arkd

## run-neutrino-mutinynet: run arkd in dev mode with neutrino on mutinynet
run-neutrino-mutinynet: clean
	@echo "Running arkd with Neutrino in dev mode on mutinynet ..."
	$(call setup_env, envs/neutrino.mutinynet.env)
	go run ./cmd/arkd

## test: runs unit and component tests
test:
	@echo "Running unit tests..."
	@go test -v -count=1 -race ./internal/...
	@find ./pkg -name go.mod -execdir go test -v -count=1 -race ./... \;

## vet: code analysis
vet:
	@echo "Running code analysis..."
	@go vet ./...

## mig_file: creates pg migration file(eg. make FILE=init mig_file)
mig_file:
	@migrate create -ext sql -dir ./internal/infrastructure/db/sqlite/migration/ $(FILE)

## mig_up: creates db schema for provided db path
mig_up:
	@echo "creating db schema..."
	@migrate -database "sqlite://$(DB_PATH)/sqlite.db" -path ./internal/infrastructure/db/sqlite/migration/ up

## mig_down: apply down migration
mig_down:
	@echo "migration down..."
	@migrate -database "sqlite://$(DB_PATH)/sqlite.db" -path ./internal/infrastructure/db/sqlite/migration/ down

## mig_down_yes: apply down migration without prompt
mig_down_yes:
	@echo "migration down..."
	@"yes" | migrate -database "sqlite://path/to/database" -path ./internal/infrastructure/db/sqlite/migration/ down

## vet_db: check if mig_up and mig_down are ok
vet_db: recreatedb mig_up mig_down_yes
	@echo "vet db migration scripts..."

## sqlc: gen sql
sqlc:
	@echo "gen sql..."
	@docker run --rm -v ./internal/infrastructure/db/sqlite:/src -w /src sqlc/sqlc generate