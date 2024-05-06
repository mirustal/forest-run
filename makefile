lint: 
	golangci-lint run --config=./golangci.yml  


all: build


build:
	docker build -t $(APP_NAME) .


up:
	docker-compose up --build
	

down:
	docker-compose down


restart: down start


clean:
	docker-compose down --rmi all --remove-orphans
	docker rmi $(APP_NAME)


logs:
	docker-compose logs -f


test:
	docker exec -it $(APP_NAME) go test ./internal/handler/


shell:
	docker exec -it $(APP_NAME) /bin/sh
