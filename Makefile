test:
	go run main.go
	hugo server -D

clean:
	rm -rf content/events
	rm config.yaml

