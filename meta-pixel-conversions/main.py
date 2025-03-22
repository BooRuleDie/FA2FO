from pprint import pprint
import os

from events import (
    product_detail_event,
    shop_detail_event,
    add_to_cart_event,
    initiate_checkout_event,
    purchase_event
)
from utils import create_user_data

from facebook_business.api import FacebookAdsApi

# conversions fake data
TEST_ACCESS_TOKEN = os.environ.get('TEST_ACCESS_TOKEN')
PROD_ACCESS_TOKEN = os.environ.get('PROD_ACCESS_TOKEN')
PIXEL_ID = os.environ.get('PIXEL_ID')
TEST_CODE = os.environ.get('TEST_CODE')

# user fake data
CUSTOMER_EMAIL = os.environ.get('CUSTOMER_EMAIL') if os.environ.get('CUSTOMER_EMAIL') else None
CUSTOMER_PHONE = os.environ.get('CUSTOMER_PHONE') if os.environ.get('CUSTOMER_PHONE') else None
TEST_TR_IP = '78.168.42.123'
TEST_USER_AGENT = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36'
FBC = 'fb.1.1612345678910.1234567890'
FBP = 'fb.2.1612345678910.1234567890'

# set the access token in use
ACCESS_TOKEN = TEST_ACCESS_TOKEN    

def main():
    FacebookAdsApi.init(access_token=ACCESS_TOKEN)

    user_data = create_user_data(
        email=CUSTOMER_EMAIL,
        phone=CUSTOMER_PHONE,
        client_ip=TEST_TR_IP,
        user_agent=TEST_USER_AGENT,
        fbc=FBC,
        fbp=FBP
    )

    # Product Detail Event
    res = product_detail_event(
        pixel_id=PIXEL_ID,
        user_data=user_data,
        product_name='Sample Product',
        product_id='ABC123',
        category='Sample Category',
        value=2999.99,
        event_source_url='https://example.com/product/ABC123',
        test_event_code=TEST_CODE
    )
    pprint(res)
    
    # Shop Detail Event
    res = shop_detail_event(
        pixel_id=PIXEL_ID,
        user_data=user_data,
        shop_name='Sample Shop',
        shop_id='ABC123',
        event_source_url='https://example.com/product/ABC123',
        test_event_code=TEST_CODE
    )
    pprint(res)
    
    # Add to Cart Event
    res = add_to_cart_event(
        pixel_id=PIXEL_ID,
        user_data=user_data,
        product_name='Sample Product',
        product_id='ABC123',
        category='Sample Category',
        value=2999.99,
        quantity=1,
        event_source_url='https://example.com/product/ABC123',
        test_event_code=TEST_CODE
    )
    pprint(res)

    # Initiate Checkout Event
    res = initiate_checkout_event(
        pixel_id=PIXEL_ID,
        user_data=user_data,
        product_ids=['ABC123', 'DEF456'],
        value=5999.98,
        num_items=2,
        event_source_url='https://example.com/checkout',
        test_event_code=TEST_CODE
    )
    pprint(res)
    
    # Purchase Event
    res = purchase_event(
        pixel_id=PIXEL_ID,
        user_data=user_data,
        product_ids=['ABC123', 'DEF456'],
        value=5999.98,
        num_items=2,
        currency='TRY',
        event_source_url='https://example.com/purchase',
        test_event_code=TEST_CODE
    )
    pprint(res)


if __name__ == "__main__":
    main()
