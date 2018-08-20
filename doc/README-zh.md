# sysmon
sysmon 是一个为 Linux 开发的前后端分离的系统资源监视器，在服务器上启动后台服务之后，可以通过任意一个能够访问该服务器的浏览器查看系统资源使用状况。

特性：
1. 单页显示 CPU，内存，磁盘，网络使用状态曲线图
2. 支持单个 CPU 使用率实时显示
3. 支持详细的内存分类显示
4. 支持查看进程详细内容（双击进程显示），包括线程，每个线程的堆栈，NUMA，SMAPS 等信息
5. 强大的进程过滤功能
6. 支持系统存储状态显示
7. 详细丰富的说明内容
8. 可以远程查看系统状态（通过 IP）
9. 移动端友好
10. WEB UI 仅 350K 左右

<br>

## Demo
https://athena-bookshelf.com:2048

http://athena-bookshelf.com:2047

http://45.32.84.249:2047


<br>

## 系统支持情况
以下是经过实际测试能够正常使用的系统，除了这些系统，主流的 Linux 发行版应该都支持。

- Ubuntu 12.04/12.10/14.04/16.04
- Ubuntu Core 16
- RHEL 7.1
- CentOS 7
- CentOS release 6.5 (Final)（系统默认 iptables 禁止外部访问，resources network 显示有问题）
- openSUSE Leap 15.0

<br>

已测不支持的系统：

- CentOS release 5.4 (Final)

<br>
<br>

# 使用
使用基本步骤是下载/编译得到应用安装包 sysmon.tar.gz。将 sysmon.tar.gz 解压到任意目录，进入解压出来的目录，执行 sysmon，然后通过浏览器访问运行的服务。

<br>

## 部署
### 原生应用
将 sysmon.tar.gz 复制到服务器上，解压并运行 sysmon
```
sudo nohup ./sysmon &
```
运行之后，通过浏览器访问 localhost:2048 即可访问服务

注意:
1. 通过这种方式运行的服务，重启设备之后需要重新手动执行上面的命令

<br>

### 容器
注意：容器化开发还未完成，以下步骤进攻测试。

获取 sysmon 的 docker image 之后，执行：
```
sudo docker run --name sysmon --restart=always -d -p 4096:2048 -v /proc:/hproc --privileged -it sysmon
```
运行容器之后，通过浏览器访问 localhost:4096 即可访问服务

<br>

### 开启 https
后端使用的是 beego 框架，要开启 https 功能，在 conf/app.conf 添加如下配置(CertFile 和 KeyFile 是你自己的整数和密钥)：
```
EnableHTTPS = true
HTTPSPort = 2049
HTTPSCertFile = xxx.cert
HTTPSKeyFile = xxx.key
```

更多配置参数参考：
https://beego.me/docs/mvc/controller/config.md


<br>

### 设置 procfs 路径
sysmon 的大多数功能依赖 procfs 提供数据，默认的 procfs 路径为 /proc，如果你的系统比较特殊或者处于其它原因，将 procfs 挂在到了其它地方，那么可以在 conf/app.conf 中配置自己的 procfs 路径：
```
procfs = /myProcPath
```


<br>
<br>

# 开发
## 搭建开发环境
1. golang
2. node(如果 npm 连接速度慢，需要考虑使用 cnpm 替换 npm)
3. musl
```
sudo apt-get install musl-tools
```

## 后端
构建好开发环境之后，在项目的本目录下执行
```
bee run -gendoc=true
```
浏览器访问 localhost:2048/swagger 即可查看后端 API
如果要调试前端，需要参考下面的说明启动前端服务，然后通过浏览器访问 localhost:2047 进行调试

## 前端
参考 web 目录下的 README.md


<br>
<br>

# 发布
首先需要搭建开发环境(参考 开发 部分的说明)

## 原生应用
首次构建的话，需要先进入 web 目录，执行
```
npm run install
```

在仓库代码目录下执行：
```
./pack.sh
```
构建生成的 sysmon.tar.gz 即为需要的本地应用

## 容器
在仓库代码目录下执行：
```
sudo ./build-docker-image.sh
```
构建容器依赖 alpine:3.7，构建完成后，生成的 docker image 为 sysmon:latest


