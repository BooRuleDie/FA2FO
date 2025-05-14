from database import get_products, get_shops, get_categories
from utils import generate_slug

import datetime
import os

BASE_URL = os.getenv("BASE_URL")
PRODUCT_ENDPOINT = os.getenv("PRODUCT_ENDPOINT")
SHOP_ENDPOINT = os.getenv("SHOP_ENDPOINT")
CATEGORY_ENDPOINT = os.getenv("CATEGORY_ENDPOINT")

NOW = datetime.date.today().strftime("%Y-%m-%d")
SITEMAP_EXPORT_PATH = "./sitemaps"


def gen_products_sitemap(ignore_priority=False):
    print("[!] fetching products...")
    products = get_products()
    if not products:
        return
    print("[!] fetching products done")
    print("[!] creating products sitemap...")

    with open(f"{SITEMAP_EXPORT_PATH}/products-sitemap.xml", "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n')

        for p in products:
            product_id = p["product_id"]
            product_name = p["product_name"]
            priority = p["priority"]

            slug = generate_slug(product_id, product_name)

            f.write("  <url>\n")
            f.write(f"    <loc>{BASE_URL}{PRODUCT_ENDPOINT}{slug}</loc>\n")
            f.write(f"    <lastmod>{NOW}</lastmod>\n")
            if not ignore_priority:
                f.write(f"    <priority>{priority}</priority>\n")
            f.write("  </url>\n")

        f.write("</urlset>")

    print("[!] creating products done")


def gen_shops_sitemap():
    print("[!] fetching shops...")
    shops = get_shops()
    if not shops:
        return
    print("[!] fetching shops done")
    print("[!] creating shops sitemap...")

    with open(f"{SITEMAP_EXPORT_PATH}/shops-sitemap.xml", "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n')

        for s in shops:
            shop_id = s["shop_id"]
            shop_name = s["shop_name"]

            slug = generate_slug(shop_id, shop_name)

            f.write("  <url>\n")
            f.write(f"    <loc>{BASE_URL}{SHOP_ENDPOINT}{slug}</loc>\n")
            f.write(f"    <lastmod>{NOW}</lastmod>\n")
            f.write("  </url>\n")

        f.write("</urlset>")

    print("[!] creating shops done")


def gen_categories_sitemap():
    print("[!] fetching categories...")
    categories = get_categories()
    if not categories:
        return
    print("[!] fetching categories done")
    print("[!] creating categories sitemap...")

    with open(f"{SITEMAP_EXPORT_PATH}/categories-sitemap.xml", "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n')

        for c in categories:
            category_id = c["category_id"]
            category_name = c["category_name"]

            slug = generate_slug(category_id, category_name)

            f.write("  <url>\n")
            f.write(f"    <loc>{BASE_URL}{CATEGORY_ENDPOINT}{slug}</loc>\n")
            f.write(f"    <lastmod>{NOW}</lastmod>\n")
            f.write("  </url>\n")

        f.write("</urlset>")

    print("[!] creating categories done")


def gen_index_sitemap():
    print("[!] creating index sitemap...")
    endpoints = [PRODUCT_ENDPOINT, SHOP_ENDPOINT, "category/"]

    with open(f"{SITEMAP_EXPORT_PATH}/sitemap-index.xml", "w") as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n')

        for endpoint in endpoints:
            f.write("  <sitemap>\n")
            f.write(f"    <loc>{BASE_URL}{endpoint}sitemap.xml</loc>\n")
            f.write(f"    <lastmod>{NOW}</lastmod>\n")
            f.write("  </sitemap>\n")

        f.write("</sitemapindex>")

    print("[!] creating index done")
