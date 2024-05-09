include .env

run:
	go run cmd/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o main_pinginciuman cmd/main.go

migrate:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=verify-full&rootcert=ap-southeast-1-bundle.pem" -path migrations up

rollback:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=verify-full&rootcert=ap-southeast-1-bundle.pem" -path migrations down

migrate-dev:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path migrations up

rollback-dev:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path migrations down

gen-swagger:
	swag init -g cmd/main.go -output cmd/docs

db-up:
	docker compose -f ./deploy/docker-compose.yaml up -d

db-down:
	docker compose -f ./deploy/docker-compose.yaml down