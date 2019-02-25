
<a href="https://github.com/Lt0/sysmon/blob/master/doc/README-zh.md">中文说明</a>


# sysmon
Sysmon is a C/S mode system monitor for Linux distribution. With server side daemon, you can remotely watch usage of your system resources via web browser or clients.

The server side supports common architecture and distribution. In addition to the browser, we also provide client applications for many system such as android/ios/windows/linux/mac os.


<img src="https://raw.githubusercontent.com/Lt0/sysmon/master/doc/img/sysmon-resources.png" />

<img src="https://raw.githubusercontent.com/Lt0/sysmon/master/doc/img/sysmon-mobile.png" />

Features：
1. Watch CPU, memory, disk, network usage in one page
2. Support single CPU usage real-time display
3. Support detailed memory classification display
4. Support for watching process details(double-click process), including threads, stack per thread, NUMA, SMAPS, etc.
5. Powerful process filtering
6. Support system storage status
7. Detailed and rich description
8. Watch System status remotely(via IP)
9. Mobile friendly
10. Small web UI(only about 350KB)
11. Support PWA(you can add the web page to Homescreen)
12. Server side supports common architecture and distribution
13. Clients support common platform
14. Monitor multiple servers by one client

<br>

## Demo

http://3.17.253.160:2048



<br>

## Platform/Arch Support
The server is a daemon service, you can access it via IP:PORT by browser or clients.

### Server Platform
The following are the systems that have been tested. In addition, sysmon should be able to run on most distribution:

- Ubuntu 12.04/12.10/14.04/16.04
- Ubuntu Core 16
- RHEL 7.1
- CentOS 7
- CentOS release 6.5 (Final) (default iptables disables http access，resources network not work)
- openSUSE Leap 15.0

<br>

Unsupported system:

- CentOS release 5.4 (Final)

<br>

### Server Architecture 
The prebuilt package includes the following architecture:

- x86-64 (sysmon-server-amd64-$version.tar.xz)
- x86 (sysmon-server-386-$version.tar.xz)
- arm (sysmon-server-arm-$version.tar.xz)
- arm64 (sysmon-server-arm64-$version.tar.xz)
- mips (sysmon-server-mips-$version.tar.xz)
- mips64 (sysmon-server-mips64-$version.tar.xz)

<br>

### Clients Platform
Mobile:

- Android: sysmon-android-$version.apk
- IOS: sysmon-ios-$version.ipa

Descktop:
- Windows(64 位): sysmon-client-win32-x64-$version.7z
- Windows(32 位): sysmon-client-win32-ia32-$version.7z
- Linux(x86 32 位): sysmon-client-linux-ia32-$version.tar.xz
- Linux(x86 64 位): sysmon-client-linux-x64-$version.tar.xz
- Linux(arm): sysmon-client-linux-armv7l-$version.tar.xz
- Linux(arm64/aarch64): sysmon-client-linux-arm64-$version.tar.xz
- Mac OS: sysmon-client-darwin-x64-$version.tar.xz
- MAS: ysmon-client-mas-x64-$version.tar.xz

<br>
<br>

# Usage
The basic step is to download/compile the application installation package sysmon-server-$arch-$version.tar.xz. Unzip it into any directory, go to the extracted directory, execute sysmon, and then access the running service via browser.

<br>

## Download
https://github.com/Lt0/sysmon/releases

or

https://pan.baidu.com/s/1EKfi2553dXZFvMWCkr60xQ

<br>

## Deploy Server Side
### Native Application
Downloading sysmon-server-$arch-$version.tar.xz to your server and then unpack it and run sysmon
```
tar Jxvf sysmon-server-$arch-$version.tar.xz 
cd sysmon
sudo nohup ./sysmon &
```
After running, access the service by accessing localhost:2048 via a browser.

Note:
You need to manually re-execute the sysmon after restarting the system. If you want to start sysmon when booting, please check start onboot section.

<br>

### Container
Note: The container version is not ready yet, the following steps is just for testing.

Once you get container version of sysmon, run it by this way:
```
sudo docker run --name sysmon --restart=always -d -p 4096:2048 -v /proc:/hproc --privileged -it sysmon
```
After running，access the service by accessing localhost:4096 via a browser.

<br>

### Start onboot
The sysmon server support upstart/systemd framworks to start onboot.If you want to start sysmon on boot, please run install.sh scripts in the extracted directory.

<br>

### Enable https
The server side of sysmon is based on beego framework，you can edit conf/app.conf to enable https (CertFile and KeyFile is your own cert and key)：
```
EnableHTTPS = true
HTTPSPort = 2049
HTTPSCertFile = xxx.cert
HTTPSKeyFile = xxx.key
```

For more details：
https://beego.me/docs/mvc/controller/config.md


<br>

### configure procfs path
sysmon depends on procfs to get system info and default path of procfs is /proc. If your procfs path is not /proc, you can configure your procfs path in conf/app.conf:
```
procfs = /myProcPath
```


<br>
<br>

# Develop
## Setup Develop Environment
Require:
1. golang
2. node
3. musl
```
sudo apt-get install musl-tools
```

## Server Side
Once the development environment is ready, run command in the project's directory.
```
bee run -gendoc=true
```
You can access the backend API by visiting localhost:2048/swagger.

If you want to debug the front end, you need to start front-end service by referring to the following instructions, and then access localhost:2047 for debugging through the browser.

## Front End
See README.md in web directory.


<br>
<br>

# Distribute
Firstly you need to setup a develop environment for sysmon.

## Native Server
For the first build, you need to go to the web directory and execute
```
npm run install
```

And then run command in the repository code directory:
```
./pack.sh
```
This command will generate Native Application sysmon-server-$arch-$version.tar.gz.

### Build for all supported arch
We can only build all supported arch on x64 server, and you need to install musl-tools and docker before building:
```
./pack.sh
``` 

It will generate built package in current directory.

Note:
  It build amd64/i386/arm64/mips64 by local go compiler, but build arm/mips by docker image lightimehpq/golang-386

<br>

## Container Server
Run command in the root path of repository: 
```
mv sysmon-xxx.tar.xz sysmon-latest.tar.xz
sudo ./build-docker-image.sh
```
This command require docker and alpine:3.7 image. It will generate a docker image sysmon:latest.

<br>

## Mobile Client
check <a href="https://github.com/Lt0/sysmon/blob/master/client/mobile/README.md">client/mobile/README.md</a> for details.

<br>

## Desktop Client
check <a href="https://github.com/Lt0/sysmon/blob/master/client/desktop/README.md">client/desktop/README.md</a> for details.


