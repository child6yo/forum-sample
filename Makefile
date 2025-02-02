build:
	docker-compose build forum-service

run:
	docker-compose up forum-service

test:
	go test -v ./...

migrate:
	migrate -path ./schemas -database 'postgres://postgres:Qwerty@db:5433/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go