from sqlalchemy.orm import Session
from app.models import URL

def create_url(
    db: Session,
    original_url: str,
    is_phishing: bool
) -> URL:
    u = URL(
        original_url=original_url,
        is_phishing=is_phishing
    )
    db.add(u)
    db.commit()
    db.refresh(u)
    return u

def get_url_by_code(db: Session, short_code: str) -> URL | None:
    return db.query(URL).filter(URL.short_code == short_code).one_or_none()
