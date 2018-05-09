var axios = require('axios');
var bapi = require('./backendAPI');
var dd = require('./dataDef');

let updater;

function startUpdater(self){
    let interval = self.interval;
    axios.get(bapi.infoAll).then(function(res){
        self.rsc = res.data
        //console.log(self.rsc);
        updater = setTimeout(startUpdater, interval, self);
    }).catch(function(err){
        console.log("get info all failed: " + err);
        updater = setTimeout(startUpdater, interval, self);
    })
}

function stopUpdater(){
    clearTimeout(updater);
}

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;