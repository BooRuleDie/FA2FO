.PHONY: build
build:
	@go build -o app

.PHONY: build
run: build
	@./app

.PHONY: get-chat-ids
get-chat-ids:
	@curl -s "https://api.telegram.org/bot$(TELEGRAM_BOT_TOKEN)/getUpdates"

.PHONY: docker-build
docker-build:
	@docker build -t osym-checker:latest .

.PHONY: docker-up
docker-up:
	@docker compose up -d
