APP_NAME=user-service
PORT=8080

build: 
	go build -o $(APP_NAME) ./cmd/main.go

run: 
	go run ./cmd/main.go

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):$(PORT) $(APP_NAME):latest