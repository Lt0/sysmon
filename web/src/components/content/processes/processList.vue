<template>
    <div class="process-container">
        <v-card id="process" v-resize="onResize">
            <v-card-title>
                <v-text-field
                    v-model="search"
                    append-icon="search"
                    label="Search"
                    single-line
                    hide-details
                ></v-text-field>
                <v-spacer></v-spacer>
                <selection style='padding-top: 12px;' :items="items" v-model='selectedItems' defaultItemNum=3 />
            </v-card-title>
            
            <v-data-table
                :headers="headers"
                :items="info.Processes"
                :rows-per-page-items="[10,25,50,100,processNum]"
                class="elevation-1"
                must-sort
                :search="search"
                id="process-table"
            >
                <!-- 定制表头 -->
                <template slot="headerCell" slot-scope="props">
                    <v-tooltip bottom>
                        <span slot="activator">
                            <!-- 表头显示内容 -->
                            {{ props.header.text }}
                        </span>
                        <span>
                            <!-- 表头的 tooltip -->
                            {{ props.header.text }}
                        </span>
                    </v-tooltip>
                </template>
                <template slot="items" slot-scope="props">
                    <td class="text-xs-left" v-show="displayComm">{{ props.item.Comm }}</td>
                    <td class="text-xs-left" v-show="displayCPU">{{ "cpu" }}</td>
                    <td class="text-xs-left" v-show="displayMEM">{{ "mem" }}</td>
                    <td class="text-xs-left" v-show="displayCPUTime">{{ "CPUTime" }}</td>
                    <td class="text-xs-left" v-show="displayTaskCPU">{{ props.item.TaskCPU }}</td>
                    <td class="text-xs-left" v-show="displayVmSize">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VmSize, 1)}}</span>
                            <span>{{props.item.VmSize}}</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayVmRSS">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VmRSS, 1)}}</span>
                            <span>{{props.item.VmRSS}}</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayVmPTE">{{ props.item.VmPTE }}</td>
                    <td class="text-xs-left" v-show="displayVmSwap">{{ props.item.VmSwap }}</td>
                    <td class="text-xs-left" v-show="displayPid">{{ props.item.Pid }}</td>
                    <td class="text-xs-left" v-show="displayNlwp">{{ props.item.Nlwp }}</td>
                    <td class="text-xs-left" v-show="displayState">{{ props.item.State }}</td>
                    <td class="text-xs-left" v-show="displayNice">{{ props.item.Nice }}</td>
                    <td class="text-xs-left" v-show="displayPriority">{{ props.item.Priority }}</td>
                    <td class="text-xs-left" v-show="displayUser">{{ props.item.User }}</td>
                    <td class="text-xs-left" v-show="displayUid">{{ props.item.Uid }}</td>
                    <td class="text-xs-left" v-show="displayCmdline">{{ props.item.Cmdline }}</td>

                </template>

                <!-- 没有数据时显示的内容 -->
                <template slot="no-data">
                    <div id="wait-tips">
                        <v-progress-circular indeterminate :width="3" color="primary"></v-progress-circular>
                    </div>
                </template>

                <!-- 搜索没有匹配的结果时显示的内容 -->
                <v-alert slot="no-results" :value="true" color="warning" icon="warning">
                    Your search for "{{ search }}" found no results.
                </v-alert>
            </v-data-table>
        </v-card>

        <div id="updatedTime" class="elevation-5">
            <v-tooltip top>
                <span slot='activator'><v-icon size="1.5em">update</v-icon>&ensp;{{ latestUpdate }}</span>
                <span>Latest Update &ensp;{{ latestUpdate }}</span>
            </v-tooltip>
        </div>
    </div>
</template>

<script>
import selection from '@/components/common/selection'
import cm from '../../../js/common'

