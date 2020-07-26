build:
	sh onedrive.sh
	go run main.go -depts

clean:
	rm -rf content config.yaml public

fresh:
	git submodule update --init --recursive
