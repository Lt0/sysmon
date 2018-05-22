var axios = require('axios');
var bapi = require('./backendAPI');
var dd = require('./dataDef');

let vColors = require('vuetify/es5/util/colors')
let vc = vColors.default;

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

// 预定义的线条颜色
// let alpha = 1;
// function genColor(r, g, b, a){
//     let c = 'rgba(';
//     c += r + ', ';
//     c += g + ', ';
//     c += b + ', ';
//     c += a + ')';
//     return c;
// }
// let colors = [
//     genColor(0, 150, 136, alpha),
//     genColor(121, 85, 72, alpha),
//     genColor(3, 169, 244, alpha),
    
//     genColor(244, 67, 54, alpha),
//     genColor(139, 195, 74, alpha),
//     genColor(0, 188, 212, alpha),
//     genColor(96, 125, 139, alpha),
//     genColor(103, 58, 183, alpha),
//     genColor(63, 81, 181, alpha),
//     genColor(33, 150, 243, alpha),
    
//     genColor(233, 30, 99, alpha),
//     genColor(156, 39, 176, alpha),
//     genColor(76, 174, 80, alpha),
//     genColor(205, 220, 57, alpha),
//     genColor(255, 235, 59, alpha),
//     genColor(255, 193, 7, alpha),
//     genColor(255, 152, 0, alpha),
//     genColor(255, 87, 34, alpha),
//     genColor(158, 158, 158, alpha),
// ]

function genColorList(){
    let colors = [];
    let colorsLists = [];
    let k = 0;
    for (let i in vc){
        let colorType = vc[i];
        if (k%2 != 0){
            k++;
            continue;
        }
        for (let j in colorType) {
            //console.log("colorType len: " + Object.keys(colorType));
            if (!colorsLists[j]){
                colorsLists[j] = [];
            }
            colorsLists[j].push(colorType[j]);
        }
        k++;
    }

    for (let name in colorsLists){
        if (name == "lighten5" || name == "lighten4" || name == "lighten3" || name == "accent4" || name == "transparent" || name == "white"){
            console.log("name: " + name);
            continue;
        }
        colors = colors.concat(colorsLists[name]);
    }
    return colors;
}

let colors = genColorList();

exports.startUpdater = startUpdater;
exports.stopUpdater = stopUpdater;
exports.colors = colors;