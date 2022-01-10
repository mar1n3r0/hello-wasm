web/app.wasm: main.go
	# GOARCH=wasm GOOS=js go build -o web/app.wasm
	tinygo build -o web/app.wasm

app: main.go
	go build -o app

run: app web/app.wasm
	./app