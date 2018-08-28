# 关于
	当前目录下的文件是用于生成移动客户端的，移动客户端目前覆盖 Android 和 IOS，通过 HBuilder 直接使用 web app 来生成移动客户端。web app 是 server 端自带的由 vue 实现的客户端（位于 sysmon/web 下）。

# 步骤
## 1. 生成 web app
在 sysmon/web 下执行：
```
npm run build
```
执行成功的话，生成的 sysmon/web 目录即为所需的 web app

Note: 如果是首次执行，则需要先在 sysmon/web 下执行 "npm install" 安装依赖。


## 2. 生成 HBuilder 移动 APP
在当前目录下（sysmon/client/mobile）执行：
```
./mobile.sh gen
```
执行成功的话，可以生成 HBuilder 移动 APP 所需的文件

## 3. 构建移动 APP
通过 HBuilder 客户端打开当前目录，即可作为一个移动 APP 项目来生成移动 APP。
