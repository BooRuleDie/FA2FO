import os
import sys
import time

from testcontainers.postgres import PostgresContainer
from testcontainers.redis import RedisContainer
from testcontainers.localstack import LocalStackContainer
from testcontainers.core.container import DockerContainer

import pytest

# Add the project root directory to Python path
project_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.path.insert(0, project_root)

# Make sure .env is not overriding our dynamic URLs
os.environ.pop("DATABASE_URL", None)
os.environ.pop("POSTGRES_HOST", None)
os.environ.pop("REDIS_HOST", None)
os.environ.pop("S3_ENDPOINT_URL", None)
os.environ.pop("PHISHING_API_URL", None)


@pytest.fixture(scope="session", autouse=True)
def docker_services():
    # 1) PostgreSQL
    user = os.environ.get("POSTGRES_USER")
    password = os.environ.get("POSTGRES_PASSWORD")
    dbname = os.environ.get("POSTGRES_DB")

    pg = PostgresContainer(
        "postgres:17.4", username=user, password=password, dbname=dbname
    )
    pg.start()
    # Export for SQLAlchemy/Alembic
    os.environ["DATABASE_URL"] = pg.get_connection_url()
    os.environ["POSTGRES_HOST"] = pg.get_container_host_ip()
    os.environ["POSTGRES_PORT"] = str(pg.get_exposed_port(pg.port))

    # 2) Redis
    rd = RedisContainer("redis:7-alpine")
    rd.start()
    os.environ["REDIS_HOST"] = rd.get_container_host_ip()
    os.environ["REDIS_PORT"] = str(rd.get_exposed_port(6379))

    # 3) LocalStack (S3)
    ls = LocalStackContainer("localstack/localstack:2.2").with_services("s3")
    ls.start()
    h, p = ls.get_container_host_ip(), ls.get_exposed_port(4566)
    # os.environ["AWS_ACCESS_KEY_ID"]     = "test"
    # os.environ["AWS_SECRET_ACCESS_KEY"] = "test"
    # os.environ["AWS_REGION"]            = "us-east-1"
    os.environ["S3_ENDPOINT_URL"] = f"http://{h}:{p}"

    # 4) WireMock
    #    We assume you have put your mappings under ./wiremock/mappings
    wm = (
        DockerContainer("wiremock/wiremock:2.35.0")
        .with_exposed_ports(8080)
        .with_command(["--verbose", "--global-response-templating"])
        .with_volume_mapping(
            os.path.abspath("wiremock"),
            "/home/wiremock/mappings",
        )
    )
    wm.start()
    h2, p2 = wm.get_container_host_ip(), wm.get_exposed_port(8080)
    os.environ["PHISHING_API_URL"] = f"http://{h2}:{p2}/phishing"

    # print("Waiting 8 seconds for services to be ready...")
    time.sleep(8)

    # Print all updated environment variables
    print("Updated environment variables:")
    updated_vars = [
        "POSTGRES_USER",
        "POSTGRES_PASSWORD",
        "POSTGRES_DB",
        "POSTGRES_HOST",
        "POSTGRES_PORT",
        "REDIS_HOST",
        "REDIS_PORT",
        # "AWS_ACCESS_KEY_ID",
        # "AWS_SECRET_ACCESS_KEY",
        # "AWS_REGION",
        "S3_ENDPOINT_URL",
        "PHISHING_API_URL",
    ]
    for var in updated_vars:
        print(f"{var}={os.environ.get(var)}")

    yield

    # Teardown
    wm.stop()
    ls.stop()
    rd.stop()
    pg.stop()


@pytest.fixture()
def client():
    # import here so that the env vars are in place first
    from fastapi.testclient import TestClient
    from app.main import app

    # The startup event will create tables automatically
    test_client = TestClient(app)

    # Force lifespan startup by using the client in a context
    with test_client as client_context:
        # Test if the client is working
        print("Testing /health/status endpoint...")
        try:
            health_response = client_context.get("/health/status")
            print(f"Health check status: {health_response.status_code}")
            # print(f"Health response: {health_response.json()}")
        except Exception as e:
            print(f"Health check failed: {e}")
            raise Exception(e)

        # Return the client for use in tests
        return client_context
