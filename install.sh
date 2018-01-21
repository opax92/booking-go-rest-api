docker-compose build
docker-compose run backend-service glide install
docker-compose run backend-service go build server.go
