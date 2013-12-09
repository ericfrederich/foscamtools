all: timelapse ui_timelapse.py

timelapse: timelapse.go
	go build timelapse.go

ui_timelapse.py: timelapse.ui
	pyuic4 -x -o ui_timelapse.py timelapse.ui

clean:
	rm -f timelapse ui_timelapse.py *.pyc

format:
	gofmt -w -tabs=false -tabwidth=4 *.go
