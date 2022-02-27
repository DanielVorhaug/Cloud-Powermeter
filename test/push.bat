scp test/test.go pi@powerpi:/home/pi/CODE/test/test.go 
ssh pi@powerpi.local "cd /home/pi/CODE/test/ && go run test.go"
