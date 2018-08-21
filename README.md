
<a href="https://github.com/Lt0/sysmon/blob/master/doc/README-zh.md">中文说明</a>


# sysmon
Sysmon is a B/S mode system monitor for Linux distribution. With server side daemon, you can remotely watch usage of your system resources via web browser.

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

<br>

## Demo
https://athena-bookshelf.com:2048

http://athena-bookshelf.com:2047

http://45.32.84.249:2047


<br>

## System Support
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

## Architecture Support
The prebuilt package includes the following architecture:

- x86-64 (sysmon-xxx-amd64.tar.gz)
- x86 (sysmon-xxx-386.tar.gz)
- ARMv6 (sysmon-xxx-armv6l.tar.gz)
- ARMv8 (sysmon-xxx-arm64.tar.gz)


<br>
<br>

# Usage
The basic step is to download/compile the application installation package sysmon.tar.gz. Unzip sysmon.tar.gz into any directory, go to the extracted directory, execute sysmon, and then access the running service via browser.

<br>

## Download
https://pan.baidu.com/s/1EKfi2553dXZFvMWCkr60xQ

<br>

## Deploy Server Side
### Native Application
Downloading sysmon.tar.gz to your server and then unpack it and run sysmon
```
tar zxvf sysmon.tar.gz
cd sysmon
sudo nohup ./sysmon &
```
After running, access the service by accessing localhost:2048 via a browser.

Note:
You need to manually re-execute the sysmon after restarting the system.

<br>

### Container
Note: The container version is not ready yet, the following steps is just for testing.

Once you get container version of sysmon, run it by this way:
```
sudo docker run --name sysmon --restart=always -d -p 4096:2048 -v /proc:/hproc --privileged -it sysmon
```
After running，access the service by accessing localhost:4096 via a browser.


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

## Native Application
For the first build, you need to go to the web directory and execute
```
npm run install
```

And then run command in the repository code directory:
```
./pack.sh
```
This command will generate Native Application sysmon.tar.gz.

## Container
Run command in the root path of repository: 
```
sudo ./build-docker-image.sh
```
This command require docker and alpine:3.7 image. It will generate a docker image sysmon:latest.


