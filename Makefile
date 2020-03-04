build:
	go run main.go
	hugo server -D

theme:
	git submodule update --init --recursive

clean:
	rm -rf content/events content/team
	rm config.yaml
