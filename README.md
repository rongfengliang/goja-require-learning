# goja with umd modules demo

> demo for load underscore && use go-bindata  for embedding static resources

## how to Running

* install go-bindata

> you can also use go generate command

```code

go get -u github.com/go-bindata/go-bindata/...
```

* generate bingdata.go

```code
go-bindata dalong.js app.js underscore-min.js
```

* build

```code

go build
```

* running

```code
./goja-learning
```