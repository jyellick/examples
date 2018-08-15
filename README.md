# Proto Linking Example

This is to show how you can link two proto files in different packages.

You may compile the protos with the following:

```
protoc --proto_path=$GOPATH/src/github.com/jyellick/examples/protos --go_out=$GOPATH/src $GOPATH/src/github.com/jyellick/examples/protos/common/common.proto
```

```
protoc --proto_path=$GOPATH/src/github.com/jyellick/examples/protos --go_out=$GOPATH/src $GOPATH/src/github.com/jyellick/examples/protos/dependant/dependant.proto
```

You may verify the protos appropriately link by looking at `cmd/example/main.go` and building the binary.
