Cloud Powermeter

A small project using a raspberry pi to track the blinks from the LED on the power meter in the house fuse box, to track power consumption. A LDR and comparator circuit is used to make a 3.3V signal when the LED is on and 0V when it's off. I started making it in python, but decided to make it in go to play with its multithreading. The data is posted to a cloud service. There is a bash script to push changes from a windows pc to the raspberry pi for easy testing and updating.
