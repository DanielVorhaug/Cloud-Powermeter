import gpiozero as gp
import time 
import Cloud_interface_post.py

BLINKS_PER_KWH = 2000
MESSAGE_PERIOD = 5 # [seconds]

ldr = gp.DigitalInputDevice(15) # Number is sensor-pin. Signal should be 3.3V when light is on and 0V when light is off.

blink_count                 = 0
last_blink_time             = time.time()
sampling_period_start_time  = time.time()
next_message_time           = time.time() + 5

while True:
    if ldr.value:
        blink_count += 1
        last_blink_time = time.time()
    
    if time.time() > next_message_time:
        post_datapoint(sampling_period_start_time-last_blink_time)

        
        
#         while ldr.value:
#             sleep(0.01)