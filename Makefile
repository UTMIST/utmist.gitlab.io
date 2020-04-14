build:
	go run main.go -depts -events

clean:
	rm -rf content/events content/team content/project content/*.md config.yaml

theme:
	git submodule update --init --recursive