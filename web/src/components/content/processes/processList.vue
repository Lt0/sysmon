<template>
    <div class="process-container">
    <v-card id="process" v-resize="onResize">
        <v-card-title>
            <selection style='padding-top: 12px;' :items="items" v-model='selectedItems' defaultItemNum=3 />

            <v-spacer></v-spacer>
            
            <v-text-field
                v-model="search"
                append-icon="search"
                label="Search"
                single-line
                hide-details
            ></v-text-field>
        </v-card-title>
        
        <v-data-table
            :headers="headers"
            :items="info.Processes"
            hide-actions
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
                <td class="text-xs-left" v-show="displayUser">{{ props.item.User }}</td>
                <td class="text-xs-left" v-show="displayCmd">{{ props.item.Cmd }}</td>
                <td class="text-xs-left" v-show="displayCPU">{{ props.item.CPU}}</td>
                <td class="text-xs-left" v-show="displayCPUTime">{{ props.item.CPUTime }}</td>
                <td class="text-xs-left" v-show="displayMEM">{{ props.item.MEM }}</td>
                <td class="text-xs-left" v-show="displaySZ">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.SZ, 2)}}</span>
                        <span>{{props.item.SZ}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left" v-show="displayRSS">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.RSS, 2)}}</span>
                        <span>{{props.item.RSS}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left" v-show="displayDRS">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.DRS, 2)}}</span>
                        <span>{{props.item.DRS}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left" v-show="displayTRS">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.TRS, 2)}}</span>
                        <span>{{props.item.TRS}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left" v-show="displayVSZ">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VSZ, 2)}}</span>
                        <span>{{props.item.VSZ}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left" v-show="displayPid">{{ props.item.Pid }}</td>
                <td class="text-xs-left" v-show="displayNlwp">{{ props.item.Nlwp }}</td>
                <td class="text-xs-left" v-show="displayState">{{ props.item.State }}</td>
                <td class="text-xs-left" v-show="displayNice">{{ props.item.Nice }}</td>
            </template>

            <!-- 没有数据时显示的内容 -->
            <template slot="no-data">
                <v-alert :value="true" color="info" icon="warning">
                    Sorry, nothing to display here :(
                </v-alert>
            </template>

            <!-- 搜索没有匹配的结果时显示的内容 -->
            <v-alert slot="no-results" :value="true" color="warning" icon="warning">
                Your search for "{{ search }}" found no results.
            </v-alert>
        </v-data-table>
    </v-card>
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
            search: '',
            items: ['Cmd', 'CPU', 'MEM', 'User', 'CPUTime', 'SZ', 'RSS', 'DRS', 'TRS', 'VSZ', 'Pid', 'Nlwp', 'State', 'Nice'],
            selectedItems: [],
            selectTypes: null,
            displayUser: false,
            displayCmd: true,
            displayCPU: true,
            displayCPUTime: false,
            displayMEM: true,
            displaySZ: false,
            displayRSS: false,
            displayDRS: false,
            displayTRS: false,
            displayVSZ: false,
            displayPid: false,
            displayNlwp: false,
            displayState: false,
            displayNice: false,
            headers: [],
        }
    },
    watch: {
        selectedItems: function(){
            this.updateHeaders();
            this.updateDisplay();
        },
        // 根据屏幕大小变化，调整显示的项目，每 135px 宽显示一个项目
        windowSize: function(){
            let num = Math.floor(this.windowSize.x/135);
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
    },
    methods: {
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
            this.displayUser = false;
            this.displayCmd = false;
            this.displayCPU = false;
            this.displayCPUTime = false;
            this.displayMEM = false;
            this.displaySZ = false;
            this.displayRSS = false;
            this.displayDRS = false;
            this.displayTRS = false;
            this.displayVSZ = false;
            this.displayPid = false;
            this.displayNlwp = false;
            this.displayState = false;
            this.displayNice = false;

            for (let i in this.selectedItems){
                switch(this.selectedItems[i]) {
                    case 'User':
                        this.displayUser = true;
                        break;
                    case 'Cmd':
                        this.displayCmd = true;
                        break;
                    case 'CPU':
                        this.displayCPU = true;
                        break;
                    case 'CPUTime':
                        this.displayCPUTime = true;
                        break;
                    case 'MEM':
                        this.displayMEM = true;
                        break;
                    case 'SZ':
                        this.displaySZ = true;
                        break;
                    case 'RSS':
                        this.displayRSS = true;
                        break;
                    case 'DRS':
                        this.displayDRS = true;
                        break;
                    case 'TRS':
                        this.displayTRS = true;
                        break;
                    case 'VSZ':
                        this.displayVSZ = true;
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
</style>
