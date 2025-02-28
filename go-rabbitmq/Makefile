install:
	@go get github.com/rabbitmq/amqp091-go

build:
	@go build -o build/app

publish: build
	@./build/app publish

consume: build
	@./build/app consume

publish-direct: build
	@if [ -n "$(filter error info warning,$(word 2,$(MAKECMDGOALS)))" ]; then \
		./build/app publish $(word 2,$(MAKECMDGOALS)) "$(wordlist 3,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))"; \
	else \
		echo "ERROR: First argument must be error, info, or warning"; \
		exit 1; \
	fi

consume-direct: build
	@./build/app consume $(filter-out consume-direct build,$(MAKECMDGOALS))

%:
	@:

.PHONY: build install publish consume publish-direct consume-direct
