

go get github.com/astaxie/beego
go get github.com/beego/bee

bee api beego-svr2

bee run


app.conf添加如下内容：
EnableAdmin = true
AdminHttpAddr = 0.0.0.0 #默认监听地址是localhost
AdminHttpPort = 8088

http://localhost:8088/qps


http://localhost:8080/web/GetAll

app.conf添加如下内容：
EnableDocs = true

bee run -gendoc=true


beego搭建api服务
http://www.cnblogs.com/ldaniel/p/5490325.html



# goconvey - 课时 1：优雅的单元测试
您可以通过以下两种方式下载安装 GoConvey：

gopm get github.com/smartystreets/goconvey
或
go get github.com/smartystreets/goconvey
























































