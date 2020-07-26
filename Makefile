build:
	sh onedrive.sh
	go run main.go

clean:
	rm -rf content config.yaml public templates

dep:
	git submodule update --init --recursive
