.PHONY: migrate
migrate:
	migrate -source file://migrations -database postgres://postgres:qwertyQAZ12345!@localhost:5433/CoffeeDb?sslmode=disable up
migrate-down:
	migrate -source file://migrations -database postgres://postgres:qwertyQAZ12345!@localhost:5433/CoffeeDb?sslmode=disable  down
migrate-status:
	migrate -database  postgres://postgres:qwertyQAZ12345!@localhost:5433/CoffeeDb?sslmode=disable -path file://migrations statu