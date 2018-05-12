<template>
    <div id="cv-container" class="chart-container">
      <line-chart :chart-data="datacollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
    </div>
</template>

<script>
import LineChart from './LineChart.js'
import cm from '../../../js/common'

// 保存内存状态信息的字符串，由 updateGlobalInfo 更新
let totalMem;
let usedMem;
let totalSwap;
let usedSwap;

let tooltipsCallback = {
  // 定制 tooltips 的显示内容
  title: function(tooltipItems, data){
    return data.datasets[tooltipItems[0].datasetIndex].label;
  },
  label: function(tooltipItem, data){
    // data.datasets[tooltipItem.datasetIndex].label 为 x 轴的内容
    let label = data.datasets[tooltipItem.datasetIndex].label || '';

    let total;
    let used;
    if (label == "Memory") {
      total = totalMem;
      used = usedMem;
    } else if (label = "Swap") {
      total = totalSwap;
      used = usedSwap;
    }

    // tooltipItem.yLabel 为 y 轴的内容
    label = tooltipItem.yLabel + "% (" + used + ") of " + total;
    return label;
  },
  
}

// 设置 x 轴坐标点的文字
let xTicksCallback = function(value, index, values){
  if (index == values.length-1){
    return value + "(s)";
  } else {
    return value;
  }
}

// 设置 y 轴坐标点的文字
let yTicksCallback = function(value, index, values){
  return value + "(%)";
}
// 设置 y 轴坐标点
let yTicks = {
  beginAtZero: true,
  max: 100,
  stepSize: 50,
  callback: yTicksCallback,
}

export default {
    name: 'RscMem',
    components: {
        LineChart,
    },
    // rsc 是从 resources 通过 rsc.startUpdater() 获取的系统资源信息数据
    // rscOp 是 resources 传进来的全局图表 options
    props: ['rsc', 'interval', 'rscOp'],
    data () {
        return {
            UsedMem: null,
            SwapMem: null,
        }
    },
    computed: {
      // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
      options: function(){
        let self = this;
        
        let op = Object.assign({}, self.rscOp.chartJsOp);
        op.title.text = "Memory and Swap History";
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
          // `this` 指向 vm 实例
          if (!this.rsc.Mem) {
              return null;
          }
          //console.log("recompute datacollection");
          return {
            labels: this.points,
            datasets: [{
                //label: `Memory (${usedMem} (${this.UsedMem[this.UsedMem.length-1]}%) of ${totalMem})`,
                label: 'Memory',
                borderColor: ['rgba(3, 169, 244, 1)'],
                backgroundColor: ['rgba(3, 169, 244, 0.05)'],
                data: this.UsedMem,
              }, {
                //label: `Swap (${usedSwap} (${this.SwapMem[this.SwapMem.length-1]}%) of ${totalSwap})`,
                label: 'Swap',
                borderColor: ['rgba(0, 150, 136, 1)'],
                backgroundColor: ['rgba(0, 150, 136, 0.05)'],
                data: this.SwapMem,
              }
            ]
          }
      }
    },
    mounted () {
        //this.fillData();
    },
    watch: {
      rsc: function(){
        this.updateUsedMem();
        this.updateSwapMem();
        this.updateGlobalInfo();
      },
    },
    methods: {
      updateUsedMem(){
        let self = this;
        let m = self.rsc.Mem;
        let percent = Math.floor((self.rsc.Mem.UsedMem)/self.rsc.Mem.MemTotal*100);
        let num = this.points.length;
        if (!self.UsedMem){
          self.UsedMem = [];
          for (let i = 0; i < num; i++) {
            self.UsedMem.push(0);
          }
        }
        updateElements(self.UsedMem, percent, self.points.length);
      },
      updateSwapMem(){
        let self = this;
        let val = Math.floor((self.rsc.Mem.SwapTotal - self.rsc.Mem.SwapFree)/self.rsc.Mem.SwapTotal*100);
        let num = this.points.length;
        if (!self.SwapMem){
          self.SwapMem = [];
          for (let i = 0; i < num; i++) {
            self.SwapMem.push(0);
          }
        }
        updateElements(self.SwapMem, val, self.points.length);
      },
      updateGlobalInfo(){
        let self = this;
        if (!totalMem){
          totalMem = cm.fmtSize.fmtKBSize(self.rsc.Mem.MemTotal);
        }

        if (!totalSwap){
          totalSwap = cm.fmtSize.fmtKBSize(self.rsc.Mem.SwapTotal);
        }

        usedMem = cm.fmtSize.fmtKBSize(self.rsc.Mem.UsedMem, 2);
        usedSwap = cm.fmtSize.fmtKBSize(self.rsc.Mem.SwapTotal - self.rsc.Mem.SwapFree, 2);
      }
    },
}

function updateElements(elemList, pointData, pointsLen){
  elemList.unshift(pointData);
  if (elemList.length > pointsLen+1){
    elemList.pop();
  }
}
</script>

<style scoped>
.chart-container {
  position: relative;
  margin: auto;
  width: 99%;
}
</style>