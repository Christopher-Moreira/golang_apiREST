.PHONY: default run build test docs clean

#vars
APP_NAME = gopportunities

#tasks
default: run

run:
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)