### 完成编译和二进制程序的安装
```shell script

$make 
go build -o myapp main.go config.go

$sudo make install
cp ./myapp /usr/local/bin
$sudo make install-cfg
mkdir -p /etc/myapp
cp ./config.ini /etc/myapp
```

### 将myapp安装为systemd的服务

```shell script
$sudo ./myapp install -user tonybai -workingdir /home/tonybai
install operation ok

$sudo systemctl status myApp
myApp.service - This is a Go daemon service demo
     Loaded: loaded (/etc/systemd/system/myApp.service; enabled; vendor preset: enabled)
     Active: inactive (dead)

```