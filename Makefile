build-wasm:	./dist/wasm/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o ./dist/wasm/a.wasm ./main.go
	cp ./web/* ./dist/wasm
	cp -r ./assets/ ./dist/wasm/

./dist/wasm/wasm_exec.js:
	mkdir -p ./dist/wasm
	curl -L -o ./dist/wasm/wasm_exec.js https://raw.githubusercontent.com/golang/go/go1.17.5/misc/wasm/wasm_exec.js

serve:
	python -m http.server --bind 127.0.0.1 --directory ./dist/wasm 8080
