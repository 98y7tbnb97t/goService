
DB_USER := postgres
DB_PASS := root
DB_HOST := localhost
DB_PORT := 5432
DB_NAME := postgres
DB_DSN := "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

BACKUP_DIR := ./backups
BACKUP_FILE := $(BACKUP_DIR)/backup_$(shell date +%Y%m%d_%H%M%S).sql

$(shell mkdir -p $(BACKUP_DIR))

migrate-new:
	@if [ -z "${NAME}" ]; then echo "Переменная NAME не установлена. Пример: make migrate-new NAME=my_migration_name"; exit 1; fi
	migrate create -ext sql -dir ./migrations ${NAME}

backup:
	@echo "Создание резервной копии базы данных в $(BACKUP_FILE)..."
	PGPASSWORD=$(DB_PASS) pg_dump -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -F c -b -v -f "$(BACKUP_FILE)"
	@echo "Резервная копия создана: $(BACKUP_FILE)"

migrate: backup
	@echo "Применение миграций..."
	$(MIGRATE) up
	@echo "Миграции применены."

migrate-down: backup
	@echo "Откат миграций..."
	$(MIGRATE) down
	@echo "Миграции отменены."

run:
	go run cmd/app/main.go

# Можно указать конкретный файл: make restore FILE=./backups/backup_YYYYMMDD_HHMMSS.sql
restore:
	@LATEST_BACKUP=$$(ls -t $(BACKUP_DIR)/*.sql | head -n 1); \
	RESTORE_FILE=${FILE:-$$LATEST_BACKUP}; \
	if [ -z "$$RESTORE_FILE" ]; then echo "Не найдены файлы бэкапов в $(BACKUP_DIR)"; exit 1; fi; \
	if [ ! -f "$$RESTORE_FILE" ]; then echo "Файл бэкапа не найден: $$RESTORE_FILE"; exit 1; fi; \
	echo "Восстановление базы данных из файла: $$RESTORE_FILE..."; \
	PGPASSWORD=$(DB_PASS) pg_restore -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -c -1 -v "$$RESTORE_FILE"; \
	echo "База данных восстановлена из $$RESTORE_FILE."

clean-backups:
	@echo "Удаление старых бэкапов из $(BACKUP_DIR)..."
	rm -f $(BACKUP_DIR)/*.sql
	@echo "Бэкапы удалены."

gen:
	@echo "Creating directories for generated code..."
	mkdir -p ./internal/web/tasks
	mkdir -p ./internal/web/users
	@echo "Generating code from OpenAPI specification..."
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go
	@echo "Code generation completed."

lint:
	@echo "Запуск линтера..."
	golangci-lint run --out-format=colored-line-number
	@echo "Линтер завершил работу."

.PHONY: migrate-new backup migrate migrate-down restore run clean-backups gen lint
