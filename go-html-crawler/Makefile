.PHONY: build
build:
	@go build -o app

.PHONY: build
run: build
	@./app

.PHONY: get_chat_ids
get_chat_ids:
	@curl -s "https://api.telegram.org/bot$(TELEGRAM_BOT_TOKEN)/getUpdates"