## Kratos
### 创建项目
#### protoc
https://github.com/protocolbuffers/protobuf/releases/
下载安装包放入环境变量Gopath设置的目录/bin下面
protoc -version 检查是否安装成功
#### protoc-gen-go 
是protoc的一个go工具可以帮我们生成代码
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
#### grpc安装
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
#### 安装Kratos
以下三种选其一
1. go get 安装
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
2. go install 安装
go install github.com/go-kratos/kratos/cmd/kratos/v2
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
3. 源码编译安装
git clone https://github.com/go-kratos/kratos
cd kratos
make install
#### 创建项目
kratos new realworld --创建项目
go mod dowload --下载依赖





