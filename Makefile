build: config.yaml
	go run main.go

clean:
	rm -rf content/events content/team content/project content/*.md
	rm config.yaml

fresh: build run

run: 
	hugo server -D

theme:
	git submodule update --init --recursive