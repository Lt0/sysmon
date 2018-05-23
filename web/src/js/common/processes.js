var axios = require('axios');
var bapi = require('./backendAPI');
var dd = require('./dataDef');

let updater = null;
let interval = 1000;

// type: all, my, active
function runUpdater(self, type){
    
}

function runUpdater(self, api) {
    // console.log("api: " + api);
    axios.get(api).then(function(res){
        self.info = res.data
        //console.log(self.rsc);
        updater = setTimeout(runUpdater, interval, self, api);
    }).catch(function(err){
        console.log("get info all failed: " + err);
        updater = setTimeout(runUpdater, interval, self, api);
    })
}
function startUpdater(self, type){
    if (updater) {
        console.log("processes updater already updated before.");
        return;
    }

    interval = self.interval;

    switch (type){
        case "all": 
            console.log("run updater for all processes");
            runUpdater(self, bapi.processAll);
            break;
        case "my":
            console.log("run updater for my processes");
            runUpdater(self, bapi.processMy);
            break;
        case "active":
            console.log("run updater for active processes");
            runUpdater(self, bapi.processActive);
            break;
        default:
            console.log("runUPdater: unknow type: " + type);
            break;
    }
}

function stopUpdater(){
    console.log("stop processes updater");
    clearTimeout(updater);
    updater = null;
}

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;