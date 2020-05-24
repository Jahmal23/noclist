echo "<---- Dep ensuring dependencies before copying files to containers ---->"
dep ensure

echo "<---- Starting Up Docker Containers ---->"
docker-compose up -d --build --remove-orphans --force-recreate

echo "<---- Containers Started, Waiting for warm up ---->"
sleep 5

echo "<---- Running tests ---->"
go test -race -coverprofile=c.out ./... && go tool cover -html=c.out -o coverage.html

echo "<---- Tests complete, stopping containers ---->"
docker-compose down

echo "<---- Cleaning your docker images you're welcome ---->"
docker image prune -f

echo "<---- Go Vet Scan ---->"
go vet  ./...
