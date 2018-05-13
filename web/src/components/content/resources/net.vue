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
            rx: 0,      // 当前的接收速率，单位为 byte
            tx: 0,      // 当前的发送速率，单位为 byte
            rxRec: [],  // 要显示的接收速率记录
            txRec: [],  // 要显示的发送速率记录
            // Nics: [],
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
        return {
            labels: self.points,
            datasets: [
                {
                    label: 'Receiving',
                    borderColor: ['rgba(3, 169, 244, 1)'],
                    backgroundColor: ['rgba(3, 169, 244, 0)'],
                    data: self.rxRec,
                },{
                    label: 'Sending',
                    borderColor: ['rgba(0, 150, 136, 1)'],
                    backgroundColor: ['rgba(3, 169, 244, 0)'],
                    data: self.txRec,
                },
            ],
        }
      }
    },
    mounted () {
    },
    watch: {
        rsc: function(){
            this.updateRxTx();
            this.updateUnit();
            this.updateRec();
            preNetInfo = this.rsc.Net;
      },
    },
    methods: {
        updateRxTx(){
            if (!preNetInfo.Nics){
                return;
            }

            let self = this;
            let net = self.rsc.Net;

            self.rx = 0;
            self.tx = 0;
            for (let i = 0; i < net.Nics.length; i++){
                self.rx += net.Nics[i].RBytes - preNetInfo.Nics[i].RBytes;
                self.tx += net.Nics[i].TBytes - preNetInfo.Nics[i].TBytes;
            }
        },
        // 更新 this.unit 和 全局 unit
        updateUnit(){
            unit = cm.fmtSize.sizeUnit(Math.max(this.rx, this.tx));
            this.unit = unit;
        },
        updateRec(){
            let self = this;
            let net = self.rsc.Net;

            let rx = cm.fmtSize.fmtSizeByUnit(self.rx, self.unit);
            let tx = cm.fmtSize.fmtSizeByUnit(self.tx, self.unit);

            updateElements(self.rxRec, rx.toFixed(2), self.points.length);
            updateElements(self.txRec, tx.toFixed(2), self.points.length);
        }
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