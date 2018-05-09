var axios = require('axios');
var bapi = require('./backendAPI');
var dd = require('./dataDef');

let interval = 1000;
let updater;

function startUpdater(self){
    axios.get(bapi.infoAll).then(function(res){
        self.rsc = res.data
        console.log(self.rsc);
        updater = setTimeout(startUpdater, interval, self);
    }).catch(function(err){
        console.log("get info all failed: " + err);
        updater = setTimeout(startUpdater, interval, self);
    })
}

// function startUpdater(rsc){
//     axios.get(bapi.infoAll).then(function(res){
//         rsc = res.data
//         console.log(rsc);
//         updater = setTimeout(startUpdater, interval, rsc);
//     }).catch(function(err){
//         console.log("get info all failed: " + err);
//         updater = setTimeout(startUpdater, interval, rsc);
//     })
// }

function stopUpdater(){
    clearTimeout(updater);
}

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;
exports.interval = interval;