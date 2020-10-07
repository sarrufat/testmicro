module github.com/sarrufat/testmicro/hello

go 1.15

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b // indirect
	github.com/golang/protobuf v1.4.2
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v3 v3.0.0-beta.2.0.20200929133051-87e898f4fc62
	github.com/micro/micro/v3 v3.0.0-beta.5.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
