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
import tips from "../../../js/tips"

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
            for (let i in this.selectedTypeName){
                let name = this.selectedTypeName[i]

                let dynamiclabel = genDynamicLabel(r[name]);

                let dataset = {
                    label: name,
                    dynamiclabel: dynamiclabel,
                    dynamiclabelPercent: r[name].percentRec[0] || 0,
                    tips: tips.resources.legend.mem[name], 
                    ctrl: r[name].ctrl,
                    borderColor: colors[colorIndex++],
                    backgroundColor: ['rgba(0, 0, 0, 0)'],
                    data: r[name].percentRec,
                }
                datasets.push(dataset);
            }
            
            this.dataCollection = {
                labels: p,
                datasets,
            }
        },
    },
}

// function recordMem(records, mem, pointsLen){
//     if (records.length < 1) {
//         initRecords(records, mem, pointsLen);
//     }
// }

// function initRecords(records, mem){
//     for(item in mem)
// }

//normal Type 表示 all 为所有内存，不需要特殊计算的统计类别
let normalTypes = [
    'Buffers', 'Cached', 
    'Active', 'Inactive', 'ActiveAnon', 'InactiveAnon', 'ActiveFile', 'InactiveFile', 'Unevictable', 
    'Mlocked', 'Dirty', 'Writeback', 'AnonPages', 'Mapped', 'Shmem', 
    'Slab', 'SReclaimable', 'SUnreclaim', 
    'KernelStack', 'PageTables', 
    'NfsUnstable', 'Bounce', 'WritebackTmp', 'CommitLimit', 'CommittedAS', 
    'HardwareCorrupted', 
    'AnonHugePages',
    'DirectMap4k', 'DirectMap2M', 'DirectMap1G'
    ];
function Records(){
    this.Memory = new MemType('Memory');
    this.Swap = new MemType('Swap');
    
    this.Buffers = new MemType('Buffers');
    this.Cached = new MemType('Cached');

    this.Active = new MemType('Active');
    this.Inactive = new MemType('Inactive');
    this.ActiveAnon = new MemType('ActiveAnon');
    this.InactiveAnon = new MemType('InactiveAnon');
    this.ActiveFile = new MemType('ActiveFile');
    this.InactiveFile = new MemType('InactiveFile');

    this.Unevictable = new MemType('Unevictable');
    this.Mlocked = new MemType('Mlocked');
    this.Dirty = new MemType('Dirty');
    this.Writeback = new MemType('Writeback');
    
    this.AnonPages = new MemType('AnonPages');
    this.Mapped = new MemType('Mapped');
    this.Shmem = new MemType('Shmem');

    this.Slab = new MemType('Slab');
    this.SReclaimable = new MemType('SReclaimable');
    this.SUnreclaim = new MemType('SUnreclaim');

    this.KernelStack = new MemType('KernelStack');
    this.PageTables = new MemType('PageTables');

    this.NfsUnstable = new MemType('NfsUnstable');
    this.Bounce = new MemType('Bounce');
    this.WritebackTmp = new MemType('WritebackTmp');
    this.CommitLimit = new MemType('CommitLimit');
    this.CommittedAS = new MemType('CommittedAS');

    this.Vmalloc = new MemType('Vmalloc');

    this.HardwareCorrupted = new MemType('HardwareCorrupted');
    this.AnonHugePages = new MemType('AnonHugePages');

    this.HugePages = new MemType('HugePages');

    this.DirectMap4k = new MemType('DirectMap4k');
    this.DirectMap2M = new MemType('DirectMap2M');
    this.DirectMap1G = new MemType('DirectMap1G');
}

// name: 类型名字，/proc/meminfo 的选项名
// rec: 该用途的记录
// allStr: 方便人类阅读的 rec 字符串
// percentRec: 该用途占用内存总量的百分比
// hide: 是否隐藏当前记录的线图
function MemType(name, all, allStr, rec, percentRec, hide, rawData){
    this.name = name;
    this.all = all || 0;
    this.allStr = allStr || null,
    this.rec = rec || [];
    this.percentRec = percentRec || [];
    this.ctrl = {
        hideChart: hide || false,
    }
    this.rawData = rawData;
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
        r.HugePages.allStr = r.HugePages.all;
    }
    updateElements(r.HugePages.rec, m.HugePagesTotal - m.HugePagesFree, p);
    // updateElements(r.HugePages.percentRec, (r.HugePages.rec[0]*100).toFixed(2), p);

    if (r.Vmalloc.all !=  m.VmallocTotal){
        r.Vmalloc.all = m.VmallocTotal;
        r.Vmalloc.allStr = cm.fmtSize.fmtKBSize(r.Vmalloc.all, 2)
    }
    updateElements(r.Vmalloc.rec, m.VmallocUsed, p);

    for(let i in normalTypes){
        let item = normalTypes[i];
        if (r[item].all !=  m.MemTotal){
            r[item].all = m.MemTotal;
            r[item].allStr = cm.fmtSize.fmtKBSize(r[item].all, 2)
        }
        updateElements(r[item].rec, m[item], p);
    }
}

// 根据 rec 记录更新 percentRec
function fmtMem(records, pointsLen){
    let r = records;
    let p = pointsLen;

    for (let item in records){
        if (!r[item].ctrl.hideChart) {
            if (r[item].rec.length - r[item].percentRec.length > 1){
                reFmtPercentRec(r[item].percentRec, r[item].rec, r[item].all);
            } else if (r[item].rec.length - r[item].percentRec.length == 1) {
                updateElements(r[item].percentRec, (r[item].rec[0]/r[item].all*100).toFixed(2), p);
            }
        } else {
            r[item].percentRec = [];
        }
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

function genDynamicLabel(record) {
    let dynamiclabel = "";
    switch(record.name) {
        case 'HugePages':
            dynamiclabel = "0";
            if(record.all != 0) {
                dynamiclabel = `${record.rec} (${record.percentRec[0]}%) of ${record.all} (pages)`;
            }
            break;
        case 'DirectMap4k':
        case 'DirectMap2M':
        case 'DirectMap1G':
        case 'HardwareCorrupted':
            dynamiclabel = cm.fmtSize.fmtKBSize(record.rec[0], 2);
            break;
        default:
        if(record.percentRec[0]) {
            dynamiclabel = `${cm.fmtSize.fmtKBSize(record.rec[0], 2)} (${record.percentRec[0]}%) of ${record.allStr}`;
        } else {
            dynamiclabel = `${cm.fmtSize.fmtKBSize(record.rec[0], 2)} of ${record.allStr})`;
        }
    }
    
    return dynamiclabel;
}
</script>
