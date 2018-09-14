var axios = require('axios');
var bapi = require('./backendAPI');

let updater = null;
let interval = 1000;

function runUpdater(self) {
    axios.get(bapi.infoDisk).then(function(res){
        self.info = res.data
        //console.log(self.rsc);
        updater = setTimeout(runUpdater, interval, self);
    }).catch(function(err){
        console.log("get disk info failed: " + err);
        self.infoErr = err;
        updater = setTimeout(runUpdater, interval, self);
    })
}
function startUpdater(self){
    if (updater) {
        console.log("fs updater already updated before.");
        return;
    }
    console.log("start fs updater");
    interval = self.interval;
    runUpdater(self);
}

function stopUpdater(){
    console.log("stop fs updater.");
    clearTimeout(updater);
    updater = null;
}

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;