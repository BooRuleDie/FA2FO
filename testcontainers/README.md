# Docker
```bash
docker build -t shortener-app:latest .
docker compose up -d
```

# Alembic
```bash
alembic init alembic
alembic revision --autogenerate -m "initial"
alembic upgrade head
```