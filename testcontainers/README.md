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

# FastAPI Dev App
```bash
uvicorn app.main:app --host 127.0.0.1 --port 8000 --reload
```

# Testing
```bash
pytest -v
```