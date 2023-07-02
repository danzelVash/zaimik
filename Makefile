.SILENT:

db:
	docker-compose up -d --build
run-local:
	go run cmd/app/main.go
run-deploy:
	sudo docker-compose up
swagger:
	swag init -g cmd/app/main.go
clean:
	docker-compose down && sudo docker image prune
clean-all:
	docker-compose down && sudo docker image prune && sudo rm -r pgdata