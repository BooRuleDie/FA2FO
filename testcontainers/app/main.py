from fastapi import FastAPI

app = FastAPI(title="URL Shortener Demo (sync)")

@app.get("/health")
def health_check():
    return {"status": "ok"}
