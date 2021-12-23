all: ci build-wasm

# some CI environments use shallow clone but we need normal clone for injecting version info
ci:
	git fetch --unshallow || true
	git fetch --tags

# go run always exit with 1 so added `|| true`
serve:
	go run tool/easy_server/main.go -dir=dist/wasm || true

build-wasm:	./dist/wasm/wasm_exec.js
	go generate ./... 
	GOOS=js GOARCH=wasm go build \
	  -o ./dist/wasm/a.wasm \
	  -ldflags "-X \"main.version=$$(git describe --tags)\" -X \"main.buildDate=$$(LC_TIME=C TZ=Asia/Tokyo date)\" -X \"main.goVersion=$$(go version)\"" ./main.go
	cp ./web/* ./dist/wasm

build-native:
	go generate ./... 
	go build -ldflags "-X \"main.version=$$(git describe --tags)\" -X \"main.buildDate=$$(LC_TIME=C TZ=Asia/Tokyo date)\" -X \"main.goVersion=$$(go version)\"" ./main.go

./dist/wasm/wasm_exec.js:
	mkdir -p ./dist/wasm
	curl -L -o ./dist/wasm/wasm_exec.js https://raw.githubusercontent.com/golang/go/$(shell go version | { read _ _ v _; echo $${v#go}; })/misc/wasm/wasm_exec.js
