from redis import Redis
import requests

from app.config import settings


def is_phishing_url(redis: Redis, url: str) -> bool:
    key = f"is_phishing:{url}"

    # Try to get from cache
    cached = redis.get(key)
    if cached is not None:
        is_phishing = cached == "1"
        print(f"Cache for URL: {url}, is_phishing: {is_phishing}, cache: {cached}")
        return is_phishing

    # Cache miss, fetch from 3rd party service
    print(f"3rd party API for URL: {url}.")

    resp = requests.get(
        f"{settings.PHISHING_API_URL}", params={"url": url}, timeout=2
    )
    resp.raise_for_status()
    is_phishing = resp.json().get("is_phishing", False)

    # Cache the result: "1" for phishing, "0" for not phishing
    cache_value = "1" if is_phishing else "0"
    redis.set(key, cache_value, ex=3600)

    return is_phishing
