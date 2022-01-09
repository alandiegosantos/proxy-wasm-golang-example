## Proxy Wasm Golang Example

This repo is an example of a WASM filter implemented in Golang. This filter only adds the header _x-wasm-custom_ with the value _FOO_.

### Building the filter

This build requires [TinyGo](https://tinygo.org/). The Makefile also supports using TinyGo directly or use a docker image to build the wasm filter.

```
$ make
```

To run Envoy with this filter:
```
$ make run
```

Then, to test the filter:
```
$ curl http://localhost:18000 -v
* Connected to localhost (127.0.0.1) port 18000 (#0)
> GET / HTTP/1.1
> Host: localhost:18000
> User-Agent: curl/7.77.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< content-length: 13
< content-type: text/plain
< date: Sun, 09 Jan 2022 17:18:09 GMT
< server: envoy
< x-envoy-upstream-service-time: 0
< x-wasm-custom: FOO
<
example body
* Connection #0 to host localhost left intact

```

To run the unit tests, run:

```
$ make test
```
