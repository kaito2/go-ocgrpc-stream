## setup
`grpc`とか`protoc`とかの設定は [ココ](https://budougumi0617.github.io/2018/01/01/hello-grpc-go/) を参考にしてください。

### generate go code

```
$ cd protos
$ protoc -I=./ grpc-stream.proto --go_out=plugins=grpc:./
```

## execute


