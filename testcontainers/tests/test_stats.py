# A small helper to generate multiple URLs and clicks
TEST_DATA = [
    ("http://good-stats.com", 1),
    ("http://safe-stats.com", 3),
    ("http://trusted-stats.com", 2),
    ("https://legitimate-stats.com", 4),
]


def test_stats_flow_multiple(client):
    # 1) Create and click each URL the appropriate number of times
    for url, clicks in TEST_DATA:
        # create short URL
        r = client.post("/create-short-url", json={"url": url})
        assert r.status_code == 201
        code = r.json()["short_code"]

        # issue the redirect 'clicks' times
        for _ in range(clicks):
            r2 = client.get(f"/{code}", follow_redirects=False)
            assert r2.status_code == 302

    # 2) Generate a stats dump
    r = client.post("/stat/generate")
    assert r.status_code == 201

    # 3) List dumps → pick the newest
    r2 = client.get("/stat/list")
    assert r2.status_code == 200
    dumps = r2.json()
    assert isinstance(dumps, list) and dumps, "Expected at least one dump"
    newest = dumps[0]
    dump_id = newest["id"]
    assert isinstance(dump_id, int)

    # 4) Download the JSON payload
    r3 = client.get(f"/stat/download/{dump_id}")
    assert r3.status_code == 200
    assert r3.headers["content-type"].startswith("application/json")

    payload = r3.json()
    # payload is a list of dicts
    assert isinstance(payload, list)

    # 5) Verify each URL has the correct click count
    for url, expected_count in TEST_DATA:
        entry = next((e for e in payload if e["original_url"] == url), None)
        assert entry is not None, f"{url} missing from stats dump"
        assert entry["total_clicks"] == expected_count
