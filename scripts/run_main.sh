dep ensure

docker-compose up -d --build --remove-orphans --force-recreate

sleep 2

go run cmd/main.go

docker-compose down




