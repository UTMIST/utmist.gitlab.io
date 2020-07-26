build:
	sh onedrive.sh
	go run main.go -depts

clean:
	rm -rf content config.yaml public

dep:
	git submodule update --init --recursive
