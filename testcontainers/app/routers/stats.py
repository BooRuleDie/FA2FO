import uuid

from fastapi import APIRouter, Depends, HTTPException, Response, status
from fastapi.responses import StreamingResponse
from sqlalchemy.orm import Session

from app.dependencies import get_db, get_s3
from app.crud import create_stat_dump, list_stat_dumps, get_stat_dump_by_id
from app.services.stats import build_stats_payload, upload_to_s3, get_s3_object

router = APIRouter(tags=["stats"])

# S3 bucket name
BUCKET = "shortener-stats"

@router.post("/stat/generate", status_code=201)
def generate_stats(db: Session = Depends(get_db), s3=Depends(get_s3)):
    # build JSON payload
    data = build_stats_payload(db)

    # derive an S3 key (use a UUID or timestamp)
    dump_key = f"stats-{uuid.uuid4()}.json"

    # upload bytes to S3
    upload_to_s3(s3, bucket=BUCKET, key=dump_key, data=data)

    # record in Postgres
    create_stat_dump(db, key=dump_key)

    # return empty body
    return Response(status_code=status.HTTP_201_CREATED)


@router.get("/stat/list")
def list_stats(db: Session = Depends(get_db)):
    return list_stat_dumps(db)


@router.get("/stat/download/{dump_id}")
def download_stats(
    dump_id: int,
    db: Session = Depends(get_db),
    s3=Depends(get_s3),
):
    # fetch the StatDump row
    dump = get_stat_dump_by_id(db, dump_id)
    if not dump:
        raise HTTPException(status.HTTP_404_NOT_FOUND, detail="Stat dump not found")

    # fetch the object from S3 using the stored key
    obj = get_s3_object(s3, bucket=BUCKET, key=str(dump.key))
    if obj is None:
        raise HTTPException(404, detail="S3 object not found")

    # stream back as a JSON file
    return StreamingResponse(
        obj["Body"],
        media_type="application/json",
        headers={"Content-Disposition": f'attachment; filename="{dump.key}"'},
    )
