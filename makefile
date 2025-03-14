# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
#DB_DSN := "host=localhost user=postgres password=pantera dbname=postgres port=5432 sslmode=disable"
DB_DSN := "postgres://postgres:pantera@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)



migrate-f:
	$(MIGRATE) force $(V)
# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go


gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	@echo "Starting the application..."
	go run cmd/app/main.go

lint:
	golangci-lint run --out-format=colored-line-number

