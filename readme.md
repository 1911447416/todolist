
### linux
``
//运行
nohup ./todo &

//查找进程
ps -ef | grep todo

//杀死进程
kill processNumber


//上传
pscp c:\todolist\todo root@ip:/root

//执行权限
chmod 777 todo


//交叉编译
C:\todolist>SET CGO_ENABLED=0

C:\todolist>SET GOOS=linux

C:\todolist>SET GOARCH=amd64

C:\todolist>go build main.go


``