<template>
    <div>
        <chart-hdr title='Disk(s) History' :items='diskNameList' v-model='selectedDiskName' />
        <disk-line-chart :rsc='rsc' :rscOp='rscOp' :interval='interval' :selectedDisks='selectedDisks' />
    </div>
</template>

<script>
import DiskLineChart from '@/components/content/resources/diskLineChart'
import chartHdr from '@/components/content/resources/chartHdr'
export default {
    name: 'RscDisk',
    props: ['rsc', 'interval', 'rscOp'],
    components: {
        DiskLineChart, chartHdr, 
    },
    data () {
        return {
            diskNameList: ["All"],
            selectedDiskName: ["All"],
            selectedDisks: [0],
        }
    },
    watch: {
        rsc: function(){
            if (this.diskNameList.length <= 1){
                this.updatediskNameList();
            }
        },
        selectedDiskName: function(){
            let self = this;
            self.selectedDisks = [];
            for (let i = 0; i < self.selectedDiskName.length; i++){
                self.selectedDisks.push(self.diskNameList.indexOf(self.selectedDiskName[i]))
            }
        }
    },
    methods: {
        updatediskNameList(){
            let self = this;
            let p = self.rsc.Disk.Partitions;
            for (let i = 0; i < p.length; i++){
                self.diskNameList.push(p[i].Name);
            }
        },
    },
}
</script>

<style scoped>

</style>

