build:
	go run main.go
	hugo server -D

theme:
	git submodule update --init --recursive

clean:
	rm -rf content/event content/team content/project content/*.md
	rm config.yaml

