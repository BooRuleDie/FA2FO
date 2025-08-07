from contextlib import asynccontextmanager
from fastapi import FastAPI

from app.routers.shortener import router as shortener_router
from app.routers.stats import router as stats_router
from app.dependencies import engine
from app.models import Base


@asynccontextmanager
async def lifespan(_: FastAPI):
    """
    Lifespan manager for the FastAPI app.
    It now uses the imported `engine` and `Base` to create tables.
    """
    print("Starting up... Creating database tables.")
    # The `Base` object imported from your models now knows about URL, Click, etc.
    # We bind it to the single engine imported from your dependencies.
    Base.metadata.create_all(bind=engine)
    
    yield
    
    print("Shutting down...")


app = FastAPI(
    title="URL Shortener Demo (sync)",
    lifespan=lifespan
)

# Include routers
app.include_router(shortener_router)
app.include_router(stats_router)


@app.get("/health/status")
def health():
    return {"status": "ok"}
