.PHONY: build clean dep fetch full

build: dust
	cp -r content_base content
	cp -r insertions_base insertions
	go run main.go

clean:
	rm -rf content content_base config.yaml public insertions insertions_base static/*.pdf static/images/profilepics

dep:
	git submodule update --init --recursive

dust: 
	rm -rf content insertions

fetch:
	sh onedrive.sh

full: clean fetch build

