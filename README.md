# sysmon

<br>
<br>

# 使用
## 部署
### 原生应用
将 sysmon.tar.gz 复制到服务器上，解压并运行 sysmon
```
nohup ./sysmon &
```
运行之后，通过浏览器访问 localhost:2048 即可访问服务

注意:
1. 通过这种方式运行的服务，重启设备之后需要重新手动执行上面的命令

### 容器
获取 sysmon 的 docker image 之后，执行：
```
docker run --name sysmon --restart=always -d -p 4096:2048 -it sysmon
```
运行容器之后，通过浏览器访问 localhost:4096 即可访问服务


<br>
<br>

# 开发
## 搭建开发环境
1. golang
2. node(如果 npm 连接速度慢，需要考虑使用 cnpm 替换 npm)
3. musl
```
apt-get install musl-tools
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
./build-docker-image.sh
```
构建容器依赖 alpine:3.7，构建完成后，生成的 docker image 为 sysmon:latest


