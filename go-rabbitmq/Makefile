build:
	@go build -o build/app

publish: build
	@./build/app publish

consume: build
	@./build/app consume

install:
	@go get github.com/rabbitmq/amqp091-go

.PHONY: build install publish consume
