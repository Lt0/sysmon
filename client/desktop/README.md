# 说明
当前目录下的工具可以使用 ../../web/dist 中的静态 web app 来构建各个桌面平台的客户端

## 支持的平台和架构

<br>

# 安装构建以来的工具
```
apt-get install wine p7zip-full
```

<br>

# 构建 sysmon 客户端
在当前目录执行：
```
npm run build
  -- 构建当前平台和硬件架构的客户端

npm run build -- --all
  -- 构建全部支持的平台和硬件架构的客户端
```
构建成功之后会将生成的客户端放在 build 目录下的对应目录中，build 下的目录以 sysmon-$平台-$架构 命名。

注意：在构建 sysmon 的客户端之前需要确保 ../../web/dist 目录是最新生成的 web app，如果不是，请先进入 ../../web 目录执行 "npm run build" 以生成最新的 web app。

<br>

# 打包
上面的步骤只是生成了应用，每个应用放在 build/sysmon-$平台-$架构 目录下，并没有版本信息，也没有打包成单个文件，下面的命令可以将 build 下的应用根据当前版本来打包为压缩包。

```
npm run compress
```
