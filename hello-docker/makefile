build:
	@mkdir -p web
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	docker build -t hello-docker .

run: build
	docker run --rm -p 7000:7000 hello-docker

clean:
	@go clean
	@rm ./web/app.wasm
