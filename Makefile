all:
	go run main.go -depts -events

clean:
	rm -rf content/events content/team content/*.md config.yaml public

depts:
	go run main.go -depts

events:
	go run main.go -events

fresh:
	git submodule update --init --recursive
