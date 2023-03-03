


run_db:
	docker start employee_management

run:
	go run cmd/server/main.go

.PHONY: run_db, run