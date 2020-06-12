build:
	go run main.go -depts -events -projects

clean:
	rm -rf content/events content/team content/*.md config.yaml public

theme:
	git submodule update --init --recursive