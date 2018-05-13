<template>
    <line-chart :chart-data="datacollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
</template>

<script>
import LineChart from './LineChart.js'
import cm from '../../../js/common'
import { fmtKBSize } from '../../../js/common/fmtSize';

// 记录上一次更新的 Net 数据
let preNetInfo = {};
let unit = "";

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
    name: 'RscNet',
    components: {
        LineChart,
    },
    // rsc 是从 resources 通过 rsc.startUpdater() 获取的系统资源信息数据
    // rscOp 是 resources 传进来的全局图表 options
    props: ['rsc', 'interval', 'rscOp'],
    data () {
        return {
            unit: "",
            Nics: [],
        }
    },
    computed: {
      // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
      options: function(){
        let self = this;
        
        let op = JSON.parse(JSON.stringify(self.rscOp.chartJsOp));
        op.title.text = "Network History";
        op.tooltips.callbacks = tooltipsCallback;
        op.scales.yAxes[0].ticks = yTicks;
        op.scales.xAxes[0].ticks.callback = xTicksCallback;
        return op;
      },
      points: function(){
        let self = this;
        let t = this.rscOp.rscChartOp.timeCap;

        let num = Math.floor(t*1000/self.interval);
        if (num < 1) {
          console.log("genSamplePoints: interval too large: " + self.interval);
          return null;
        }

        let ps = [];
        for(let i = 0; i <= num; i++){
          let p = self.interval/1000*i
          ps.push(p.toFixed(1));
        }
        // console.log("recomputed points: " + ps);
        return ps;
      },
      datacollection: function () {
        let self = this;
        if (!self.rsc.Net){
            return null;
        }

        let datasets = [];
        for (let i = 0; i < self.Nics.length; i++) {
            let c = self.Nics[i];
            let rxDataset = {
                label: c.name + '(rx)',
                borderColor: colors[i],
                backgroundColor: ['rgba(0, 0, 0, 0)'],
                data: c.rxRec,
            }
            datasets.push(rxDataset);

            let txDataset = {
                label: c.name + '(tx)',
                borderColor: colors[i],
                backgroundColor: ['rgba(0, 0, 0, 0)'],
                data: c.txRec,
            }
            datasets.push(txDataset);
        }

        return {
            labels: self.points,
            datasets: datasets,
        }
      }
    },
    mounted () {
        //this.fillData();
    },
    watch: {
        rsc: function(){
            this.updateUnit();
            this.updateNICs();
      },
    },
    methods: {
        // 更新 this.unit 和 全局 unit
        updateUnit(){
            if (!preNetInfo.Nics){
                return;
            }
            let self = this;
            let max = 0;
            let net = self.rsc.Net;
            for (let i = 0; i < net.Nics.length; i++){
                let RBytes = net.Nics[i].RBytes - preNetInfo.Nics[i].RBytes;
                let TBytes = net.Nics[i].TBytes - preNetInfo.Nics[i].TBytes;
                if (RBytes > max){
                    max = RBytes;
                }

                if (TBytes > max){
                    max = TBytes;
                }
            }
            self.unit = cm.fmtSize.sizeUnit(max);
            unit = self.unit;
        },
        updateNICs(){
            let self = this;
            let net = self.rsc.Net;         
            // 如果没有初始化过 self.cores，则初始化
            if (self.Nics.length < 1){
                for (let i = 0; i < net.Nics.length; i++) {
                    self.Nics.push(new Nic(net.Nics[i].Name));
                }
            }

            for (let i = 0; i < net.Nics.length; i++){
                // nic current record
                let c = net.Nics[i];
                let percent = 0;

                let RBytes = "";
                let TBytes = "";
                if (preNetInfo.Nics) {
                    // nic previous record
                    let p = preNetInfo.Nics[i];
                    RBytes = c.RBytes - p.RBytes;
                    TBytes = c.TBytes - p.TBytes;
                } else {
                    RBytes = 0;
                    TBytes = 0;
                }
                RBytes = cm.fmtSize.fmtSizeByUnit(RBytes, self.unit);
                TBytes = cm.fmtSize.fmtSizeByUnit(TBytes, self.unit);
                updateElements(self.Nics[i].rxRec, RBytes.toFixed(2), self.points.length);
                updateElements(self.Nics[i].txRec, TBytes.toFixed(2), self.points.length);
            }
            preNetInfo = Object.assign({}, net)
        },
    },
}

// name
function Nic(name, rxRec, txRec){
    this.name = name;

    if (!rxRec) {
        this.rxRec = [];
    } else {
        this.rxRec = rec;
    }

    if (!txRec) {
        this.txRec = [];
    } else {
        this.txRec = txRec;
    }
}

function updateElements(elemList, pointData, pointsLen){
  elemList.unshift(pointData);
  if (elemList.length > pointsLen+1){
    elemList.pop();
  }
}
</script>

<style scoped>
</style>