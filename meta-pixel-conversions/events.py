from facebook_business.adobjects.serverside.content import Content
from facebook_business.adobjects.serverside.custom_data import CustomData

from utils import send_event, generate_event_id_and_time

def product_detail_event(pixel_id, user_data, product_name, product_id, category, value, event_source_url=None, test_event_code=None):
    """Track ViewContent event for a product page."""
    custom_data = CustomData(
        content_type='product',
        content_name=product_name,
        content_ids=[product_id],
        content_category=category,
        currency='TRY',
        value=value
    )

    # for redundant setup it's important
    event_id, event_time = generate_event_id_and_time()

    return send_event(event_id, event_time, 'ViewContent', pixel_id, user_data, custom_data, event_source_url, test_event_code)

def shop_detail_event(pixel_id, user_data, shop_name, shop_id, event_source_url="", test_event_code=None):
    """Track ViewContent event for a shop page."""
    custom_data = CustomData(
        content_type='shop',
        content_name=shop_name,
        content_ids=[shop_id]
    )

    # for redundant setup it's important
    event_id, event_time = generate_event_id_and_time()

    return send_event(event_id, event_time, 'ViewContent', pixel_id, user_data, custom_data, event_source_url, test_event_code)

def add_to_cart_event(pixel_id, user_data, product_name, product_id, category, value, quantity, event_source_url=None, test_event_code=None):
    """Track AddToCart event."""
    content = Content(
        product_id=product_id,
        quantity=quantity
    )

    custom_data = CustomData(
        content_type='product',
        content_name=product_name,
        content_ids=[product_id],
        content_category=category,
        contents=[content],
        currency='TRY',
        value=value
    )

    # for redundant setup it's important
    event_id, event_time = generate_event_id_and_time()

    return send_event(event_id, event_time, 'AddToCart', pixel_id, user_data, custom_data, event_source_url, test_event_code)

def initiate_checkout_event(pixel_id, user_data, product_ids, value, num_items, event_source_url=None, test_event_code=None):
    """Track InitiateCheckout event."""
    custom_data = CustomData(
        content_type='product',
        content_ids=product_ids,
        currency='TRY',
        value=value,
        num_items=num_items
    )

    # for redundant setup it's important
    event_id, event_time = generate_event_id_and_time()

    return send_event(event_id, event_time, 'InitiateCheckout', pixel_id, user_data, custom_data, event_source_url, test_event_code)

def purchase_event(pixel_id, user_data, product_ids, value, num_items, currency='TRY', event_source_url=None, test_event_code=None):
    """Track Purchase event."""
    custom_data = CustomData(
        content_type='product',
        content_ids=product_ids,
        currency=currency,
        value=value,
        num_items=num_items
    )

    # for redundant setup it's important
    event_id, event_time = generate_event_id_and_time()

    return send_event(event_id, event_time, 'Purchase', pixel_id, user_data, custom_data, event_source_url, test_event_code)
