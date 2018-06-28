<template>
    <div class="card">
        <div class="hdr">
            <v-tooltip bottom>
                <span slot="activator">
                    <h2 class="hdr-title">{{storage.Fs}}</h2>
                </span>
                <span>Device: {{storage.Fs}}</span>
            </v-tooltip>
            
        </div>
        <div class="chart">
            <doughnut :chart-data="datacollection" :options="options"></doughnut>
        </div>
        <div class="detail">
            <h3>Details</h3>
            <div class="details-item">Type: {{storage.Type}}</div>
            <div class="details-item">Total: {{storage.Size}}</div>
            <div class="details-item">Used: {{storage.Used}}</div>
            <div class="details-item">Avail: {{storage.Avail}}</div>
            <div class="details-item">Use%: {{storage.UsePercent}}</div>
            <v-tooltip bottom>
                <span slot="activator">
                    <div class="details-item">MountPoint: {{storage.MountPoint}}</div>
                </span>
                <span>MountPoint: {{storage.MountPoint}}</span>
            </v-tooltip>
        </div>
        
    </div>
</template>

<script>
import Doughnut from '../../common/Doughnut.js'
import cm from '../../../js/common'

let tooltipsCallback = {
  label: function(tooltipItem, data){
    // tooltipItem 是一些控制项，包括 datasets 下标之类的
    // data 是绘图时传进去的 data

    // console.log("tooltipItem:", JSON.stringify(tooltipItem))
    // console.log("data:", data.datasets[tooltipItem.datasetIndex].data[tooltipItem.index])
    let size = data.datasets[tooltipItem.datasetIndex].data[tooltipItem.index];
    return " " + cm.fmtSize.fmtKBSize(size*1024*1024, 1);
  },
}

export default {
    name: 'diskCard',
    components: {Doughnut},
    props: ['storage'],
    data () {
        return {
            datacollection: null,
        }
    },
    computed: {
        // options 是要通过 vue-chartjs 传给图表组件的 chartjs 配置
        options: function(){
            let self = this;
            let op = {tooltips: {callbacks: null}};
            op.tooltips.callbacks = tooltipsCallback;
            return op;
        },
    },
    mounted () {
      this.fillData()
    },
    methods: {
        fillData () {
            this.datacollection = {
                labels: ["Used", "Avail"],
                datasets: [
                    {
                        label: 'Total',
                        backgroundColor: ['#36a2eb', '#ffce56'],
                        data: [cm.fmtSize.str2GSize(this.storage.Used), cm.fmtSize.str2GSize(this.storage.Avail)]
                    }
                ]
            }
        },
    }
}
</script>

<style scoped>
    .hdr {
        width: 100%;
    }
    .hdr-title {
        overflow: hidden; 
        white-space: nowrap; 
        text-overflow: ellipsis; 
    }
    .card {
        background-color: #ffffff;
        
        margin: 1em;
        padding: 1em;

        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .chart {
        width: 95%;
        padding: 1em;
    }

    .detail {
        padding: 1em;

        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: flex-start;

        color: #888888;
        font-size: 0.8em;
        overflow: hidden;

        width: 100%;
    }

    .details-item {
        padding: 0.3em 0em 0em 0em;

        overflow: hidden; 
        white-space: nowrap; 
        text-overflow: ellipsis; 
    }
</style>
