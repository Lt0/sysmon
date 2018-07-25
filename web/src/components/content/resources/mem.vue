<template>
    <div>
        <chart-hdr title='Memory Usage History' :items='typeNameList' v-model='selectedTypeName' :defaultItemNum='2' dataKey="resources-mem" />
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
            typeNameList: [
                'Memory', 'Swap', 'Buffers', 'Cached', 
                'Active', 'Inactive', 'ActiveAnon', 'InactiveAnon', 'ActiveFile', 'InactiveFile', 'Unevictable', 
                'Mlocked', 'Dirty', 'Writeback', 'AnonPages', 'Mapped', 'Shmem', 
                'Slab', 'SReclaimable', 'SUnreclaim', 
                'KernelStack', 'PageTables', 
                'NfsUnstable', 'Bounce', 'WritebackTmp', 'CommitLimit', 'CommittedAS', 
                'Vmalloc', 
                'HardwareCorrupted', 
                'AnonHugePages', 
                'HugePages', 
                'DirectMap4k', 'DirectMap2M', 'DirectMap1G'],
            selectedTypeName: ['Memory', 'Swap'],
        }
    },
}
</script>