export default {
    name: 'processList',
    components: {
        selection, 
    },
    props: ['info'],
    data() {
        return {
            cm: cm,
            windowSize: {
                x: 0,
                y: 0
            },
            latestUpdate: null,
            search: '',
            items: ['Comm', 'CPU', 'MEM', 'CPUTime', 'TaskCPU', 'VmSize', 'VmRSS', 'VmPTE', 'VmSwap', 'Pid', 'Nlwp', 'State', 'Nice', 'Priority', 'User', 'Uid', 'Cmdline'],  // 所有可显示的项目
            selectedItems: [],  // 实际显示的项目
            selectTypes: null,
            
            displayComm: true,
            displayCPU: true,
            displayMEM: true,
            displayCPUTime: false,
            displayTaskCPU: false,
            displayVmSize: false,
            displayVmRSS: false,
            displayVmPTE: false,
            displayVmSwap: false,
            displayPid: false,
            displayNlwp: false,
            displayState: false,
            displayNice: false,
            displayPriority: false,
            displayUser: false,
            displayUid: false,
            displayCmdline: false,
            headers: [],
        }
    },
    computed: {
        processNum: function() {
            if(this.info.Processes) {
                return this.info.Processes.length;
            } else {
                return 0;
            }
        },
    },
    watch: {
        selectedItems: function(){
            this.updateHeaders();
            this.updateDisplay();
        },
        // 根据屏幕大小变化，调整显示的项目，每 105px 宽显示一个项目
        windowSize: function(){
            let num = Math.floor(this.windowSize.x/125);
            this.selectedItems = [];
            for (let i = 0; i < num; i++){
                if (this.items[i]){
                    this.selectedItems.push(this.items[i]);
                } else {
                    break;
                }
                
            }
            console.log("this.selectedItems: " + this.selectedItems);
            this.updateHeaders();
            this.updateDisplay();
        },
        info: function(){
            console.log("cm.sysInfo.hw.mem: ", cm.sysInfo.hw.mem)
            this.callTest();
            let time = new Date();
            this.latestUpdate = time.toLocaleTimeString();
        },
    },
    methods: {
        callTest() {
            cm.sysInfo.setMem();
        },
        updateHeaders(){
            // console.log("updateHeaders: " + this.selectedItems);
            this.headers = [];
            for (let item in this.selectedItems){
                let t = this.selectedItems[item];
                if (this.selectedItems[item] == 'CPU' || this.selectedItems[item] == 'MEM'){
                    t += "(%)";
                }
                let hdr = {text: t, value: this.selectedItems[item]};
                this.headers.push(hdr);
            }
        },
        updateDisplay(){
            displayComm: true;
            displayCPU: true;
            displayMEM: true;
            displayCPUTime: false;
            displayTaskCPU: false;
            displayVmSize: false;
            displayVmRSS: false;
            displayVmPTE: false;
            displayVmSwap: false;
            displayPid: false;
            displayNlwp: false;
            displayState: false;
            displayNice: false;
            displayPriority: false;
            displayUser: false;
            displayUid: false;
            displayCmdline: false;

            for (let i in this.selectedItems){
                switch(this.selectedItems[i]) {
                    case 'Comm':
                        this.displayComm = true;
                        break;
                    case 'CPU':
                        this.displayCPU = true;
                        break;
                    case 'MEM':
                        this.displayMEM = true;
                        break;
                    case 'CPUTime':
                        this.displayCPUTime = true;
                        break;
                    case 'TaskCPU':
                        this.displayTaskCPU = true;
                        break;
                    case 'VmSize':
                        this.displayVmSize = true;
                        break;
                    case 'VmRSS':
                        this.displayVmRSS = true;
                        break;
                    case 'VmPTE':
                        this.displayVmPTE = true;
                        break;
                    case 'VmSwap':
                        this.displayVmSwap = true;
                        break;
                    case 'Pid':
                        this.displayPid = true;
                        break;
                    case 'Nlwp':
                        this.displayNlwp = true;
                        break;
                    case 'State':
                        this.displayState = true;
                        break;
                    case 'Nice':
                        this.displayNice = true;
                        break;
                    case 'Priority':
                        this.displayPriority = true;
                        break;
                    case 'User':
                        this.displayUser = true;
                        break;
                    case 'Uid':
                        this.displayUid = true;
                        break;
                    case 'Cmdline':
                        this.displayCmdline = true;
                        break;
                }
            }
        },
        onResize(){
            this.windowSize = { x: window.innerWidth, y: window.innerHeight }
        },
    },
}
</script>

<style scoped>
.process-container {
    padding: 0 0 3em 0;
}

#updatedTime {
    position: fixed;
    bottom: 0px;
    left: 0px;

    background: #ffffff;

    padding: 0.3em 0.5em 0.3em 0.5em;

    opacity: 0.6;
    border-radius: 3px;

    font-size: 0.8em;
}

#wait-tips {
    padding: 5em 0em;
}
</style>
