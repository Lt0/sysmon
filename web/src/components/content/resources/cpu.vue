

<template>
    <div>
        <chart-hdr title='CPU(s) Usage History' :items='coreList' v-model='selectedCores' defaultItemNum=1 />
        <cpu-line-chart :cpu='cpu' :rscOp='rscOp' :points='points' :selectedCores='selectedCores' />
    </div>
</template>

<script>
import CpuLineChart from '@/components/content/resources/cpuLineChart'
import chartHdr from '@/components/content/resources/chartHdr'
import cm from '../../../js/common'

export default {
    name: 'RscCpu',
    components: {
        CpuLineChart, chartHdr, 
    },
    // cpu: resources 传进来的 rsc.Cpu
    // points: 绘图点数
    // rscOp: 包含主要的 resources 级别的配置
    props: ['cpu', 'points', 'rscOp'],
    data () {
        return {
            coreList: ['Avg'],
            selectedCores: ['Avg'],
        }
    },
    watch: {
        cpu: function(){
            if (this.coreList.length > 1){
                return;
            }

            for (let i in this.cpu.Cores){
                this.coreList.push(this.cpu.Cores[i].Name);
            }
        },
    },
}
</script>
