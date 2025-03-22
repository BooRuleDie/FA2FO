import time
import random
import string

from facebook_business.adobjects.serverside.event import Event
from facebook_business.adobjects.serverside.user_data import UserData
from facebook_business.adobjects.serverside.action_source import ActionSource
from facebook_business.adobjects.serverside.event_request import EventRequest

def create_user_data(email=None, phone=None, client_ip=None, user_agent=None, fbc=None, fbp=None):
    """Create and return a UserData object with the provided parameters."""
    return UserData(
        email=email,
        phone=phone,
        client_ip_address=client_ip,
        client_user_agent=user_agent,
        fbc=fbc,
        fbp=fbp,
    )

def send_event(event_id, event_time, event_name, pixel_id, user_data, custom_data, event_source_url=None, test_event_code=None):
    """Create and send an event with the provided parameters."""
    event = Event(
        event_id=event_id,
        event_name=event_name,
        event_time=event_time,
        user_data=user_data,
        custom_data=custom_data,
        event_source_url=event_source_url,
        action_source=ActionSource.WEBSITE,
    )

    event_request = EventRequest(
        events=[event],
        pixel_id=pixel_id,
        test_event_code=test_event_code
    )

    return event_request.execute()

# very crucial function for event setup
def generate_event_id_and_time():
    # Get current timestamp in milliseconds
    timestamp = int(time.time() * 1000)  # timestamp in milliseconds
    event_time = int(time.time())  # event_time in seconds

    # Generate random string (equivalent to Math.random().toString(36).substring(2, 15))
    random_chars = ''.join(random.choices(string.ascii_lowercase + string.digits, k=13))

    # Format: fb_[timestamp]_[random-string]
    return f"fb_{timestamp}_{random_chars}", event_time
