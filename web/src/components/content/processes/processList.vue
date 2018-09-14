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
                :rows-per-page-items='[10,25,50, {text: "ALL", value: -1}]'
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
                            <!-- 表头的 tooltip, updateHeaders 时填充 -->
                            {{ props.header.tips }}
                        </span>
                    </v-tooltip>
                </template>
                <template slot="items" slot-scope="props">
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayComm">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Comm }}</span>
                            <span>{{props.item.Cmdline || props.item.Comm}}</span>
                        </v-tooltip>
                    </td>

                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayCPU">{{props.item.CPU}}%</td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayMEM">
                        <v-tooltip bottom>
                            <span slot="activator">{{props.item.MEM}}%</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.VmRSS, 1)}}</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayCPUTime">
                        <v-tooltip bottom>
                            <span slot="activator">{{ timeParseSec(props.item.CPUTime) }}</span>
                            <span>{{props.item.CPUTime}} sec</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayTaskCPU">{{ props.item.TaskCPU }}</td>

                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayRRate">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.RRate, 1)}}</span>
                            <span>{{props.item.RRate}} byte(s)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayWRate">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.WRate, 1)}}</span>
                            <span>{{props.item.WRate}} byte(s)</span>
                        </v-tooltip>
                    </td>

                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayVmSize">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VmSize, 1)}}</span>
                            <span>{{props.item.VmSize}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayVmRSS">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VmRSS, 1)}}</span>
                            <span>{{props.item.VmRSS}} KB</span>
                        </v-tooltip>
                    </td>

                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayPid">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Pid }}</span>
                            <span>线程 ID(s): {{ props.item.Task }}</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayThreads">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Threads }}</span>
                            <span>线程 ID(s): {{ props.item.Task }}</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayState">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.State }}</span>
                            <span>{{ tips.processes.state[props.item.State] }}</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayNice">{{ props.item.Nice }}</td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayPriority">{{ props.item.Priority }}</td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayUser">{{ props.item.User }}</td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayUid">{{ props.item.Uid }}</td>

                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayVmPTE">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VmPTE, 1)}}</span>
                            <span>{{props.item.VmPTE}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayVmSwap">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VmSwap, 1)}}</span>
                            <span>{{props.item.VmSwap}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayRead">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.Read, 1)}}</span>
                            <span>{{props.item.Read}} byte(s)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayWrite">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtSize(props.item.Write, 1)}}</span>
                            <span>{{props.item.Write}} byte(s)</span>
                        </v-tooltip>
                    </td>

                    <td class="text-xs-left" @dblclick.stop="clickItem(props.item.Pid)" v-if="displayCmdline">{{ props.item.Cmdline }}</td>

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
                <span slot='activator'><icon>update</icon>&ensp;{{ latestUpdate }}</span>
                <span>Latest Update &ensp;{{ latestUpdate }}</span>
            </v-tooltip>
        </div>
    </div>
</template>

<script>
import selection from '@/components/common/selection'
import cm from '../../../js/common'
import tips from '../../../js/tips'

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
            tips: tips,
            serverInfo: serverInfo,
            windowSize: {
                x: 0,
                y: 0
            },
            latestUpdate: null,
            search: '',
            items: ['Comm', 'CPU', 'MEM', 'CPUTime', 'TaskCPU', 'RRate', 'WRate', 'VmSize', 'VmRSS', 'Pid', 'Threads', 'State', 'Nice', 'Priority', 'User', 'Uid', 'VmPTE', 'VmSwap', 'Read', 'Write', 'Cmdline'],  // 所有可显示的项目
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
            displayPid: false,
            displayThreads: false,
            displayState: false,
            displayNice: false,
            displayPriority: false,
            displayUser: false,
            displayUid: false,
            displayVmPTE: false,
            displayVmSwap: false,
            displayRead: false,
            displayWrite: false,
            displayCmdline: false,
            headers: [],

            USER_HZ: 100, // 单个核心从系统启动以来平均每秒产生的节拍数
            preUsedCPU: {},
            preTimeStamp: 0,
            PreIOReadBytes: {},
            PreIOWriteBytes: {},
        }
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
            this.formatInfo();

            let time = new Date();
            this.latestUpdate = time.toLocaleTimeString();
        },
    },
    methods: {
        formatInfo() {
            let tmpPreUsedCPU = {};
            let tmpPreIOReadBytes = {};
            let tmpPreIOWriteBytes = {};

            let timeOffset = (this.info.TimeStamp - this.preTimeStamp)/1000000000;
            let processes = this.info.Processes;
            for(let i in processes) {
                let p = processes[i];
                // format MEM
                p.MEM = (p.VmRSS/serverInfo.SysInfo.HW.Mem.PhySize*100).toFixed(2);

                // format CPU
                if(!this.preUsedCPU[p.Pid]) {
                    p.CPU = "0.00";
                } else {
                    // p.CPU = ((p.UsedCPU-this.preUsedCPU[p.Pid])/this.USER_HZ*100).toFixed(2);
                    p.CPU = ((p.UsedCPU-this.preUsedCPU[p.Pid])/timeOffset).toFixed(2);
                }
                tmpPreUsedCPU[p.Pid] = p.UsedCPU;
                // format CPU end

                // format TaskCPU
                p.TaskCPU = p.TaskCPU.toString();
                if(p.TaskCPU.substr(0, 1) != 'c') {
                    p.TaskCPU = "cpu" + p.TaskCPU;
                }

                // format CPUTime
                // 进程一共消耗的节拍数/单个核心一秒一共能产生的节拍数（USER_HZ, x86 架构下为 100） = 进程使用单个核心的 CPU 时长（sec）
                p.CPUTime = p.UsedCPU/this.USER_HZ;

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

            this.preUsedCPU = tmpPreUsedCPU;
            // console.log("this.preStartTime: " + this.preStartTime)
            this.preTimeStamp = this.info.TimeStamp;
            this.PreIOReadBytes = tmpPreIOReadBytes;
            this.PreIOWriteBytes = tmpPreIOWriteBytes;
        },
        updateHeaders(){
            this.headers = [];

            for(let i in this.items) {
                for (let j in this.selectedItems){
                    if(this.items[i] == this.selectedItems[j]) {
                        let t = this.selectedItems[j];
                        let hdr = {text: t, value: t};
                        // 每一项表头的 tips 都定义在 tips.processes.hdr 这个对象的同名项中
                        hdr.tips = tips.processes.hdr[this.selectedItems[j]];
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
            this.displayPid = false;
            this.displayThreads = false;
            this.displayState = false;
            this.displayNice = false;
            this.displayPriority = false;
            this.displayUser = false;
            this.displayUid = false;
            this.displayVmPTE = false;
            this.displayVmSwap = false;
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
                    case 'Pid':
                        this.displayPid = true;
                        break;
                    case 'Threads':
                        this.displayThreads = true;
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
                    case 'VmPTE':
                        this.displayVmPTE = true;
                        break;
                    case 'VmSwap':
                        this.displayVmSwap = true;
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
            let h = parseInt(sec/3600).toString();
            if(h.length < 2) {
                h = "0" + h
            }

            sec = sec%3600;
            let m = parseInt(sec/60).toString();
            if(m.length < 2) {
                m = "0" + m
            }

            let s = (sec%60).toFixed(2).toString()
            if(s.length < 5) {
                s = "0" + s
            }

            return h + ":" + m + ":" + s
        },
        clickItem(pid){
            console.log("click item pid: ", pid);
            this.$emit('show-process-details', pid);
        }
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
