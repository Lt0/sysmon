<template>
    <div>
        <chart-hdr title='Memory Usage History' :items='typeNameList' v-model='selectedTypeName' />
        <mem-line-chart :mem='mem' :rscOp='rscOp' :points='points' :selectedTypeName='selectedTypeName' />
    </div>
</template>

<script>
import MemLineChart from '@/components/content/resources/memLineChart'
import chartHdr from '@/components/content/resources/chartHdr'
import cm from '../../../js/common'

export default {
    name: 'RscMem',
    components: {
        MemLineChart, chartHdr, 
    },
    // mem: resources 传进来的 rsc.Mem
    // points: 绘图点数
    // rscOp: 包含主要的 resources 级别的配置
    props: ['mem', 'points', 'rscOp'],
    data () {
        return {
            typeNameList: ['Memory', 'Swap', 'HugePages'],
            selectedTypeName: ['Memory', 'Swap'],
        }
    },
}
</script>
