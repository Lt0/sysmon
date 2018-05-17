<template>
    <div>
        <line-chart :chart-data="dataCollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
        <chart-legend-bar :dataCollection='dataCollection' @updateChart="repaintChart" class="rsc-legend-bar"/>
    </div>
</template>

<script>
import LineChart from './LineChart.js'
import chartLegendBar from '@/components/content/resources/chartLegendBar'
import cm from '../../../js/common'

let colors = cm.rsc.colors;
// let tooltipsCallback = {
//     label: function(tooltipItem, data){
//         return data.datasets[tooltipItem.datasetIndex].label + ": " + tooltipItem.yLabel;
//     }
// }

let tooltipsCallback = {
  label: function(tooltipItem, data){
    // data.datasets[tooltipItem.datasetIndex].label 为 x 轴的内容
    let label = data.datasets[tooltipItem.datasetIndex].label || '';
    // tooltipItem.yLabel 为 y 轴的内容
    return label + ": " + tooltipItem.yLabel + "%";
  },
}

// 设置 x 轴坐标点的文字
let xTicksCallback = function(value, index, values){
    if (index == values.length-1){
        return value + "s";
    } else {
        return value;
    }
}

// 设置 y 轴坐标点的文字
let yTicksCallback = function(value, index, values){
    return value + "%";
}
// 设置 y 轴坐标点
let yTicks = {
    beginAtZero: true,
    max: 100,
    stepSize: 50,
    callback: yTicksCallback,
}

export default {
    name: 'MemChart',
    components: {
        LineChart, chartLegendBar, 
    },
    props: ['mem', 'points', 'rscOp', 'selectedTypeName'],
    data () {
        return {
            records: new Records(),               // 所有的内存的用途记录信息
            dataCollection: {},
        }
    },
    computed: {
        // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
        options: function(){
            let self = this;
            let op = JSON.parse(JSON.stringify(self.rscOp.chartJsOp));
            op.tooltips.callbacks = tooltipsCallback;
            op.scales.yAxes[0].ticks = yTicks;
            op.scales.xAxes[0].ticks.callback = xTicksCallback;
            return op;
        },
    },
    // mounted(){
    //     console.log("mounted: mem: " + JSON.stringify(this.mem));
    // },
    watch: {
        mem: function(){
            // console.log("watch mem: " + JSON.stringify(this.mem));
            recordMem(this.records, this.mem, this.points.length);
            fmtMem(this.records, this.points.length);
            this.updateDataCollection();
        },
        selectedTypeName: function(){
            this.repaintChart();
        },
    },
    methods: {
        repaintChart(){
            // console.log("repaintChart");
            fmtMem(this.records, this.points.length);
            this.updateDataCollection();
        },
        updateDataCollection(){
            let r = this.records;
            let p = this.points;

            let colorIndex = 0;
            let datasets = [];
            // console.log("r.Memory.hide: " + r.Memory.hide);
            for (let i = 0; i < this.selectedTypeName.length; i++){
                if (this.selectedTypeName[i] == 'Memory') {
                    let dataset = {
                        label: 'Memory',
                        dynamiclabel: `${cm.fmtSize.fmtKBSize(r.Memory.rec[0], 2)} (${r.Memory.percentRec[0]}%) of ${r.Memory.allStr})`,
                        ctrl: r.Memory.ctrl,
                        borderColor: colors[colorIndex++],
                        backgroundColor: ['rgba(0, 0, 0, 0)'],
                        data: r.Memory.percentRec,
                    }
                    datasets.push(dataset);
                    continue;
                }
                if (this.selectedTypeName[i] == 'Swap') {
                    let dataset = {
                        label: 'Swap',
                        dynamiclabel: `${cm.fmtSize.fmtKBSize(r.Swap.rec[0], 2)} (${r.Swap.percentRec[0]}%) of ${r.Swap.allStr})`,
                        ctrl: r.Swap.ctrl,
                        borderColor: colors[colorIndex++],
                        backgroundColor: ['rgba(0, 0, 0, 0)'],
                        data: r.Swap.percentRec,
                    }
                    datasets.push(dataset);
                    continue;
                }
                if (this.selectedTypeName[i] == 'HugePages') {
                    let dataset = {
                        label: 'HugePages',
                        dynamiclabel: `${cm.fmtSize.fmtKBSize(r.HugePages.rec[0], 2)} (${r.HugePages.percentRec[0]}%) of ${r.HugePages.allStr})`,
                        ctrl: r.HugePages.ctrl,
                        borderColor: colors[colorIndex++],
                        backgroundColor: ['rgba(0, 0, 0, 0)'],
                        data: r.HugePages.percentRec,
                    }
                    datasets.push(dataset);
                    continue;
                }
            }
            
            this.dataCollection = {
                labels: p,
                datasets,
            }
        },
    },
}

