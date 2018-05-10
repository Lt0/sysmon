var axios = require('axios');
var bapi = require('./backendAPI');
var dd = require('./dataDef');

let updater = null;
let interval = 1000;

function runUpdater(self) {
    axios.get(bapi.infoAll).then(function(res){
        self.rsc = res.data
        //console.log(self.rsc);
        updater = setTimeout(runUpdater, interval, self);
    }).catch(function(err){
        console.log("get info all failed: " + err);
        updater = setTimeout(runUpdater, interval, self);
    })
}
function startUpdater(self){
    if (updater) {
        console.log("resources updater already updated before.");
        return;
    }
    console.log("start resources updater");
    interval = self.interval;
    runUpdater(self);
}

function stopUpdater(){
    clearTimeout(updater);
    updater = null;
}

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;