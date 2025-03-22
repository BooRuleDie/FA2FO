# What's Conversions API

In order to track user actions like adding a product to cart, going to checkout page, or purchasing a product, Meta Pixel is used on the client side. However, modern browsers have some strict policies that might block 3rd party requests that Meta Pixel would make.

To solve this problem and offer a more robust solution, **Meta Conversions API** was introduced. Instead of tracking user actions on the client side, these requests are sent from the server side. This means the backend handles this work, and to communicate with Meta, these Conversions APIs are used.

For more detailed information about implementation and best practices, refer to the [official Meta Conversions API documentation](https://developers.facebook.com/docs/marketing-api/conversions-api/).

# What's Redundant Setup?

Redundant setup refers to implementing both Meta Pixel and Conversions API simultaneously to track the same events. This dual implementation provides more reliable tracking since if one method fails (e.g. Meta Pixel is blocked by browser), the other method (Conversions API) can still capture the event data. It's considered a best practice to use both systems together to ensure maximum data accuracy and reliability in your Meta event tracking.

# Meta Conversions API Key Parameters

When implementing Meta Conversions API, certain parameters are more critical than others for effective tracking and attribution. Here's a breakdown of the most important parameters:

## Most Critical (Must Match)
- **event_id** - The unique identifier for each event (critical for deduplication)
- **event_name** - The standardized action name (Purchase, AddToCart, etc.)
- **event_time** - Timestamp of when the event occurred

## Very Important
User identification parameters (in order of importance):
- **fbp** (Facebook Browser Pixel cookie) - Critical for web events
- **email** (hashed) - Primary user identifier
- **external_id** (your customer ID) - If you use your own system IDs
- **fbc** (Facebook Click ID cookie) - Important for attribution
- **phone** (hashed) - Additional identifier

## Important
Event-specific parameters:
- For purchase events: **value** and **currency**
- **content_ids** for product-related events
- **content_type** (product, product_group, etc.)

# Meta Pixel Event Call Examples

Before any fbq event call, fbq init must be called first. Best practice is calling it in the head section of each page to prepare for single or multiple event calls by the fbq function.
```js
// Initialize Meta Pixel with your pixel ID
fbq('init', '123456789012345');
```

```js
// Product Detail Page
fbq('track', 'ViewContent', {
    content_type: 'product',
    currency: 'TRY',

    // dynamic fields
    content_name: 'Product Name',
    content_ids: ['ABC123'],
    content_category: 'Category Name',
    value: 2999.99,
});

// Shop Detail Page
fbq('track', 'ViewContent', {
    content_type: 'shop',

    // dynamic fields
    content_name: 'Shop Name',
    content_ids: ['SHOP123'],
});

// Add to Cart
fbq('track', 'AddToCart', {
    content_type: 'product',
    currency: 'TRY'

    // dynamic fields
    content_name: 'Product Name',
    content_ids: ['ABC123'],
    content_category: 'Category Name',
    value: 2999.99,
    quantity: 1,
});

// Checkout Page
fbq('track', 'InitiateCheckout', {
    content_type: 'product',
    currency: 'TRY'

    // dynamic fields
    content_ids: ['ABC123', 'DEF456'],
    value: 4599.98,
    num_items: 2,
});

```
