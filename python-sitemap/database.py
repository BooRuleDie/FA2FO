import mysql.connector
from utils import read_file

import os


DB_DATABASE = os.getenv("DB_DATABASE")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_USER = os.getenv("DB_USER")
DB_HOST = os.getenv("DB_HOST")


CATEGORY_SITEMAP_SQL = read_file("SQLs/category.sql")
PRODUCT_SITEMAP_SQL = read_file("SQLs/product.sql")
SHOP_SITEMAP_SQL = read_file("SQLs/shop.sql")


def get_connection() -> tuple:
    try:
        conn = mysql.connector.connect(
            host=DB_HOST,
            user=DB_USER,
            password=DB_PASSWORD,
            database=DB_DATABASE,
        )
        cursor = conn.cursor(dictionary=True)
        return cursor, conn
    except Exception as error:
        print(f"get_connection error: {error}")
        return None, None


def get_products():
    cursor, conn = None, None
    try:
        cursor, conn = get_connection()
        cursor.execute(PRODUCT_SITEMAP_SQL)
        result = cursor.fetchall()
        return result
    except Exception as error:
        print(f"get_products error: {error}")
    finally:
        if cursor:
            cursor.close()
        if conn:
            conn.close()


def get_shops():
    cursor, conn = None, None
    try:
        cursor, conn = get_connection()
        cursor.execute(SHOP_SITEMAP_SQL)
        result = cursor.fetchall()
        return result
    except Exception as error:
        print(f"get_shops error: {error}")
    finally:
        if cursor:
            cursor.close()
        if conn:
            conn.close()


def get_categories():
    cursor, conn = None, None
    try:
        cursor, conn = get_connection()
        cursor.execute(CATEGORY_SITEMAP_SQL)
        result = cursor.fetchall()
        return result
    except Exception as error:
        print(f"get_categories error: {error}")
    finally:
        if cursor:
            cursor.close()
        if conn:
            conn.close()
