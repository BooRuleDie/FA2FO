from sitemap import gen_products_sitemap, gen_shops_sitemap, gen_categories_sitemap, gen_index_sitemap

def main():
    gen_products_sitemap(ignore_priority=True)
    gen_shops_sitemap()
    gen_categories_sitemap()
    gen_index_sitemap()

if __name__ == "__main__":
    main()