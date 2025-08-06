from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session

from app.schemas import CreateShortURLRequest
from app.dependencies import get_db
from app.crud import create_url
from app.services.phishing import is_phishing_url

router = APIRouter(tags=["shortener"])


@router.post(
    "/create-short-url",
    status_code=201,
)
def create_short(
    payload: CreateShortURLRequest,
    db: Session = Depends(get_db),
):
    phishing_flag = is_phishing_url(payload.url)
    if phishing_flag:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST, detail="Phishing URL detected."
        )

    url_obj = create_url(db, payload.url, phishing_flag)

    response = {
        "short_code": url_obj.short_code,
        "is_phishing": url_obj.is_phishing,
    }

    return response
