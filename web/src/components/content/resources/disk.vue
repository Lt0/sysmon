<template>
    <div class="rsc-chart-panel">
        <div class="rsc-chart-hdr">
            <div class="rsc-chart-hdr-title">
                <h4>Disk(s) History</h4>
            </div>
            <v-dialog v-model="dialog" max-width="500px" >
                <v-btn slot="activator" flat icon color="teal"><v-icon>list</v-icon></v-btn>

                <v-card>
                    <v-card-title>Select Disks</v-card-title>

                    <v-card-text>
                        <v-select 
                            :items="diskNameList" 
                            v-model="selectedDiskName" 
                            multiple 
                            chips
                            deletable-chips
                        >
                        </v-select>
                    </v-card-text>

                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" flat @click.native="dialog = false">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </div>

        <disk-line-chart :rsc='rsc' :rscOp='rscOp' :interval='interval' :selectedDisks='selectedDisks' />
    </div>
</template>

<script>
import DiskLineChart from '@/components/content/resources/diskLineChart'
export default {
    name: 'RscDisk',
    props: ['rsc', 'interval', 'rscOp'],
    components: {
        DiskLineChart, 
    },
    data () {
        return {
            dialog: false,
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

