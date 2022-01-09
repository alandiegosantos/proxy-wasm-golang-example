all: example_go_filter.wasm

mod: go.sum go.mod
	go mod download

example_go_filter.wasm: main.go
ifeq (, $(shell which tinygo))
	docker run --rm -v ${CURDIR}:/src -w /src tinygo/tinygo:0.21.0 tinygo build -o example_go_filter.wasm -scheduler=none -target=wasi main.go
else
	tinygo build -o example_go_filter.wasm -scheduler=none -target=wasi main.go
endif
	

.PHONY: test
test: 
	go test -tags=proxytest ./...

.PHONE: run
run: example_go_filter.wasm
	envoy --config-path ./test/envoy.yaml