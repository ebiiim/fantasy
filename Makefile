all: ci build-wasm

ci:
	git fetch --unshallow || true
	git fetch --tags

serve:
	python -m http.server --bind 127.0.0.1 --directory ./dist/wasm 8080

build-wasm:	./dist/wasm/wasm_exec.js
	go generate ./... 
	GOOS=js GOARCH=wasm go build \
	  -o ./dist/wasm/a.wasm \
	  -ldflags "-X \"main.version=$$(git describe --tags)\" -X \"main.buildDate=$$(date --iso-8601=seconds)\" -X \"main.goVersion=$$(go version)\"" ./main.go
	cp ./web/* ./dist/wasm

build-native:
	go generate ./... 
	go build -ldflags "-X \"main.version=$$(git describe --tags)\" -X \"main.buildDate=$$(date --iso-8601=seconds)\" -X \"main.goVersion=$$(go version)\"" ./main.go

./dist/wasm/wasm_exec.js:
	mkdir -p ./dist/wasm
	curl -L -o ./dist/wasm/wasm_exec.js https://raw.githubusercontent.com/golang/go/go1.17.5/misc/wasm/wasm_exec.js
