from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import AnyUrl

class Settings(BaseSettings):
    # Postgres
    POSTGRES_USER: str
    POSTGRES_PASSWORD: str
    POSTGRES_DB: str
    POSTGRES_PORT: int
    POSTGRES_HOST: str
    DATABASE_URL: str 

    # Redis
    REDIS_HOST: str
    REDIS_PORT: int

    # S3 / LocalStack
    AWS_ACCESS_KEY_ID: str
    AWS_SECRET_ACCESS_KEY: str
    AWS_REGION: str
    S3_ENDPOINT_URL: AnyUrl

    # WireMock phishing API
    PHISHING_API_URL: AnyUrl

    # FastAPI serve
    APP_HOST: str = "0.0.0.0"
    APP_PORT: int = 8000

    model_config = SettingsConfigDict(env_file=".env")


settings = Settings()