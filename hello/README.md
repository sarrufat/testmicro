# Github.Com/Sarrufat/Testmicro/Hello Service

This is the Github.Com/Sarrufat/Testmicro/Hello service

Generated with

```
micro new github.com/sarrufat/testmicro/hello
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- Alias: github.com/sarrufat/testmicro/hello

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./github.com/sarrufat/testmicro/hello
```

Build a docker image
```
make docker
```