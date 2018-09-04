var axios = require('axios');
var bapi = require('./backendAPI');

// sysInfo 包含系统静态信息，一般首次打开应用时更新一次即可

// let SysInfo = {};

function getSysInfo(serverInfo) {
    console.log("updateSysInfoAll");
    axios.get(bapi.sysInfoAll).then(function(res){
        serverInfo.SysInfo = res.data
    }).catch(function(err){
        console.log("get SysInfo all failed: " + err);
    })
}

// 获取当前连接的 server （activeServer）的 hostname，保存到 self 参数的 hostname 属性中
function getHostname(self) {
    console.log("getHostname");
    axios.get(bapi.sysInfoHostname).then(function(res){
    	console.log("hostname: ", res.data)
        self.hostname = res.data;
    }).catch(function(err){
        console.log("get SysInfo all failed: " + err);
    })
}

exports.getSysInfo = getSysInfo;
exports.getHostname = getHostname;