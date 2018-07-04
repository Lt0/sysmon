var axios = require('axios');
var bapi = require('./backendAPI');

// sysInfo 包含系统静态信息，一般首次打开应用时更新一次即可

let SysInfo = {};

function updateSysInfoAll() {
    console.log("updateSysInfoAll");
    axios.get(bapi.sysInfoAll).then(function(res){
        SysInfo = res.data
        console.log("SysInfo: ", SysInfo);
    }).catch(function(err){
        console.log("get SysInfo all failed: " + err);
    })
}

exports.SysInfo = SysInfo;
exports.updateSysInfoAll = updateSysInfoAll;