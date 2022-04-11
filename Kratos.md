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
go mod download --下载依赖
##### 生成proto代码
创建proto模板
kratos proto add api/user/v1/user.proto
make api 根据对应需求make不同的命令
这时会生成对应的pb.go文件

##### MakeFile
###### 问题
make 提示不是内部命令？
安装MinGw,另外由于Kratos的MakeFile是按照Linux和MAC写的所以不能直接使用cmd。可以采用git的终端进行处理
### Docker
windows docker 无法启动 wsl 版本是1
https://docs.microsoft.com/zh-cn/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package
#### 项目开发步骤
##### api
通过定义Protoc 生成api接口
##### internal/service
创建realworld.go
```go
type RealworldService struct {
	v1.UnimplementedRealworldServer ---对外的Api

	uc  *biz.GreeterUsecase --内部业务逻辑
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewrealworldService(uc *biz.GreeterUsecase, logger log.Logger) *RealworldService {
	return &RealworldService{uc: uc, log: log.NewHelper(logger)}
}
```
service 修改依赖注入
```go
// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewrealworldService)
```
创建user --这个根据自身判断是否创建，这里用来测试接口是否成功
```go
//实现realworld-grpc.pb.go realworld_http.pb.go 接口
func (realworld *RealworldService ) Login(context.Context, *v1.LoginRequest)  (*v1.LoginReply, error) {
	userInfo:= v1.UserInfo{
		Email: "786797661@qq.com",
		Username: "yzc",
		Token:"",
	}
	return &v1.LoginReply{User: &userInfo},nil
}

```
##### 测试
http://127.0.0.1:8752/api/users/login
返回
```json
{
    "user": {
        "email": "786797661@qq.com",
        "token": "",
        "username": "yzc",
        "bio": "",
        "image": ""
    }
}
```
##### 文件结构

### Git
推送新项目
create a new repository 

1、git init
2、git remote add origin https://github.com/786797661/knsh.git







