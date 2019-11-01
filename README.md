# Go micro Casbin Gin Gorm JWT 用户认证权限系统微服务

这是一个基于 Go Micro 的用户认证和权限的微服务

```
micro new test --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)


protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/test/test.proto

## Configuration

- FQDN: go.micro.srv.test
- Type: srv
- Alias: test

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./test-srv
```

Build a docker image
```
make docker
```