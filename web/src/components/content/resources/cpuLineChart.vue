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

let preCpuInfo = {};
let colors = cm.rsc.colors;
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
    name: 'CpuChart',
    components: {
        LineChart, chartLegendBar, 
    },
    mounted(){

    },
    props: ['cpu', 'points', 'rscOp', 'selectedCores'],
    data () {
        return {
            records: [],               // 所有 CPU 核心（包括平均值）的记录信息
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
    watch: {
        cpu: function(){
            recordCpu(this.records, this.cpu, this.points.length);
            fmtRec(this.records);
            this.updateDataCollection();
            preCpuInfo = this.cpu;
        },
        selectedCores: function(){
            this.repaintChart();
        },
    },
    methods: {
        repaintChart(){
            fmtRec(this.records);
            this.updateDataCollection();
        },
        updateDataCollection(){
            let datasets = [];
            let colorIndex = 0;

            for (let i in this.records){
                let r = this.records[i];
                if (this.selectedCores.indexOf(r.name) < 0) {
                    continue;
                }

                let dynamiclabel = "";
                if(r.fmtRec[0]) {
                    dynamiclabel = `(${r.fmtRec[0]}%)`
                }

                let dataset = {
                    label: r.name,
                    dynamiclabel: dynamiclabel,
                    dynamiclabelPercent: r.fmtRec[0] || 0,
                    ctrl: r.ctrl,
                    borderColor: colors[colorIndex++],
                    backgroundColor: ['rgba(0, 0, 0, 0)'],
                    data: r.fmtRec,
                }
                datasets.push(dataset);
            }
            
            this.dataCollection = {
                labels: this.points,
                datasets: datasets,
            }

        },
    },
}

function initRec(records, cpu){
    records.push({
        name: 'Avg',
        rec: [0],
        fmtRec: [0.00],
        ctrl: {
            hideChart: false,
        },
    });

    for (let i in cpu.Cores){
        let core = {
            name: cpu.Cores[i].Name,
            rec: [0],
            fmtRec: [0.00],
            ctrl: {
                hideChart: false,
            },
        }
        records.push(core);
    }
}

function recordCpu(records, cpu, pointsLen){
    if (records.length < 1) {
        initRec(records, cpu);
        return;
    }

    let total = 0;
    for (let i = 1; i <= cpu.Cores.length; i++){
        let c = cpu.Cores[i-1];           // current cpu info
        let p = preCpuInfo.Cores[i-1];    // previous cpu info
        
        let user = c.User - p.User;
        let nice = c.Nice - p.Nice;
        let system = c.System - p.System;
        let idle = c.Idle - p.Idle;
        let iowait = c.Iowait - p.Iowait;
        let irq = c.Irq - p.Irq;
        let softirq = c.Softirq - p.Softirq;
        let steal = c.Steal - p.Steal;
        let guest = c.Guest - p.Guest;
        let guestNice = c.GuestNice - p.GuestNice;

        let used = ((1-idle/(user + nice + system + idle + iowait + irq + softirq + steal + guest + guestNice))*100).toFixed(2);
        updateElements(records[i].rec, used, pointsLen);
        total += parseFloat(used);
    }

    updateElements(records[0].rec, (total/cpu.Cores.length).toFixed(2), pointsLen);
}

// 更新绘图数组
function updateElements(elemList, pointData, pointsLen){
  elemList.unshift(pointData);
  if (elemList.length > pointsLen+1){
    elemList.pop();
  }
}

function fmtRec(records){
    for (let i = 0; i < records.length; i ++){
        let r = records[i];
        if (r.ctrl.hideChart){
            r.fmtRec = [];
            continue;
        } else {
            r.fmtRec = r.rec;
        }
    }
}
</script>
