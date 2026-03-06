test:
	go test -v -cover ./internal/handlers
build:
	go build ./cmd/taskmanager
up:
	docker compose up --build -d
down:
	docker compose down