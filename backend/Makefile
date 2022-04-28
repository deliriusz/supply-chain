test:
	go test ./...

clean:
	rm -rf bin
	rm -rf build
	rm -rf artifacts

run:
	go run .

build:
	go run -o bin/main main.go

eth:
	mkdir -p build
	mkdir -p artifacts
	scripts/generate-contract-artifacts.sh