<template>
    <div>
        <line-chart :chart-data="datacollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
        <chart-legend-bar :dataCollection='datacollection' @updateChart="repaintChart" class="rsc-legend-bar"/>
    </div>
</template>

<script>
import LineChart from './LineChart.js'
import chartLegendBar from '@/components/content/resources/chartLegendBar'
import cm from '../../../js/common'

// 记录上一次更新的 net 数据
let preNetInfo = null;

// 记录上一次的单位
let preUnit = "B";

// 当前的单位
let unit = "B";

let colors = cm.rsc.colors;
let tooltipsCallback = {
    label: function(tooltipItem, data){
        return data.datasets[tooltipItem.datasetIndex].label + ": " + tooltipItem.yLabel + unit;
    }
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
    return value + unit;
}

let yTicks = {
    beginAtZero: true,
    callback: yTicksCallback,
}

export default {
    name: 'NetChart',
    components: {
        LineChart, chartLegendBar, 
    },
    // net 是从 net.vue 传进来的 rsc.Net 信息
    // rscOp 是 resources 传进来的全局图表 options
    // selectedNics: 当前需要显示状态的网卡
    props: ['net', 'points', 'rscOp', 'selectedNics'],
    data () {
        return {
            nics: [],               // 所有的网卡记录信息
            datacollection: {},
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
    watch: {
        net: function(){
            this.updateNetInfo();
        },
        selectedNics: function(){
            this.repaintChart();
        },
    },
    methods: {
        //重绘图像
        repaintChart(){
            this.forceRefmtSelectedNics();
            this.updateDatacollection();
        },
        // updateNetInfo 用来更新 nics 内的数据
        updateNetInfo(){
            if (!preNetInfo){
                preNetInfo = this.net
            }

            this.recordNics();
            this.updateUnit();
            this.fmtSelectedNics();
            this.updateDatacollection();
            preNetInfo = this.net;
        },
        recordNics(){
            let self = this;
            let nics = self.net.Nics;

            if (self.nics.length < 1){
                self.nics.push(new Nic("Sum"));
                for (let i = 0; i < nics.length; i++){
                    self.nics.push(new Nic(nics[i].Name));
                }
            }

            // 上行总数
            let totalTRate = 0;
            // 下行总数
            let totalRRate = 0;
            // self.nics 的第一个 nic 是 Sum, 所以 self.nics[i] 对应的是 nics[i-1]
            for (let i = 1; i < self.nics.length; i++){
                let c = nics[i-1];
                let p = preNetInfo.Nics[i-1];
                let RBytes = c.RBytes - p.RBytes;
                let TBytes = c.TBytes - p.TBytes;
                totalRRate += RBytes;
                totalTRate += TBytes;
                recordNic(self.nics[i], RBytes, TBytes, self.points.length);

            }
            recordNic(self.nics[0], totalRRate, totalTRate, self.points.length);

            
        },
        // 更新全局 unit
        updateUnit(){
            let self = this;
            let max = 0;
            for (let i = 0; i < self.selectedNics.length; i++){
                let index = self.selectedNics[i];
                let n = self.nics[index];
                let dMax = Math.max(maxItem(n.rRec), maxItem(n.tRec));
                if (dMax > max){
                    max = dMax;
                }
            }
            preUnit = unit;
            unit = cm.fmtSize.sizeUnit(max);
        },
        fmtSelectedNics(){
            let self = this;
            for(let i = 0; i < self.selectedNics.length; i++){
                let index = self.selectedNics[i];
                fmtNic(self.nics[index], self.points.length);
            }
        },
        // 强制全部刷新已选 nics 的 fmtRec
        forceRefmtSelectedNics(){
            let self = this;
            for(let i = 0; i < self.selectedNics.length; i++){
                let index = self.selectedNics[i];
                if(self.nics[index]) {
                    forceFmtNic(self.nics[index]);
                }
            }
        },
        updateDatacollection(){
            let self = this;
            let datasets = [];
            let colorIndex = 0;
            
            for (let i = 0; i < self.selectedNics.length; i++){
                if(!self.nics[self.selectedNics[i]]) {
                    continue;
                }
                let n = self.nics[self.selectedNics[i]];
                let rDataset = {
                    label: n.name + "(rx) ",
                    dynamiclabel: cm.fmtSize.fmtSize(n.rRec[0], 1),
                    ctrl: n.rCtrl,
                    borderColor: colors[colorIndex++],
                    backgroundColor: ['rgba(0, 0, 0, 0)'],
                    data: n.rFmtRec,
                }
                datasets.push(rDataset);

                let tDataset = {
                    label: n.name + "(tx) ",
                    dynamiclabel: cm.fmtSize.fmtSize(n.tRec[0], 1),
                    ctrl: n.tCtrl,
                    borderColor: colors[colorIndex++],
                    backgroundColor: ['rgba(0, 0, 0, 0)'],
                    data: n.tFmtRec,
                }
                datasets.push(tDataset);
            }
            self.datacollection = {
                labels: self.points,
                datasets: datasets,
            }
            
        },
    },
    
}

// name: 网卡名
// rRec：接收速率记录
// tRec: 发送速率记录
// rFmtRec：根据单位格式化后的读取记录
// tFmtRec：根据单位格式化后的写入记录
function Nic(name, rRec, tRec, unit, rFmtRec, hideRRec, tFmtRec, hideTRec){
    this.name = name;
    if (!rRec) {
        this.rRec = [];
    } else {
        this.rRec = rRec;
    }

    if (!tRec) {
        this.tRec = [];
    } else {
        this.tRec = tRec;
    }
    this.rFmtRec = [];
    this.tFmtRec = [];

    this.rCtrl = {
        hideChart: hideRRec || false,
    }

    this.tCtrl = {
        hideChart: hideTRec || false,
    }
}

// 仅记录网口收发速率，不更新绘图部分的数值
function recordNic(nic, rRate, tRate, pointsLen){
    updateElements(nic.rRec, rRate, pointsLen);
    updateElements(nic.tRec, tRate, pointsLen);
}

// 根据 recordNic 记录的 rRec 和 tRec，更新 nic, rFmtRec, tFmtRec
function fmtNic(nic, pointsLen){
    if (!nic.rCtrl.hideChart) {
        if (preUnit != unit || (nic.rRec.length-1)>(nic.rFmtRec.length)){
            reFmtRec(nic.rRec, nic.rFmtRec, unit);
        } else {
            updateElements(nic.rFmtRec, cm.fmtSize.fmtSizeByUnit(nic.rRec[0], unit).toFixed(2), pointsLen);
        }
    } else {
        nic.rFmtRec = [];
    }

    if (!nic.tCtrl.hideChart){
        if (preUnit != unit || (nic.tRec.length-1)>(nic.tFmtRec.length)){
            reFmtRec(nic.tRec, nic.tFmtRec, unit);
        } else {
            updateElements(nic.tFmtRec, cm.fmtSize.fmtSizeByUnit(nic.tRec[0], unit).toFixed(2), pointsLen);
        }
    } else {
        nic.tFmtRec = [];
    }
}

function forceFmtNic(nic){
    //console.log("forceFmtNic");
    if (!nic.rCtrl.hideChart) {
        reFmtRec(nic.rRec, nic.rFmtRec, unit);
    } else {
        nic.rFmtRec = [];
    }
console.log("forceFmtNic");
    if (!nic.tCtrl.hideChart) {
        reFmtRec(nic.tRec, nic.tFmtRec, unit);
    } else {
         nic.tFmtRec = [];
    }
}

function maxItem(list){
    let max = list[0];
    for (let i = 1; i < list.length; i++){
        if (list[i] > max){
            max = list[i]
        }
    }
    return max
}

// 根据 unit 格式化一个 fmtRec
function reFmtRec(rec, fmtRec, unit){
    for (let i = 0; i < rec.length; i++){
        fmtRec[i] = cm.fmtSize.fmtSizeByUnit(rec[i], unit).toFixed(2);
    }
}

// 更新绘图数组
function updateElements(elemList, pointData, pointsLen){
  elemList.unshift(pointData);
  if (elemList.length > pointsLen+1){
    elemList.pop();
  }
}
</script>

<style scoped>
</style>