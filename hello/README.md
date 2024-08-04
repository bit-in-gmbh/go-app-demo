# hello

hello is a demo that shows how to use the [go-app package](https://github.com/maxence-charriere/go-app) to build a GUI.

## Build

Go to the hello directory:

```sh
cd ./go-app-demo/hello
```

### Build the client

```sh
GOARCH=wasm GOOS=js go build -o web/app.wasm
```

### Build the server

```sh
go build
```

### Run the app
```sh
./hello
```

## Contribute

Help to develop the [go-app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
