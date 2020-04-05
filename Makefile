build:
	go run main.go

buildrun: build run

clean:
	rm -rf content/events content/team content/project content/*.md
	rm config.yaml

run: 
	hugo server -D

theme:
	git submodule update --init --recursive