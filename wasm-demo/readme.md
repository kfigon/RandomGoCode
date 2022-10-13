# golang wasm demo

* copy wasm_exec.js
```
$(go env GOROOT)/misc/wasm/wasm_exec.js
```

* create boilerplate html
* create a server that will host the app

* build wasm
```
GOOS=js GOARCH=wasm go build -o public/main.wasm ./public
```
* start the server


### "hot reload"
keep the server running, just rebuild wasm and refresh