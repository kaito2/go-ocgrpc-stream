## なに?
gRPC Stream (`server -> client`) の単純なサンプル

## 環境

`go version go1.11 darwin/amd64`

## setup (memo)
protos はコンパイルされているし、`mod.go`があるから go get する必要もないけど覚書

`grpc`とか`protoc`とかの設定は [ココ](https://budougumi0617.github.io/2018/01/01/hello-grpc-go/) を参考にしてください。


### generate go code

```
$ cd protos
$ protoc -I=./ grpc-stream.proto --go_out=plugins=grpc:./
```

### go get 

```
$ vgo get -u google.golang.org/grpc 
```


## execute

```
# client
$ go run client/main.go

# server
$ go run server/main.go
```
