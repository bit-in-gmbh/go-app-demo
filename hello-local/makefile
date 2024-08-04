build:
	@mkdir -p web
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: build
	./hello-local

clean:
	@go clean -v ./...
	@-rm web/app.wasm
