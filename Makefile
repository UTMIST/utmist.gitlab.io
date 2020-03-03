test:
	go run main.go
	hugo server -D

clean:
	rm -rf content/events content/team content/events.md
	rm config.yaml

