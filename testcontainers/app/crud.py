from sqlalchemy.orm import Session
from app.models import URL, Click, StatDump
import uuid

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
    try:
        code_uuid = uuid.UUID(short_code)
    except (ValueError, TypeError):
        return None

    return db.query(URL).filter(URL.short_code == code_uuid).one_or_none()


def create_click(db: Session, url_id: int) -> Click:
    click = Click(url_id=url_id)
    db.add(click)
    db.commit()
    db.refresh(click)
    return click


def create_stat_dump(db: Session, key: str) -> StatDump:
    dump = StatDump(key=key)
    db.add(dump)
    db.commit()
    db.refresh(dump)
    return dump


def get_stat_dump_by_id(db: Session, dump_id: int) -> StatDump | None:
    return db.query(StatDump).filter(StatDump.id == dump_id).one_or_none()


def list_stat_dumps(db: Session) -> list[StatDump]:
    return db.query(StatDump).all()
