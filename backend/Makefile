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
	solc @openzeppelin=../frontend/node_modules/@openzeppelin \
		--abi ../frontend/src/contracts/Authorization.sol > build/Authorization.abi
	solc @openzeppelin=../frontend/node_modules/@openzeppelin \
		--bin ../frontend/src/contracts/Authorization.sol > build/Authorization.bin
	abigen --bin=build/Authorization.bin --abi=build/Authorization.abi --pkg authorization --out=artifacts/Authorization.go