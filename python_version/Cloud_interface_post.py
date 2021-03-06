import os
import time
import requests  # pip install requests==2.25.1
import random

# Environment variables for authentication credentials.
SERVICE_ACCOUNT_KEY_ID  = os.environ.get("DT_SERVICE_ACCOUNT_KEY_ID")
SERVICE_ACCOUNT_SECRET  = os.environ.get("DT_SERVICE_ACCOUNT_SECRET")
SENSOR_ID               = os.environ.get("DT_SENSOR_ID")
PROJECT_ID              = os.environ.get("DT_PROJECT_ID")

URL = f"https://emulator.d21s.com/v2/projects/{PROJECT_ID}/devices/{SENSOR_ID}:publish"

def post_datapoint(datapoint):
    post = requests.post(
        url=URL,
        auth=(SERVICE_ACCOUNT_KEY_ID, SERVICE_ACCOUNT_SECRET),
        data=str( {"temperature": {"value": str(datapoint)}} ).replace("\'", "\"")
    )
    print(post)



for i in range(6):
    before = time.time()
    post_datapoint(random.randrange(-10**4, 10**4)/10**(2))
    print(time.time() - before)
    time.sleep(1)


