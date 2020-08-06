.PHONY: base build clean dep fetch full

build: dust
	cp -r content_base content
	cp -r templates_base templates
	go run main.go

clean:
	rm -rf content content_base config.yaml public templates

dep:
	git submodule update --init --recursive

dust: 
	rm -rf content templates

fetch:
	sh onedrive.sh

full: clean fetch base build

