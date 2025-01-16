server:
    cd server && go run cmd/main.go

docker:
    docker-compose up -d

web:
    cd web && npm run start
