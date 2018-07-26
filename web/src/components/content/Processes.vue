<template>
    <div>
        <process-list :info='infoCtrl.info' @show-process-details="showProcessDetailsHandler($event)" />

        <v-dialog v-model="dialog" lazy scrollable persistent fullscreen>
        <v-card>
            <v-toolbar dark color="teal lighten-1">
                <v-toolbar-title><span>Details of pid {{ detailsCtrl.pid }}</span></v-toolbar-title>
                <v-spacer></v-spacer>
                <v-btn icon dark @click.native="closeProcessDetailsHandler()"><v-icon>close</v-icon></v-btn>
            </v-toolbar>
            <!-- <v-card-title class="headline"> </v-card-title> -->

            <v-card-text>
                <v-tabs centered v-model="tabs">
                    <v-tab>
                        <v-tooltip bottom>
                            <span slot="activator">thread</span>
                            <span>{{tips.processes.detailsTabs.thread}}</span>
                        </v-tooltip>
                    </v-tab>
                    <!-- <v-tab>
                        <v-tooltip bottom>
                            <span slot="activator">limit</span>
                            <span>{{tips.processes.detailsTabs.limit}}</span>
                        </v-tooltip>
                    </v-tab> -->
                    <v-tab>
                        <v-tooltip bottom>
                            <span slot="activator">stack</span>
                            <span>{{tips.processes.detailsTabs.stack}}</span>
                        </v-tooltip>
                    </v-tab>
                    <v-tab>
                        <v-tooltip bottom>
                            <span slot="activator">smaps</span>
                            <span>{{tips.processes.detailsTabs.smaps}}</span>
                        </v-tooltip>
                    </v-tab>
                    <v-tab>
                        <v-tooltip bottom>
                            <span slot="activator">numa</span>
                            <span>{{tips.processes.detailsTabs.numaMaps}}</span>
                        </v-tooltip>
                    </v-tab>
                    <v-tab>
                        <v-tooltip bottom>
                            <span slot="activator">other</span>
                            <span>{{tips.processes.detailsTabs.other}}</span>
                        </v-tooltip>
                    </v-tab>

                    <v-tabs-items v-model="tabs">
                        <v-tab-item>
                            <process-list :info='detailsCtrl.info' @show-process-details="showProcessDetailsHandler($event)" />
                        </v-tab-item>
                        <!-- <v-tab-item>
                            <details-limits :limitsInfo='detailsCtrl.info.Limits'></details-limits>
                        </v-tab-item> -->
                        <v-tab-item>
                            <details-stack :stacksInfo='detailsCtrl.info.Stacks'></details-stack>
                        </v-tab-item>
                        <v-tab-item>
                            <details-smaps :smapsInfo='detailsCtrl.info.Smaps'></details-smaps>
                        </v-tab-item>
                        <v-tab-item>
                            <details-numa-maps :numaMapsInfo='detailsCtrl.info.NumaMaps'></details-numa-maps>
                        </v-tab-item>
                        <v-tab-item>
                            <details-other :info='detailsCtrl.info'></details-other>
                        </v-tab-item>
                    </v-tabs-items>
                </v-tabs>
            </v-card-text>
        </v-card>
        </v-dialog>
    </div>
</template>

<script>
import cm from '../../js/common'
import tips from '../../js/tips'
import processList from '@/components/content/processes/processList'
import detailsLimits from '@/components/content/processes/detailsLimits'
import detailsStack from '@/components/content/processes/detailsStack'
import detailsSmaps from '@/components/content/processes/detailsSmaps'
import detailsNumaMaps from '@/components/content/processes/detailsNumaMaps'
import detailsOther from '@/components/content/processes/detailsOther'

let infoCtrl = {
    type: "all",
    updater: null,
    interval: 1000,         // 更新时间间隔
    info: new Object(),
}
let detailsCtrl = {
    type: "details",
    updater: null,
    interval: 1000,         // 更新时间间隔
    info: new Object(),
}

export default {
    name: 'Processes',
    components: {
        processList, 
        detailsLimits, 
        detailsStack, 
        detailsSmaps, 
        detailsNumaMaps, 
        detailsOther, 
    },
    data () {
        return {
            // 进程数据控制结构，由 processes updater 负责填充
            infoCtrl: infoCtrl,

            // 单个进程的详细信息控制结构，打开 dialog 的时候启动 updater，关闭 dialog 的时候停止 updater
            detailsCtrl: detailsCtrl,

            dialog: false,
            tabs: null,
            tips: tips,
        }
    },
    beforeRouteEnter: (to, from, next) => {
        next(vm => {
            // console.log("processes beforeRouteEnter next");
            cm.process.startUpdater(vm.infoCtrl, vm.infoCtrl.type);
        })
    },
    beforeRouteUpdate: (to, from, next) => {
        // console.log("processes beforeRouteUpdate rsc");
        next();
    },
    beforeRouteLeave: (to, from, next) => {
        // 调用 beforeRouteLeave 的时候无法访问 this 对象，所以使用外部定义的指针
        cm.process.stopUpdater(infoCtrl);
        cm.process.stopUpdater(detailsCtrl);
        next();
    },

    methods: {
        showProcessDetailsHandler(pid) {
            cm.process.stopUpdater(infoCtrl);
            cm.process.stopUpdater(detailsCtrl);

            this.detailsCtrl.pid = pid;
            cm.process.startUpdater(this.detailsCtrl);
            this.dialog = true;
        },
        closeProcessDetailsHandler(){
            this.dialog = false;
            cm.process.stopUpdater(this.detailsCtrl);
            cm.process.startUpdater(infoCtrl, infoCtrl.type);
        }
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
