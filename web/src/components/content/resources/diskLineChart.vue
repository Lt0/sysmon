<template>
    <div>
        <line-chart :chart-data="datacollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
        <div class="rsc-legend-container">
            <div v-for="dataset in datacollection.datasets" :key="dataset.label">
                <chart-legend 
                    :borderColor=dataset.borderColor 
                    :backgroundColor=dataset.backgroundColor
                    :label=dataset.label
                    :dynamicLabel=dataset.dynamiclabel
                    dynamicLabelMinWidth='5em'
                    v-model="dataset.ctrl.hideChart"
                    @updateChart="updateDiskInfo"
                >
                </chart-legend>
            </div>
        </div>
    </div>
</template>

<script>
import LineChart from './LineChart.js'
import chartLegend from '@/components/content/resources/chartLegend'
import cm from '../../../js/common'

// 记录上一次更新的 Disk 数据
let preDiskInfo = null;

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
    name: 'DiskChart',
    components: {
        LineChart, chartLegend, 
    },
    // rsc 是从 resources 通过 rsc.startUpdater() 获取的系统资源信息数据
    // rscOp 是 resources 传进来的全局图表 options
    // selectedDisks: 当前需要显示的 disk
    props: ['rsc', 'interval', 'rscOp', 'selectedDisks'],
    data () {
        return {
            disks: [],  
            datacollection: {},
        }
    },
    computed: {
      // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
      options: function(){
        let self = this;
        
        let op = JSON.parse(JSON.stringify(self.rscOp.chartJsOp));
        // op.title.text = "Disk(s) History";
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
    },
    mounted () {
    },
    watch: {
        rsc: function(){
            this.updateDiskInfo();
      },
      selectedDisks: function(){
        this.updateDiskInfo();
      },
    },
    methods: {
        // updateDiskInfo 用来更新 disks 内的数据，也可以用来重绘图像
        updateDiskInfo(){
            if (!preDiskInfo){
                preDiskInfo = this.rsc.Disk
            }
            this.recordDisks();
            this.updateUnit();
            this.fmtSelectedDisks();
            this.updateDatacollection();
            preDiskInfo = this.rsc.Disk;
        },
        recordDisks(){
            let self = this;
            let partitions = self.rsc.Disk.Partitions;

            if (self.disks.length < 1){
                self.disks.push(new Disk("All"));
                for (let i = 0; i < partitions.length; i++){
                    self.disks.push(new Disk(partitions[i].Name));
                }
            }

            let totalWRate = 0;
            let totalRRate = 0;
            // 第一个 disk 是 All, 所以 self.disks[i] 对应的是 self.rsc.Disk.Partitions[i-1]
            for (let i = 1; i < self.disks.length; i++){
                let d = self.disks[i];
                let c = partitions[i-1];

                let p = preDiskInfo.Partitions[i-1];
                let rSector = c.SectorsRead - p.SectorsRead;
                let wSector = c.SectorsWrite - p.SectorsWrite;
                let rRate = rSector * c.LogSector;
                totalRRate += rRate;
                let wRate = wSector * c.LogSector;
                totalWRate += wRate;
                recordDisk(self.disks[i], rRate, wRate, self.points.length);

            }
            recordDisk(self.disks[0], totalRRate, totalWRate, self.points.length);

            
        },
        // 更新全局 unit
        updateUnit(){
            let self = this;
            let max = 0;
            for (let i = 0; i < self.selectedDisks.length; i++){
                let index = self.selectedDisks[i];
                let d = self.disks[index];
                let dMax = Math.max(maxItem(d.rRec), maxItem(d.wRec));
                if (dMax > max){
                    max = dMax;
                }
            }
            preUnit = unit;
            unit = cm.fmtSize.sizeUnit(max);
        },
        fmtSelectedDisks(){
            let self = this;
            for(let i = 0; i < self.selectedDisks.length; i++){
                let index = self.selectedDisks[i];
                fmtDisk(self.disks[index], self.points.length);
            }
        },
        updateDatacollection(){
            let self = this;
            let datasets = [];
            let colorIndex = 0;
            
            for (let i = 0; i < self.selectedDisks.length; i++){
                let d = self.disks[self.selectedDisks[i]];
                // console.log("d.rFmtRec: " + d.rFmtRec);
                // console.log("d.wFmtRec: " + d.wFmtRec);
                let rDataset = {
                    label: d.name + "(r)",
                    dynamiclabel: cm.fmtSize.fmtSize(d.rRec[0], 1),
                    ctrl: d.rCtrl,
                    borderColor: colors[colorIndex++],
                    backgroundColor: ['rgba(0, 0, 0, 0)'],
                    data: d.rFmtRec,
                }
                datasets.push(rDataset);

                let wDataset = {
                    label: d.name + "(w)",
                    dynamiclabel: cm.fmtSize.fmtSize(d.wRec[0], 1),
                    ctrl: d.wCtrl,
                    borderColor: colors[colorIndex++],
                    backgroundColor: ['rgba(0, 0, 0, 0)'],
                    data: d.wFmtRec,
                }
                datasets.push(wDataset);
            }
            self.datacollection = {
                labels: self.points,
                datasets: datasets,
            }
            
        },
    },
    
}

// name: 磁盘名
// rRec：读取速率记录
// wRec：写入速率记录
// rFmtRec：根据单位格式化后的读取记录
// wFmtRec：根据单位格式化后的写入记录
function Disk(name, rRec, wRec, unit, rFmtRec, hideRRec, wFmtRec, hideWRec){
    this.name = name;
    if (!rRec) {
        this.rRec = [];
    } else {
        this.rRec = rRec;
    }

    if (!wRec) {
        this.wRec = [];
    } else {
        this.wRec = wRec;
    }
    this.rFmtRec = [];
    this.wFmtRec = [];

    this.rCtrl = {
        hideChart: hideRRec || false,
    }

    this.wCtrl = {
        hideChart: hideWRec || false,
    }
}

// 仅记录磁盘的读写速率，不更新绘图部分的数值
function recordDisk(disk, rRate, wRate, pointsLen){
    updateElements(disk.rRec, rRate, pointsLen);
    updateElements(disk.wRec, wRate, pointsLen);
}

// 根据 recordDisk 记录的 rRec 和 wRec，更新 disk 的 nit, rFmtRec, wFmtRec
function fmtDisk(disk, pointsLen){
    if (!disk.rCtrl.hideChart) {
        if (preUnit != unit || (disk.rRec.length-1)>(disk.rFmtRec.length)){
            reFmtRec(disk.rRec, disk.rFmtRec, unit);
        } else {
            updateElements(disk.rFmtRec, cm.fmtSize.fmtSizeByUnit(disk.rRec[0], unit).toFixed(2), pointsLen);
        }
    } else {
        disk.rFmtRec = [];
    }

    if (!disk.wCtrl.hideChart){
        if (preUnit != unit || (disk.wRec.length-1)>(disk.wFmtRec.length)){
            reFmtRec(disk.wRec, disk.wFmtRec, unit);
        } else {
            updateElements(disk.wFmtRec, cm.fmtSize.fmtSizeByUnit(disk.wRec[0], unit).toFixed(2), pointsLen);
        }
    } else {
        disk.wFmtRec = [];
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