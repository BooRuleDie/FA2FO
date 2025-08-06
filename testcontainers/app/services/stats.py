from io import BytesIO
import json

from sqlalchemy.orm import Session
from app.models import URL, Click
from sqlalchemy import func

def build_stats_payload(db: Session) -> bytes:
    # Query URLs with total click counts
    results = (
        db.query(
            URL.short_code,
            URL.original_url,
            URL.is_phishing,
            func.count(Click.id).label("total_clicks"),
        )
        .outerjoin(Click, Click.url_id == URL.id)
        .group_by(URL.id)
        .all()
    )
    payload = [
        {
            "short_code": str(short_code),
            "original_url": original_url,
            "is_phishing": is_phishing,
            "total_clicks": total_clicks or 0,
        }
        for short_code, original_url, is_phishing, total_clicks in results
    ]
    return json.dumps(payload, indent=4).encode("utf-8")

def upload_to_s3(s3_client, bucket: str, key: str, data: bytes):
    try:
        s3_client.create_bucket(Bucket=bucket)
    except s3_client.exceptions.BucketAlreadyOwnedByYou:
        pass
    except s3_client.exceptions.BucketAlreadyExists:
        pass

    s3_client.put_object(Bucket=bucket, Key=key, Body=BytesIO(data))
    
    
def get_s3_object(s3_client, bucket: str, key: str):
    try:
        obj = s3_client.get_object(Bucket=bucket, Key=key)
    except s3_client.exceptions.NoSuchKey:
        return None
    return obj
