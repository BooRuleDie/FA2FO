import os
import sys
from logging.config import fileConfig

from alembic import context
from sqlalchemy import engine_from_config, pool

config = context.config

if config.config_file_name is not None:
    fileConfig(config.config_file_name)

sys.path.insert(0, os.getcwd())

from app.models import Base  # noqa
target_metadata = Base.metadata

db_url = os.getenv("DATABASE_URL", "").strip()
if not db_url:
    user = os.getenv("POSTGRES_USER")
    pw = os.getenv("POSTGRES_PASSWORD")
    host = os.getenv("POSTGRES_HOST", "localhost")
    port = os.getenv("POSTGRES_PORT", "5432")
    db   = os.getenv("POSTGRES_DB")
    if not all([user, pw, db]):
        raise RuntimeError(
            "Neither DATABASE_URL nor POSTGRES_USER/PASSWORD/DB are fully set in environment"
        )
    db_url = f"postgresql+psycopg2://{user}:{pw}@{host}:{port}/{db}"

config.set_main_option("sqlalchemy.url", db_url)

def run_migrations_offline() -> None:
    """Run migrations in 'offline' mode."""
    url = config.get_main_option("sqlalchemy.url")
    context.configure(
        url=url,
        target_metadata=target_metadata,
        literal_binds=True,
        dialect_opts={"paramstyle": "named"},
    )

    with context.begin_transaction():
        context.run_migrations()


def run_migrations_online() -> None:
    """Run migrations in 'online' mode."""
    connectable = engine_from_config(
        config.get_section(config.config_ini_section),
        prefix="sqlalchemy.",
        poolclass=pool.NullPool,
    )
    with connectable.connect() as connection:
        context.configure(
            connection=connection,
            target_metadata=target_metadata,
        )
        with context.begin_transaction():
            context.run_migrations()


if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()