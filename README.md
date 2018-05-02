# sysmon

<br>

## 构建
### 原生应用
在仓库代码目录下执行：
```
./pack.sh
```
构建生成的 sysmon.tar.gz 即为需要的本地应用

### 容器
在仓库代码目录下执行：
```
./build-docker-image.sh
```
构建容器依赖 alpine:3.7，构建完成后，生成的 docker image 为 sysmon:latest


<br>

## 部署
### 原生应用
直接将打包生成 sysmon.tar.gz 复制到服务器上，解压并运行 sysmon
```
nohup ./sysmon &
```

### 容器
获取 sysmon 的 docker image 之后，执行：
```
docker run --name sysmon --restart=always -d -p 4096:2048 -it sysmon
```


<br>

## 开发
### 后端
```
bee run -gendoc=true
```

浏览器访问 localhost:2048 即可访问默认网页
浏览器访问 localhost:2048/swagger 即可查看后端 API

### 前端
参考 web 目录下的 README.md

