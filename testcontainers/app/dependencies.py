import os

from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
import redis
import boto3

from app.config import settings

DATABASE_URL = os.getenv("DATABASE_URL", "")

# echo=True will log all SQL; turn off in prod
engine = create_engine(DATABASE_URL, pool_pre_ping=True, echo=False)

SessionLocal = sessionmaker(
    autocommit=False,
    autoflush=False,
    bind=engine,
)

def get_db():
    """Yield a SQLAlchemy Session, closing it afterwards."""
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


_redis_client = None
def get_redis():
    """Return a singleton Redis client."""
    global _redis_client
    if _redis_client is None:
        _redis_client = redis.Redis(
            host=settings.REDIS_HOST,
            port=settings.REDIS_PORT,
            decode_responses=True,
        )
    return _redis_client



_s3_client = None
def get_s3():
    """Return a singleton boto3 S3 client pointed at LocalStack."""
    global _s3_client
    if _s3_client is None:
        _s3_client = boto3.client(
            "s3",
            aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
            aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
            region_name=settings.AWS_REGION,
            endpoint_url=str(settings.S3_ENDPOINT_URL),
        )
    return _s3_client
