from fastapi import APIRouter, Depends, HTTPException, status, Response
from sqlalchemy.orm import Session
from redis import Redis

from app.schemas import CreateShortURLRequest
from app.dependencies import get_db, get_redis
from app.crud import create_url, get_url_by_code, create_click
from app.services.phishing import is_phishing_url

import typing

router = APIRouter(tags=["shortener"])


@router.post(
    "/create-short-url",
    status_code=201,
)
def create_short(
    payload: CreateShortURLRequest,
    db: Session = Depends(get_db),
    redis: Redis = Depends(get_redis),
):
    phishing_flag = is_phishing_url(redis, payload.url)
    
    url_obj = create_url(db, payload.url, phishing_flag)

    response = {
        "short_code": url_obj.short_code,
        "is_phishing": url_obj.is_phishing,
    }

    return response

@router.get("/{short_code}")
def redirect_short(
    short_code: str,
    db: Session = Depends(get_db)
):
    # we still want to verify the phishing flag in DB
    url = get_url_by_code(db, short_code)
    if not url:
        raise HTTPException(status.HTTP_404_NOT_FOUND, "Short code not found")
    elif url.is_phishing is True:
        raise HTTPException(status.HTTP_403_FORBIDDEN, "URL flagged as phishing")

    # record click and redirect
    create_click(db, typing.cast(int, url.id))

    return Response(status_code=302, headers={"Location": str(url.original_url)})