function Records(){
    this.Memory = new MemType('Memory');
    this.Swap = new MemType('Swap');
    this.HugePages = new MemType('HugePages');
}

// name: 类型名字，/proc/meminfo 的选项名
// rec: 该用途的记录
// allStr: 方便人类阅读的 rec 字符串
// percentRec: 该用途占用内存总量的百分比
// hide: 是否隐藏当前记录的线图
function MemType(name, all, allStr, rec, percentRec, hide){
    this.name = name;
    this.all = all || 0;
    this.allStr = allStr || null,
    this.rec = rec || [];
    this.percentRec = percentRec || [];
    this.ctrl = {
        hideChart: hide || false,
    }
}

function recordMem(records, currentMem, pointsLen){
    let r = records;
    let m = currentMem;
    let p = pointsLen;

    if (r.Memory.all != m.MemTotal) {
        r.Memory.all = m.MemTotal;
        r.Memory.allStr = cm.fmtSize.fmtKBSize(r.Memory.all, 2)
    }
    updateElements(r.Memory.rec, m.UsedMem, p);
    // 
    
    if (r.Swap.all !=  m.SwapTotal){
        r.Swap.all = m.SwapTotal;
        r.Swap.allStr = cm.fmtSize.fmtKBSize(r.Swap.all, 2)
    }
    updateElements(r.Swap.rec, m.SwapTotal - m.SwapFree, p);
    // updateElements(r.Swap.percentRec, (r.Swap.rec[0]*100).toFixed(2), p);

    if (r.HugePages.all !=  m.HugePagesTotal){
        r.HugePages.all = m.HugePagesTotal;
        r.HugePages.allStr = cm.fmtSize.fmtKBSize(r.HugePages.all, 2)
    }
    updateElements(r.HugePages.rec, m.HugePagesTotal - m.HugePagesFree, p);
    // updateElements(r.HugePages.percentRec, (r.HugePages.rec[0]*100).toFixed(2), p);
}

// 根据 rec 记录更新 percentRec
function fmtMem(records, pointsLen){
    let r = records;
    let p = pointsLen;

    // console.log("r.Memory.rec.length: " + r.Memory.rec.length);
    if (!r.Memory.ctrl.hideChart) {
        if (r.Memory.rec.length - r.Memory.percentRec.length > 1){
            reFmtPercentRec(r.Memory.percentRec, r.Memory.rec, r.Memory.all);
        } else if (r.Memory.rec.length - r.Memory.percentRec.length == 1) {
            updateElements(r.Memory.percentRec, (r.Memory.rec[0]/r.Memory.all*100).toFixed(2), p);
        }
    } else {
        r.Memory.percentRec = [];
    }

    if (!r.Swap.ctrl.hideChart) {
        if (r.Swap.rec.length - r.Swap.percentRec.length > 1){
            reFmtPercentRec(r.Swap.percentRec, r.Swap.rec, r.Swap.all);
        } else if (r.Swap.rec.length - r.Swap.percentRec.length == 1){
            updateElements(r.Swap.percentRec, (r.Swap.rec[0]/r.Swap.all*100).toFixed(2), p);
        }
    } else {
        r.Swap.percentRec = [];
    }

    if (!r.HugePages.ctrl.hideChart) {
        if (r.HugePages.rec.length - r.HugePages.percentRec.length > 1){
            reFmtPercentRec(r.HugePages.percentRec, r.HugePages.rec, r.HugePages.all);
        } else if (r.HugePages.rec.length - r.HugePages.percentRec.length == 1) {
            updateElements(r.HugePages.percentRec, (r.HugePages.rec[0]/r.HugePages.all*100).toFixed(2), p);
        }
    } else {
        r.HugePages.percentRec = [];
    }
}

// 更新绘图数组
function updateElements(elemList, pointData, pointsLen){
  elemList.unshift(pointData);
  if (elemList.length > pointsLen+1){
    elemList.pop();
  }
}

function reFmtPercentRec(percentRec, rec, all){
    for (let i = 0; i < rec.length; i++){
        percentRec[i] = (rec[i]/all*100).toFixed(2)
    }
}


// // 根据 recordNic 记录的 rRec 和 tRec，更新 nic, rFmtRec, tFmtRec
// function fmtNic(nic, pointsLen){
//     if (!nic.rCtrl.hideChart) {
//         if (preUnit != unit || (nic.rRec.length-1)>(nic.rFmtRec.length)){
//             reFmtRec(nic.rRec, nic.rFmtRec, unit);
//         } else {
//             updateElements(nic.rFmtRec, cm.fmtSize.fmtSizeByUnit(nic.rRec[0], unit).toFixed(2), pointsLen);
//         }
//     } else {
//         nic.rFmtRec = [];
//     }

