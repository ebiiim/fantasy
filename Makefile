build-wasm:	
	rm -rf ./dist/wasm
	mkdir -p ./dist/wasm
	GOOS=js GOARCH=wasm go build -o ./dist/wasm/a.wasm ./main.go
	curl -L -o ./dist/wasm/wasm_exec.js https://raw.githubusercontent.com/golang/go/go1.17.5/misc/wasm/wasm_exec.js
	cp ./web/* ./dist/wasm
	cp -r ./assets/ ./dist/wasm/
