build:
	go build -o bin/gdcache

run: build
	./bin/gdcache