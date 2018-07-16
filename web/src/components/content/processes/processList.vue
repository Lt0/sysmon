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
                    <td class="text-xs-left" v-show="displayCPU">{{props.item.CPU}}%</td>
                    <td class="text-xs-left" v-show="displayMEM">{{props.item.MEM}}%</td>
                    <td class="text-xs-left" v-show="displayCPUTime">
                        <v-tooltip bottom>
                            <span slot="activator">{{ timeParseSec(props.item.CPUTime) }}</span>
                            <span>{{props.item.CPUTime}} sec</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayTaskCPU">{{ props.item.TaskCPU }}</td>

                    <td class="text-xs-left" v-show="displayRRate">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.RRate, 1)}}</span>
                            <span>{{props.item.RRate}} byte(s)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayWRate">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.WRate, 1)}}</span>
                            <span>{{props.item.WRate}} byte(s)</span>
                        </v-tooltip>
                    </td>

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

                    <td class="text-xs-left" v-show="displayRead">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.Read, 1)}}</span>
                            <span>{{props.item.Read}} byte(s)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayWrite">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.Write, 1)}}</span>
                            <span>{{props.item.Write}} byte(s)</span>
                        </v-tooltip>
                    </td>

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

// 获取 server 信息，当前仅获取 SysInfo
let serverInfo = {};
cm.sysInfo.getSysInfo(serverInfo)

export default {
    name: 'processList',
    components: {
        selection, 
    },
    props: ['info'],
    data() {
        return {
            cm: cm,
            serverInfo: serverInfo,
            windowSize: {
                x: 0,
                y: 0
            },
            latestUpdate: null,
            search: '',
            items: ['Comm', 'CPU', 'MEM', 'CPUTime', 'TaskCPU', 'RRate', 'WRate', 'VmSize', 'VmRSS', 'VmPTE', 'VmSwap', 'Pid', 'Nlwp', 'State', 'Nice', 'Priority', 'User', 'Uid', 'Read', 'Write', 'Cmdline'],  // 所有可显示的项目
            selectedItems: [],  // 实际显示的项目，由 selection 返回
            selectTypes: null,
            
            displayComm: true,
            displayCPU: true,
            displayMEM: true,
            displayCPUTime: false,
            displayTaskCPU: false,
            displayRRate: false,
            displayWRate: false,
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
            displayRead: false,
            displayWrite: false,
            displayCmdline: false,
            headers: [],

            avgJiffies: 0,
            preJiffies: 0,

            prePidJiffies: {},

            secJiffies: 0,  // 单个核心从系统启动以来平均每秒产生的节拍数

            PreIOReadBytes: {},
            PreIOWriteBytes: {},
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
            // console.log("this.selectedItems: " + this.selectedItems);
            this.updateHeaders();
            this.updateDisplay();
        },
        info: function(){
            this.updateJiffies();
            this.formatInfo();

            let time = new Date();
            this.latestUpdate = time.toLocaleTimeString();
        },
    },
    methods: {
        formatInfo() {
            let tmpPrePidJiffies = {};
            let tmpPreIOReadBytes = {};
            let tmpPreIOWriteBytes = {};

            let processes = this.info.Processes;
            for(let i in processes) {
                let p = processes[i];
                // format MEM
                p.MEM = (p.VmRSS/serverInfo.SysInfo.HW.Mem.PhySize*100).toFixed(2);

                // format CPU
                if(!this.prePidJiffies[p.Pid]) {
                    p.CPU = (p.UsedCPU/this.avgJiffies*100).toFixed(2);
                } else {
                    p.CPU = ((p.UsedCPU-this.prePidJiffies[p.Pid])/this.avgJiffies*100).toFixed(2);
                }
                tmpPrePidJiffies[p.Pid] = p.UsedCPU;
                // format CPU end

                // format TaskCPU
                p.TaskCPU = "cpu" + p.TaskCPU;

                // format CPUTime
                // 进程一共消耗的节拍数/单个核心一秒一共能产生的节拍数 = 进程使用单个核心的 CPU 时长（sec）
                p.CPUTime = parseInt(p.UsedCPU/this.secJiffies);

                // format Read/Write
                p.Read = p.IOReadBytes;
                p.Write = p.IOWriteBytes;

                // format IO R/W
                if(!this.PreIOReadBytes[p.Pid]) {
                    p.RRate = p.IOReadBytes;
                } else {
                    p.RRate = p.IOReadBytes - this.PreIOReadBytes[p.Pid];
                }
                tmpPreIOReadBytes[p.Pid] = p.IOReadBytes;

                if(!this.PreIOWriteBytes[p.Pid]) {
                    p.WRate = p.IOWriteBytes;
                } else {
                    p.WRate = p.IOWriteBytes - this.PreIOWriteBytes[p.Pid];
                }
                tmpPreIOWriteBytes[p.Pid] = p.IOWriteBytes;
            }

            this.prePidJiffies = tmpPrePidJiffies;
            this.PreIOReadBytes = tmpPreIOReadBytes;
            this.PreIOWriteBytes = tmpPreIOWriteBytes;
        },
        updateJiffies() {
            let c = this.info.Cores[0];
            let cur = c.User + c.Nice + c.System + c.Idle + c.Iowait + c.Irq + c.Softirq + c.Steal + c.Guest + c.Guest_nice
            this.avgJiffies = (cur - this.preJiffies)/this.info.CoreNum;
            this.preJiffies = cur;

            this.secJiffies = cur/this.info.UpTime.Uptime/this.info.CoreNum;
        },
        updateHeaders(){
            this.headers = [];

            for(let i in this.items) {
                for (let j in this.selectedItems){
                    if(this.items[i] == this.selectedItems[j]) {
                        let t = this.selectedItems[j];
                        // if (this.selectedItems[j] == 'CPU'){
                        //     t += "(%)";
                        // }
                        let hdr = {text: t, value: this.selectedItems[j]};
                        this.headers.push(hdr);
                        break;
                    }
                }
            }
        },
        updateDisplay(){
            this.displayComm = false;
            this.displayCPU = false;
            this.displayMEM = false;
            this.displayCPUTime = false;
            this.displayTaskCPU = false;
            this.displayRRate = false;
            this.displayWRate = false;
            this.displayVmSize = false;
            this.displayVmRSS = false;
            this.displayVmPTE = false;
            this.displayVmSwap = false;
            this.displayPid = false;
            this.displayNlwp = false;
            this.displayState = false;
            this.displayNice = false;
            this.displayPriority = false;
            this.displayUser = false;
            this.displayUid = false;
            this.displayRead = false;
            this.displayWrite = false;
            this.displayCmdline = false;

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
                    case 'RRate':
                        this.displayRRate = true;
                        break;
                    case 'WRate':
                        this.displayWRate = true;
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
                    case 'Read':
                        this.displayRead = true;
                        break;
                    case 'Write':
                        this.displayWrite = true;
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
        timeParseSec(sec) {
            sec = parseInt(sec)
            let h = parseInt(sec/3600).toString();
            if(h.length < 2) {
                h = "0" + h
            }

            sec = sec%3600;
            let m = parseInt(sec/60).toString();
            if(m.length < 2) {
                m = "0" + m
            }

            let s = (sec%60).toString()
            if(s.length < 2) {
                s = "0" + s
            }

            return h + ":" + m + ":" + s
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
