bare:
	go run main.go

depts:
	go run main.go -depts

events:
	go run main.go -events

all:
	go run main.go -depts -events

clean:
	rm -rf content/events content/team content/*.md config.yaml public

fresh:
	git submodule update --init --recursive
