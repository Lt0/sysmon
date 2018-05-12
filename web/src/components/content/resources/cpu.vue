<template>
    <line-chart :chart-data="datacollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
</template>

<script>
import LineChart from './LineChart.js'
import cm from '../../../js/common'

// 记录上一次更新的 CPU 数据
let preCpuInfo = {};


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
    stepSize: 20,
    callback: yTicksCallback,
}

export default {
    name: 'RscCpu',
    components: {
        LineChart,
    },
    // rsc 是从 resources 通过 rsc.startUpdater() 获取的系统资源信息数据
    // rscOp 是 resources 传进来的全局图表 options
    props: ['rsc', 'interval', 'rscOp'],
    data () {
        return {
            cores: [],
        }
    },
    computed: {
      // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
      options: function(){
        let self = this;
        
        let op = JSON.parse(JSON.stringify(self.rscOp.chartJsOp));
        op.title.text = "CPU(s) History";
        //op.tooltips.callbacks = tooltipsCallback;
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
        if (!self.rsc.Cpu){
            return null;
        }

        let datasets = [];
        for (let i = 0; i < self.cores.length; i++) {
            let c = self.cores[i];
            let dataset = {
                label: c.name,
                borderColor: colors[i],
                backgroundColor: ['rgba(0, 0, 0, 0)'],
                data: c.rec,
            }
            datasets.push(dataset);
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
        this.updateCores();
      },
    },
    methods: {
        updateCores(){
            let self = this;
            let cpu = self.rsc.Cpu;         
            // 如果没有初始化过 self.cores，则初始化
            if (self.cores.length < 1){
                for (let i = 0; i < cpu.Cores.length; i++) {
                    self.cores.push(new Core(cpu.Cores[i].Name));
                }
            }

            for (let i = 0; i < cpu.Cores.length; i++){
                let c = cpu.Cores[i];
                let percent = 0;
                if (preCpuInfo.Cores) {
                    let p = preCpuInfo.Cores[i]
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
                    percent = ((1-idle/(user + nice + system + idle + iowait + irq + softirq + steal + guest + guestNice))*100).toFixed(2);
                } else {
                    percent = ((1-c.Idle/(c.User + c.Nice + c.System + c.Idle + c.Iowait + c.Irq + c.Softirq + c.Steal + c.Guest + c.GuestNice))*100).toFixed(2);
                }
                
                updateElements(self.cores[i].rec, percent, self.points.length);
            }
            preCpuInfo = Object.assign({}, cpu)
        },
    },
}

// name
function Core(name, rec){
    this.name = name;
    if (!rec) {
        this.rec = [];
    } else {
        this.rec = rec;
    }
}

function updateElements(elemList, pointData, pointsLen){
  elemList.unshift(pointData);
  if (elemList.length > pointsLen+1){
    elemList.pop();
  }
}

// 预定义的线条颜色
let alpha = 1;
function genColor(r, g, b, a){
    let c = 'rgba(';
    c += r + ', ';
    c += g + ', ';
    c += b + ', ';
    c += a + ')';
    return c;
}
let colors = [
    genColor(0, 150, 136, alpha),
    genColor(121, 85, 72, alpha),
    genColor(3, 169, 244, alpha),
    
    genColor(244, 67, 54, alpha),
    genColor(139, 195, 74, alpha),
    genColor(0, 188, 212, alpha),
    genColor(96, 125, 139, alpha),
    genColor(103, 58, 183, alpha),
    genColor(63, 81, 181, alpha),
    genColor(33, 150, 243, alpha),
    
    genColor(233, 30, 99, alpha),
    genColor(156, 39, 176, alpha),
    genColor(76, 174, 80, alpha),
    genColor(205, 220, 57, alpha),
    genColor(255, 235, 59, alpha),
    genColor(255, 193, 7, alpha),
    genColor(255, 152, 0, alpha),
    genColor(255, 87, 34, alpha),
    genColor(158, 158, 158, alpha),
]
</script>

<style scoped>
</style>