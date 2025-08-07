import string
import random

GOOD_URL = "https://good.com"
BAD_URL = "https://bad.com"


def random_url():
    suffix = "".join(
        random.choices(string.ascii_lowercase + string.digits, k=8)
    )
    return f"https://{suffix}.com/"


def test_non_phishing_url(client):
    r = client.post("/create-short-url", json={"url": GOOD_URL})
    assert r.status_code == 201
    data = r.json()
    assert data["is_phishing"] is False
    assert isinstance(data["short_code"], str)


def test_phishing_url(client):
    r = client.post("/create-short-url", json={"url": BAD_URL})
    assert r.status_code == 201
    data = r.json()
    print("problematic value", data)
    assert data["is_phishing"]
    # subsequent redirect must be forbidden
    follow = client.get(f"/{data['short_code']}", follow_redirects=False)
    assert follow.status_code == 403


def test_random_phishing_distribution(client):
    """
    Send 20 random URLs, expect at least one phishing=True.
    Then pick that URL and issue 5 more create calls against it,
    verifying is_phishing stays True.
    """
    seen_flagged = None
    for _ in range(20):
        url = random_url()
        r = client.post("/create-short-url", json={"url": url})
        assert r.status_code == 201
        js = r.json()
        if js["is_phishing"]:
            seen_flagged = url
            break
    
    assert seen_flagged, "No phishing responses in 20 requests → test failed"

    # repeat 5 times on the same original URL
    for i in range(5):
        r = client.post("/create-short-url", json={"url": seen_flagged})
        assert r.status_code == 201
        assert r.json()[
            "is_phishing"
        ], f"Expected phishing=True on retry #{i+1}"


def test_redirect_logic(client):
    # create a safe URL
    r = client.post("/create-short-url", json={"url": GOOD_URL})
    code = r.json()["short_code"]

    # immediate redirect
    r2 = client.get(f"/{code}", follow_redirects=False)
    assert r2.status_code == 302
    assert r2.headers["location"] == GOOD_URL

    # clicking again still works
    r3 = client.get(f"/{code}", follow_redirects=False)
    assert r3.status_code == 302
    assert r3.headers["location"] == GOOD_URL
