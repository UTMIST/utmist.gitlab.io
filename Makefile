build:
	go run main.go

full:
	make clean
	sh onedrive.sh
	go run main.go

dep:
	git submodule update --init --recursive

clean:
	rm -rf content config.yaml public templates
