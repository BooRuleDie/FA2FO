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
	
publish-topic: build
	@if [ -n "$(filter kern.error user.error mail.error kern.critical user.critical mail.critical,$(word 2,$(MAKECMDGOALS)))" ]; then \
		./build/app publish $(word 2,$(MAKECMDGOALS)) "$(wordlist 3,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))"; \
	else \
		echo "ERROR: First argument must be kern.error, user.error, mail.error, kern.critical, user.critical or mail.critical"; \
		exit 1; \
	fi

consume-topic: build
	@./build/app consume $(filter-out consume-topic build,$(MAKECMDGOALS))
	
rpc-client: build
	@./build/app rpc-client $(filter-out rpc-client build,$(MAKECMDGOALS))
	
rpc-server: build
	@./build/app rpc-server

%:
	@:

.PHONY: build install publish consume publish-direct consume-direct publish-topic consume-topic
