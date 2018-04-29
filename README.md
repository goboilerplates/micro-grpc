# Boilerplate for RPC Microservice in Go
[![Build Status](https://travis-ci.org/goboilerplates/micro-rpc.svg?branch=master)](https://travis-ci.org/goboilerplates/micro-rpc)
[![codecov](https://codecov.io/gh/goboilerplates/micro-rpc/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/micro-rpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/micro-rpc)](https://goreportcard.com/report/github.com/goboilerplates/micro-rpc)
[![GoDoc](https://godoc.org/github.com/goboilerplates/micro-rpc?status.svg)](https://godoc.org/github.com/goboilerplates/micro-rpc)
[![](https://images.microbadger.com/badges/image/goboilerplates/micro-rpc.svg)](https://microbadger.com/images/goboilerplates/micro-rpc)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/micro-rpc/blob/master/LICENSE)

## Features
- Unary, Server-side Streaming and Client-side Streaming RPCs API
- Bidirectional Streaming RPC
- CI with Travis
- Docker Build

## Installation

Get the micro-rpc repository

```
go get github.com/goboilerplates/micro-rpc

cd echo $GOPATH/src/github.com/goboilerplates/micro-rpc
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Generate Go Protocol Buffers
```
protoc proto/main.proto --go_out=plugins=grpc:./
```

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve RPC at localhost:50052
```

Build and run native binary

``` bash
bash script/Build.sh

./micro-rpc.out
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash script/BuildMulti.sh
```
## Docker support 

Build docker image

```
bash script/Dockerbuild.sh
```

Run docker container

```
docker run -d --name micro-rpc -p 50052:50052 goboilerplates/micro-rpc
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/micro-rpc/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

