build:
	@mkdir -p web
	GOARCH=wasm GOOS=js go build -o web/app.wasm

clean:
	@go clean
	@-rm web/app.wasm
