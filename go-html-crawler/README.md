# ÖSYM Checker

This project is a lightweight Go application that continuously monitors the ÖSYM website (`https://ais.osym.gov.tr/`) and sends Telegram notifications based on the content changes or specific conditions.

## Features

- Periodically checks the ÖSYM page
- Sends Telegram notifications using a bot
- Built in pure Go — no external dependencies
- Runs in a `scratch` Docker container for minimal attack surface and size

## Setup

### 1. Clone the project

```bash
git clone https://github.com/your-username/osym-checker.git
cd osym-checker
```

### 2. Configure your `.env`

Create a `.env` file with your Telegram Bot Token and other configs:

```bash
TELEGRAM_BOT_TOKEN=YOUR-TOKEN
TELEGRAM_CHAT_ID=YOUR-CHAT
```

### 3. Build and Run Locally

```bash
make run
```

### 4. Run in Docker

Build the container:

```bash
make docker-build
```

Run the service with Docker Compose:

```bash
make docker-up
```

## Docker Compose Structure

- `volumes:` are used to persist data like `exams.db`
- `.env` is used to inject secrets/environment variables
- Image is built using a multi-stage setup for production-grade deployment

## Why CGO is Disabled?

Originally, this app used the `mattn/go-sqlite3` package, which requires **CGO**.
This means it depends on **C libraries** (like `libsqlite3.so`) that are **not available in minimal containers like `scratch`**.

To make the app fully self-contained and suitable for a `scratch` image:

- We switched to `modernc.org/sqlite`, a **pure Go** SQLite driver
- We disabled CGO using `CGO_ENABLED=0`
- We built a fully static binary
- We run it in a `FROM scratch` Docker image
- The image is ultra-small and secure

## Getting Your Telegram Chat ID

Use this Makefile command to get your chat ID:

```bash
make get-chat-ids
```

This sends a request to the Telegram Bot API to list recent messages and chat IDs.
