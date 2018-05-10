<template>
  <div id="vc-box" style="width: 100%">
    <div id="cv-container" class="chart-container">
      <line-chart :chart-data="datacollection" :options="options"></line-chart>
    </div>
  </div>
</template>

<script>
import LineChart from './LineChart.js'
import cm from '../../../js/common'

let totalMem;
let usedMem;
let totalSwap;
let usedSwap;

let tooltipsCallback = {
  // 定制 tooltips 的显示内容
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
    label += ": " + tooltipItem.yLabel + "% (" + used + ") of " + total;
    return label;
  }
}

export default {
    name: 'RscMem',
    components: {
        LineChart,
    },
    props: ['rsc', 'interval', 'lineChartOptions'],
    data () {
        return {
            UsedMem: null,
        }
    },
    computed: {
      options: function(){
        let self = this;
        
        let op = Object.assign({}, self.lineChartOptions);
        op.title.text = "Memory and Swap History";
        //op.scales.yAxes[0].scaleLabel.display = true;
        //op.scales.yAxes[0].scaleLabel.labelString = "Percent(%)";
        //op.scales.xAxes[0].scaleLabel.display = true;
        //op.scales.xAxes[0].scaleLabel.labelString = "Time(seconds)";
        op.tooltips.callbacks = tooltipsCallback;
        return op;
      },
      points: function(){
        let self = this;
        let t = 60;

        let num = Math.floor(t*1000/self.interval);
        if (num < 1) {
          console.log("genSamplePoints: interval too large: " + self.interval);
          return null;
        }

        let ps = [];
        for(let i = 0; i <= num; i++){
          let p = self.interval/1000*i
          ps.unshift(p.toFixed(1));
        }
        ps[0] = ps[0] + "(s)";
        // console.log("recomputed points: " + ps);
        return ps;
      },
      datacollection: function () {
          // `this` 指向 vm 实例
          if (!this.rsc.Disk) {
              return null;
          }
          //console.log("recompute datacollection");
          return {
            labels: this.points,
            datasets: [{
                label: 'Memory',
                borderColor: ['rgba(3, 169, 244, 1)'],
                backgroundColor: ['rgba(3, 169, 244, 0.2)'],
                data: this.UsedMem,
              }, {
                label: 'Swap',
                borderColor: ['rgba(0, 150, 136, 1)'],
                backgroundColor: ['rgba(0, 150, 136, 0.2)'],
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
        let val = Math.floor((self.rsc.Mem.MemTotal - self.rsc.Mem.MemAvailable)/self.rsc.Mem.MemTotal*100);
        // console.log("MemTotal: " + self.rsc.Mem.MemTotal + ", MemAvailable: " + self.rsc.Mem.MemAvailable + ", val: " + val);
        let num = this.points.length;
        if (!self.UsedMem){
          self.UsedMem = [];
          for (let i = 0; i < num; i++) {
            self.UsedMem.push(0);
          }
        }

        self.UsedMem.shift();
        self.UsedMem.push(val);
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

        self.SwapMem.shift();
        self.SwapMem.push(val);
      },
      updateGlobalInfo(){
        let self = this;
        if (!totalMem){
          totalMem = cm.fmtSize.fmtKBSize(self.rsc.Mem.MemTotal);
        }

        if (!totalSwap){
          totalSwap = cm.fmtSize.fmtKBSize(self.rsc.Mem.SwapTotal);
        }

        usedMem = cm.fmtSize.fmtKBSize(self.rsc.Mem.MemTotal - self.rsc.Mem.MemAvailable, 2);
        usedSwap = cm.fmtSize.fmtKBSize(self.rsc.Mem.SwapTotal - self.rsc.Mem.SwapFree, 2);
      }
    },
}

function genSamplePoints(interval){
  let num = 50/interval;

  
}
</script>

<style scoped>
.chart-container {
  position: relative;
  margin: auto;
  width: 99%;
}
</style>