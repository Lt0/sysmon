<a href="https://github.com/Lt0/sysmon/blob/master/README.md"> English </a>

# sysmon
sysmon 是一个为 Linux 开发的前后端分离的系统资源监视器，在服务器上启动后台服务之后，可以通过任意一个能够访问该服务器的浏览器或客户端查看系统资源使用状况。

服务端支持常见的硬件架构和 linux 发行版，客户端除了浏览器，还提供了 android/ios/windows/linux/mac os 等系统和架构的本地客户端应用。

<img src="https://raw.githubusercontent.com/Lt0/sysmon/master/doc/img/sysmon-resources.png" />

<img src="https://raw.githubusercontent.com/Lt0/sysmon/master/doc/img/sysmon-mobile.png" />

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
11. 支持 PWA（可以在浏览器上将网页添加到主屏幕，以提供类似本地应用的体验）
12. 服务端支持常见的硬件架构和发行版
13. 客户端覆盖常见平台
14. 在单个客户端上监控多个服务器

<br>

## Demo
http://3.17.253.160:2048


<br>

## 平台/架构支持情况
服务端是一个 daemon 进程，运行之后，可以通过浏览器直接访问服务端的 IP:PORT 来访问，也可以通过客户端应用，设置服务端的 IP:PORT 之后使用。

### 服务端平台支持情况
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

### 服务端架构支持情况
预发布的版本包括以下架构：

- x86-64 (sysmon-server-amd64-$version.tar.xz)
- x86 (sysmon-server-386-$version.tar.xz)
- arm (sysmon-server-arm-$version.tar.xz)
- arm64 (sysmon-server-arm64-$version.tar.xz)
- mips (sysmon-server-mips-$version.tar.xz)
- mips64 (sysmon-server-mips64-$version.tar.xz)

<br>

### 客户端平台支持情况
移动平台
Android：sysmon-android-$version.apk
IOS(需越狱)：sysmon-ios-$version.ipa

桌面平台
- Windows（64 位）：sysmon-client-win32-x64-$version.7z
- Windows（32 位）：sysmon-client-win32-ia32-$version.7z
- Linux（x86 32 位）：sysmon-client-linux-ia32-$version.tar.xz
- Linux（x86 64 位）：sysmon-client-linux-x64-$version.tar.xz
- Linux（arm）：sysmon-client-linux-armv7l-$version.tar.xz
- Linux（arm64/aarch64）：sysmon-client-linux-arm64-$version.tar.xz
- Mac OS（未测试）：sysmon-client-darwin-x64-$version.tar.xz
- MAS（未测试）：sysmon-client-mas-x64-$version.tar.xz

<br>
<br>

# 使用
使用基本步骤是下载/编译得到应用安装包 sysmon-server-$arch-$version.tar.xz。将其解压到任意目录，进入解压出来的目录，执行 sysmon，然后通过浏览器访问运行的服务。

<br>

## 下载
https://github.com/Lt0/sysmon/releases

或

https://pan.baidu.com/s/1EKfi2553dXZFvMWCkr60xQ

<br>

## 部署
### 原生应用
将 sysmon-server-$arch-$version.tar.gz 复制到服务器上，解压并运行 sysmon
```
tar Jxvf sysmon-server-$arch-$version.tar.gz
cd sysmon
sudo nohup ./sysmon &
```
运行之后，通过浏览器访问 localhost:2048 即可访问服务

注意:
1. 通过这种方式运行的服务，重启设备之后需要重新手动执行上面的命令，要让服务开机运行，参考下面的 开机启动 章节

<br>

### 容器
注意：容器化开发还未完成，以下步骤进攻测试。

获取 sysmon 的 docker image 之后，执行：
```
sudo docker run --name sysmon --restart=always -d -p 4096:2048 -v /proc:/hproc --privileged -it sysmon
```
运行容器之后，通过浏览器访问 localhost:4096 即可访问服务

<br>

### 开机启动
目前 sysmon 支持 upstart 和 systemd 两种启动器的操作系统的开机启动，如果需要让 sysmon 开机启动，请在解压出 server 端的安装包之后，执行里边的 install.sh 脚本。

如果该脚本不能识别你的系统，请根据自己的系统启动管理器类型，将 etc/init 下对应的启动配置文件复制到你的操作系统的对应配置目录下。

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

## 服务端
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

## 本地服务端
首次构建的话，需要先进入 web 目录，执行
```
npm run install
```

在仓库代码目录下执行：
```
./pack.sh
```
构建生成的 sysmon-server-$arch-$version.tar.gz 即为需要的本地应用

## 构建全部支持的平台
目前仅支持在 x64 的环境下交叉编译出所有支持平台的服务端，需要先安装 musl-tools 和 docker

```
./pack.sh all
```

编译成功的话会在当前目录下生成各个对应的压缩包。

Note: x86 下编译的 amd64/i386/arm64/mips64 版本是通过本地的 go 编译器编译，arm/mips 两个则是通过 lightimehpq/golang-386 容器镜像来编译的。

<br>

## 容器服务端
在仓库代码目录下执行：
```
mv sysmon-xxx.tar.xz sysmon-latest.tar.xz
sudo ./build-docker-image.sh
```
构建容器依赖 alpine:3.7，构建完成后，生成的 docker image 为 sysmon:latest

<br>

## 移动客户端
参考 <a href="https://github.com/Lt0/sysmon/blob/master/client/mobile/README.md">client/mobile/README.md</a>

<br>

## 桌面客户端
参考 <a href="https://github.com/Lt0/sysmon/blob/master/client/desktop/README.md">client/desktop/README.md</a>

