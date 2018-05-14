<template>
    <line-chart :chart-data="datacollection" :options="options" :styles="rscOp.vueChartOp.styles"></line-chart>
</template>

<script>
import LineChart from './LineChart.js'
import cm from '../../../js/common'

// 记录上一次更新的 Net 数据
let preDiskInfo = null;
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
    name: 'RscDisk',
    components: {
        LineChart,
    },
    // rsc 是从 resources 通过 rsc.startUpdater() 获取的系统资源信息数据
    // rscOp 是 resources 传进来的全局图表 options
    props: ['rsc', 'interval', 'rscOp'],
    data () {
        return {
            disks: [],
            selectedDisk: 0,    // 当前显示哪一个 disk，默认为 0, 也即 total，所有 disk 的统计信息
            datacollection: {},
        }
    },
    computed: {
      // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
      options: function(){
        let self = this;
        
        let op = JSON.parse(JSON.stringify(self.rscOp.chartJsOp));
        op.title.text = "Disk(s) History";
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
            if (!preDiskInfo){
                preDiskInfo = this.rsc.Disk
            }
            this.updateDisks();
            this.updateUnit();
            this.updateDatacollection();
            preDiskInfo = this.rsc.Disk;
      },
    },
    methods: {
        // // 更新全局 unit
        updateUnit(){
            unit = this.disks[this.selectedDisk].unit;
        },
        updateDisks(){
            console.log("fill: " + (count += 1));
            let self = this;
            let partitions = self.rsc.Disk.Partitions;

            if (self.disks.length < 1){
                self.disks.push(new Disk("total"));
                for (let i = 0; i < partitions.length; i++){
                    if (partitions[i].IsMajorHD) {
                        self.disks.push(new Disk(partitions[i].Name));
                    }
                }
            }

            let totalWRate = 0;
            let totalRRate = 0;
            // 第一个 disk 是 total，不需要查找
            for (let i = 1; i < self.disks.length; i++){
                
                let d = self.disks[i];
                for (let j = 0; j < partitions.length; j++){
                    let c = partitions[j];
                    if (d.name == c.Name) {
                        let p = preDiskInfo.Partitions[j];
                        let rSector = c.SectorsRead - p.SectorsRead;
                        let wSector = c.SectorsWrite - p.SectorsWrite;
                        let rRate = rSector * c.LogSector;
                        totalRRate += rRate;
                        let wRate = wSector * c.LogSector;
                        totalWRate += wRate;
                        if (i == self.selectedDisk){
                            updateDisk(self.disks[i], rRate, wRate, self.points.length);
                        } else {
                            recordDisk(self.disks[i], rRate, wRate, self.points.length);
                        }
                        
                    }
                }
            }
            if (0 == self.selectedDisk){
                updateDisk(self.disks[0], totalRRate, totalWRate, self.points.length);
            } else {
                recordDisk(self.disks[0], totalRRate, totalWRate, self.points.length);
            }
            console.log("self.disks[0]: " + JSON.stringify(self.disks[0]));
        },
        updateDatacollection(){
            let self = this;
            self.datacollection = {
                labels: self.points,
                datasets: [
                    {
                        label: 'Read',
                        borderColor: ['rgba(3, 169, 244, 1)'],
                        backgroundColor: ['rgba(3, 169, 244, 0)'],
                        data: self.disks[self.selectedDisk].rFmtRec,
                    },{
                        label: 'Write',
                        borderColor: ['rgba(0, 150, 136, 1)'],
                        backgroundColor: ['rgba(3, 169, 244, 0)'],
                        data: self.disks[self.selectedDisk].wFmtRec,
                    },
                ],
            }
            
        },
    },
    
}

let count = 0;

// name: 磁盘名
// rRec：读取速率记录
// wRec：写入速率记录
// rFmtRec：根据单位格式化后的读取记录
// wFmtRec：根据单位格式化后的写入记录
function Disk(name, rRec, wRec, unit, rFmtRec, wFmtRec){
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

    if (!unit) {
        this.unit = "B";
    } else {
        this.unit = unit;
    }

    this.rFmtRec = [];
    this.wFmtRec = [];
}

// 仅记录磁盘的读写速率，不更新绘图部分的数值
function recordDisk(disk, rRate, wRate, pointsLen){
    updateElements(disk.rRec, rRate, pointsLen);
    updateElements(disk.wRec, wRate, pointsLen);
}

// 更新单个 disk 记录
function updateDisk(disk, rRate, wRate, pointsLen){
    updateElements(disk.rRec, rRate, pointsLen);
    updateElements(disk.wRec, wRate, pointsLen);

    let unit = decideUnit(disk);
    if (disk.unit != unit) {
        console.log("change unit from " + disk.unit + " to " + unit);
        reFmtRec(disk.rRec, disk.rFmtRec, unit);
        reFmtRec(disk.wRec, disk.wFmtRec, unit);
        disk.unit = unit;
    } else {
        updateElements(disk.rFmtRec, cm.fmtSize.fmtSizeByUnit(rRate, disk.unit).toFixed(2), pointsLen);
        updateElements(disk.wFmtRec, cm.fmtSize.fmtSizeByUnit(wRate, disk.unit).toFixed(2), pointsLen);
    }
}

function decideUnit(disk){
    let max = Math.max(maxItem(disk.rRec), maxItem(disk.wRec));
    return cm.fmtSize.sizeUnit(max);
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
    fmtRec = [];
    console.log("rec.length: " + rec.length);
    for (let i = 0; i < rec.length; i++){
        fmtRec[i] = cm.fmtSize.fmtSizeByUnit(rec[i], unit).toFixed(2);
        console.log("fmtRec[i]: " + fmtRec[i]);
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