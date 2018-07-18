var axios = require('axios');
var bapi = require('./backendAPI');
var dd = require('./dataDef');

let interval = 1000;

function runUpdater(ctrl, api) {
    // console.log("api: " + api);
    axios.get(api).then(function(res){
        ctrl.info = res.data
        //console.log(ctrl.rsc);
        ctrl.updater = setTimeout(runUpdater, interval, ctrl, api);
        // console.log(ctrl.type, ctrl.pid + " updater: ", ctrl.updater);
    }).catch(function(err){
        console.log("get info all failed: " + err);
        ctrl.updater = setTimeout(runUpdater, interval, ctrl, api);
    })
}
function startUpdater(ctrl){
    if (ctrl.updater) {
        console.log("processes updater of this ctrl already updated before.");
        return;
    }

    if(ctrl.interval) {
        interval = ctrl.interval;
    }

    switch (ctrl.type){
        case "all": 
            console.log("run updater for all processes");
            runUpdater(ctrl, bapi.processAll);
            break;
        case "details":
            console.log("run updater for process details of ", ctrl.pid);
            runUpdater(ctrl, bapi.processDetails + "?pid=" + ctrl.pid);
            break;
        default:
            console.log("runUPdater: unknow type: " + type);
            break;
    }
}

function stopUpdater(ctrl){
    if(!ctrl.updater) {
        console.log("updater of this ctrl already stoped before.");
        return;
    }

    console.log("stop processes updater");
    clearTimeout(ctrl.updater);
    ctrl.updater = null;
}

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;