//     if (!nic.tCtrl.hideChart){
//         if (preUnit != unit || (nic.tRec.length-1)>(nic.tFmtRec.length)){
//             reFmtRec(nic.tRec, nic.tFmtRec, unit);
//         } else {
//             updateElements(nic.tFmtRec, cm.fmtSize.fmtSizeByUnit(nic.tRec[0], unit).toFixed(2), pointsLen);
//         }
//     } else {
//         nic.tFmtRec = [];
//     }
// }





// // 仅记录网口收发速率，不更新绘图部分的数值
// function recordNic(nic, rRate, tRate, pointsLen){
//     updateElements(nic.rRec, rRate, pointsLen);
//     updateElements(nic.tRec, tRate, pointsLen);
// }

// // 根据 recordNic 记录的 rRec 和 tRec，更新 nic, rFmtRec, tFmtRec
// function fmtNic(nic, pointsLen){
//     if (!nic.rCtrl.hideChart) {
//         if (preUnit != unit || (nic.rRec.length-1)>(nic.rFmtRec.length)){
//             reFmtRec(nic.rRec, nic.rFmtRec, unit);
//         } else {
//             updateElements(nic.rFmtRec, cm.fmtSize.fmtSizeByUnit(nic.rRec[0], unit).toFixed(2), pointsLen);
//         }
//     } else {
//         nic.rFmtRec = [];
//     }

//     if (!nic.tCtrl.hideChart){
//         if (preUnit != unit || (nic.tRec.length-1)>(nic.tFmtRec.length)){
//             reFmtRec(nic.tRec, nic.tFmtRec, unit);
//         } else {
//             updateElements(nic.tFmtRec, cm.fmtSize.fmtSizeByUnit(nic.tRec[0], unit).toFixed(2), pointsLen);
//         }
//     } else {
//         nic.tFmtRec = [];
//     }
// }

// function forceFmtNic(nic){
//     console.log("forceFmtNic");
//     if (!nic.rCtrl.hideChart) {
//         reFmtRec(nic.rRec, nic.rFmtRec, unit);
//     } else {
//         nic.rFmtRec = [];
//     }
// console.log("forceFmtNic");
//     if (!nic.tCtrl.hideChart) {
//         reFmtRec(nic.tRec, nic.tFmtRec, unit);
//     } else {
//          nic.tFmtRec = [];
//     }
// }

// function maxItem(list){
//     let max = list[0];
//     for (let i = 1; i < list.length; i++){
//         if (list[i] > max){
//             max = list[i]
//         }
//     }
//     return max
// }

// // 根据 unit 格式化一个 fmtRec
// function reFmtRec(rec, fmtRec, unit){
//     for (let i = 0; i < rec.length; i++){
//         fmtRec[i] = cm.fmtSize.fmtSizeByUnit(rec[i], unit).toFixed(2);
//     }
// }



// let tooltipsCallback = {
//     label: function(tooltipItem, data){
//         return data.datasets[tooltipItem.datasetIndex].label + ": " + tooltipItem.yLabel + unit;
//     }
// }

// // 保存内存状态信息的字符串，由 updateGlobalInfo 更新
// let totalMem;
// let usedMem;
// let totalSwap;
// let usedSwap;

// let tooltipsCallback = {
//   // 定制 tooltips 的显示内容
//   title: function(tooltipItems, data){
//     return data.datasets[tooltipItems[0].datasetIndex].label;
//   },
//   label: function(tooltipItem, data){
//     // data.datasets[tooltipItem.datasetIndex].label 为 x 轴的内容
//     let label = data.datasets[tooltipItem.datasetIndex].label || '';

//     let total;
//     let used;
//     if (label == "Memory") {
//       total = totalMem;
//       used = usedMem;
//     } else if (label = "Swap") {
//       total = totalSwap;
//       used = usedSwap;
//     }

//     // tooltipItem.yLabel 为 y 轴的内容
//     label = tooltipItem.yLabel + "% (" + used + ") of " + total;
//     return label;
//   },
  
// }

// // 设置 x 轴坐标点的文字
// let xTicksCallback = function(value, index, values){
//   if (index == values.length-1){
//     return value + "s";
//   } else {
//     return value;
//   }
// }

// // 设置 y 轴坐标点的文字
// let yTicksCallback = function(value, index, values){
//   return value + "%";
// }
// // 设置 y 轴坐标点
// let yTicks = {
//   beginAtZero: true,
//   max: 100,
//   stepSize: 50,
//   callback: yTicksCallback,
// }
</script>
