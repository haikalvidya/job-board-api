run-docker:
	@echo "Running docker server"
	@docker-compose up -d

run-migration:
	@echo "Running migrations"
	@go run main.go -migrate

run-local:
	@echo "Running local server"
	@go mod tidy
	@go run main.go
