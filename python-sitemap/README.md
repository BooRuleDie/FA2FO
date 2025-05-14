# Sitemap XML Guidelines

## Key Points

- **Sitemap Index**: For sites with multiple sitemaps, use a sitemap index file (`sitemap-index.xml`) to list all individual sitemap files
- **Size Limits**: 
  - Maximum 50,000 URLs per sitemap file
  - Maximum file size of 50MB (uncompressed)
  - Maximum 50,000 sitemaps in a sitemap index
- **File Format**: XML format with UTF-8 encoding

## Best Practices

- **Compression**: Use GZIP compression for large sitemap files (.xml.gz)
- **URLs**:
  - Must be properly encoded
  - Should use canonical URLs
  - Include protocol (http/https)
- **Update Frequency**: Update sitemaps when content changes
- **Submission**: Submit sitemap URL to search engines via:
  - robots.txt file
  - Direct submission in search console
  - Direct HTTP ping

## Required Elements

- `<loc>`: Full URL of the page (required)
- `<lastmod>`: Date of last modification (recommended)
- `<changefreq>`: How often page changes (optional)
- `<priority>`: Relative importance 0.0 to 1.0 (optional)

## Example Structure

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://example.com/page1</loc>
    <lastmod>2023-01-01</lastmod>
    <changefreq>monthly</changefreq>
    <priority>0.8</priority>
  </url>
</urlset>
```

## Example Index Sitemap Structure

```xml
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <sitemap>
      <loc>http://www.example.com/sitemap1.xml.gz</loc>
      <lastmod>2004-10-01T18:23:17+00:00</lastmod>
   </sitemap>
   <sitemap>
      <loc>http://www.example.com/sitemap2.xml.gz</loc>
      <lastmod>2005-01-01</lastmod>
   </sitemap>
</sitemapindex>
```


## Important Notes

- Don't include URLs that return non-200 status codes
- Include only canonical versions of URLs
- Keep sitemaps up to date to ensure proper indexing
- Monitor sitemap errors in search console
- Consider using dynamic sitemap generation for large sites