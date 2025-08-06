from contextlib import asynccontextmanager
from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from sqlalchemy import text

from app.routers.shortener import router as shortener_router
from app.dependencies import get_db, engine
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

@app.get("/health")
def health(db: Session = Depends(get_db)):
    """
    Basic health check, now cleaner and using the shared `get_db` dependency.
    """
    try:
        # In SQLAlchemy 2.0+, it's recommended to wrap raw SQL in text()
        db.execute(text("SELECT 1"))
        print("I'm coming here no problem!")
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Database error: {e}")

    return {"status": "ok 2"}
