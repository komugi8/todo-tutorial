MIGRATE_CMD = migrate create -ext sql -dir migrations -seq

.PHONY: migrate

migrate/create:
	@if [ -z "$(file)" ]; then \
		echo "Error: file is not set. Usage: make migrate file=<filename>"; \
		exit 1; \
	fi
	$(MIGRATE_CMD) $(file)

%:
	@$(MAKE) migrate file=$@

migrate:
	migrate -path ./migrations -database 'mysql://root:password@tcp(localhost:3306)/db' up

migrate/down:
	@if [ -z "$(version)" ]; then \
		echo "Error: version is not set. Usage: make migrate/force version=<version>"; \
		exit 1; \
	fi
	migrate -path ./migrations -database 'mysql://root:password@tcp(localhost:3306)/db' down $(version)

migrate/force:
	@if [ -z "$(version)" ]; then \
		echo "Error: version is not set. Usage: make migrate/force version=<version>"; \
		exit 1; \
	fi
	migrate -path ./migrations -database 'mysql://root:password@tcp(localhost:3306)/db' force $(version)