### GRPC使用

* 安装protobuf编译工具
```
$ wget https://github.com/google/protobuf/releases/download/v3.2.0/protobuf-php-3.2.0.tar.gz
$ tar -zxvf protobuf-php-3.2.0.tar.gz
$ cd protobuf-php-3.2.0
$ ./configure --prefix=/usr/local/protobuf
$ sudo make 
$ sudo make install

检测是否安装成功:
/usr/local/protobuf/bin/protoc –version

安装成功后，请将protoc可执行文件copy或者软链到/usr/local/bin和/usr/bin目录下
```

* 安装golang依赖包
```
$ git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
$ git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
$ git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
以上三个步骤因为网络原因可能下载不下来，可以直接下载zip包，然后解压到对应的目录下
 
// 安装 Protoc Golang 插件
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
 
$ git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
以上步骤也可以直接下载zip包，然后解压到对应的目录下
 
$ cd $GOPATH/src/
$ go install google.golang.org/grpc

```