<template>
    <div>
        <chart-hdr title='Network History' :items='nicNameList' v-model='selectedNicName' defaultItemNum=1 />
        <net-line-chart :net='net' :rscOp='rscOp' :points='points' :selectedNics='selectedNics' />
    </div>
</template>

<script>
import NetLineChart from '@/components/content/resources/netLineChart'
import chartHdr from '@/components/content/resources/chartHdr'
import cm from '../../../js/common'

export default {
    name: 'RscNet',
    components: {
        NetLineChart, chartHdr, 
    },
    // net: resources 传进来的 rsc.Net
    // points: 绘图点数
    // rscOp: 包含主要的 resources 级别的配置
    props: ['net', 'points', 'rscOp'],
    data () {
        return {
            nicNameList: ["Sum"],
            selectedNicName: ["Sum"],
            selectedNics: [0],
        }
    },
    computed: {
    },
    watch: {
        net: function(){
            if (this.nicNameList.length <= 1){
                this.updateNicNameList();
            }
        },
        selectedNicName: function(){
            let self = this;
            self.selectedNics = [];
            for (let i = 0; i < self.selectedNicName.length; i++){
                self.selectedNics.push(self.nicNameList.indexOf(self.selectedNicName[i]))
            }
        }
    },
    methods: {
        updateNicNameList(){
            let self = this;
            let n = self.net.Nics;
            for (let i = 0; i < n.length; i++){
                self.nicNameList.push(n[i].Name);
            }
        },
    }
}
</script>

<style scoped>
</style>