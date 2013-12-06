all: timelapse ui_timelapse.py

timelapse: main.go
	go build

ui_timelapse.py: timelapse.ui
	pyuic4 -x -o ui_timelapse.py timelapse.ui

clean:
	rm -f timelapse ui_timelapse.py *.pyc

format:
	gofmt -w -tabs=false -tabwidth=4 *.go
