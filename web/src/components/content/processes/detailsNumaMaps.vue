<template>
    <div class="numa-maps-container">
        <v-card id="numa-maps-card" v-resize="onResize">
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
                :items="NumaMaps"
                class="elevation-1"
                must-sort
                :search="search"
                id="process-table"

                :rows-per-page-items='[10,25,50, {text: "ALL", value: -1}]'
                rows-per-page-text = ""
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
                    <td class="text-xs-left file-item" v-if="displayAddr">{{props.item.Addr}}</td>
                    <td class="text-xs-left file-item" v-if="displayPolicy">{{props.item.Policy}}</td>
                    <td class="text-xs-left file-item" v-if="displayFile">{{props.item.File}}</td>
                    <td class="text-xs-left file-item" v-if="displayNodes">{{props.item.Nodes}}</td>
                    <td class="text-xs-left file-item" v-if="displayMapType">{{props.item.MapType}}</td>
                    <td class="text-xs-left" v-if="displayAnon">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Anon }}</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.Anon * (props.item.KernelPageSizeKB || 4), 1)}} ({{props.item.Anon * (props.item.KernelPageSizeKB || 4)}} KB)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-if="displayDirty">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Dirty }}</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.Dirty * (props.item.KernelPageSizeKB || 4), 1)}} ({{props.item.Dirty * (props.item.KernelPageSizeKB || 4)}} KB)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-if="displayMapped">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Mapped }}</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.Mapped * (props.item.KernelPageSizeKB || 4), 1)}} ({{props.item.Mapped * (props.item.KernelPageSizeKB || 4)}} KB)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left file-item" v-if="displayMapMax">{{props.item.MapMax}}</td>
                    <td class="text-xs-left" v-if="displaySwapCache">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.SwapCache }}</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.SwapCache * (props.item.KernelPageSizeKB || 4), 1)}} ({{props.item.SwapCache * (props.item.KernelPageSizeKB || 4)}} KB)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-if="displayActive">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.Active }}</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.Active * (props.item.KernelPageSizeKB || 4), 1)}} ({{props.item.Active * (props.item.KernelPageSizeKB || 4)}} KB)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-if="displayWriteBack">
                        <v-tooltip bottom>
                            <span slot="activator">{{ props.item.WriteBack }}</span>
                            <span>{{cm.fmtSize.fmtKBSize(props.item.WriteBack * (props.item.KernelPageSizeKB || 4), 1)}} ({{props.item.WriteBack * (props.item.KernelPageSizeKB || 4)}} KB)</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left file-item" v-if="displayKernelPageSizeKB">{{props.item.KernelPageSizeKB}}</td>
                    
                </template>

                <!-- 没有数据时显示的内容 -->
                <template slot="no-data">
                    Nothing to show
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
import tips from '../../../js/tips'

import Vuetify from 'vuetify'

export default {
    name: 'detailsNumaMaps',
    props: ['numaMapsInfo'],
    components: {
        selection, 
    },
    data() {
        return {
            cm: cm,
            tips: tips,
            windowSize: {
                x: 0,
                y: 0
            },
            // latestUpdate: null,
            search: '',
            items: ['Addr', 'Policy', 'File', 'Nodes', 'MapType', 'Anon', 'Dirty', 'Mapped', 'MapMax', 'SwapCache', 'Active', 'WriteBack', 'KernelPageSizeKB'],  // 所有可显示的项目
            selectedItems: [],  // 实际显示的项目，由 selection 返回
            selectAddr: null,
            displayPolicy: true,
            displayFile: false,
            displayNodes: true,
            displayMapType: true,
            displayAnon: false,
            displayDirty: false,
            displayMapped: false,
            displayMapMax: false,
            displaySwapCache: false,
            displayActive: false,
            displayWriteBack: false,
            displayKernelPageSizeKB: false,
            headers: [],

            // rowsPerPageItems: [5,10,25,{text: Vuetify.dataIterator.rowsPerPageAll, value: -1}],
        }
    },
    computed: {
        NumaMaps: function() {
            if(this.numaMapsInfo && this.numaMapsInfo.Mappings) {
                let Mappings = this.numaMapsInfo.Mappings;
                // 按要求格式化所有的字符串
                for(let i in Mappings) {
                    let mapping = Mappings[i];

                    // 格式化 Nodes
                    let nodesStr = "";
                    for(let j in mapping.Nodes) {
                        nodesStr += mapping.Nodes[j].Node + ": " + mapping.Nodes[j].NrPages + "; "
                    }
                    mapping.Nodes = nodesStr;
                    // 格式化 Nodes End
                }

                return Mappings;
            }

            return [];
        },
    },
    watch: {
        selectedItems: function(){
            this.updateHeaders();
            this.updateDisplay();
        },
        // 根据屏幕大小变化，调整显示的项目，每 150px 宽显示一个项目
        windowSize: function(){
            let num = Math.floor(this.windowSize.x/150);
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
        // numaMapsInfo: function() {
        //     console.log("numaMapsInfo: ", this.numaMapsInfo)
        // },
    },
    methods: {
        updateHeaders(){
            this.headers = [];

            for(let i in this.items) {
                for (let j in this.selectedItems){
                    if(this.items[i] == this.selectedItems[j]) {
                        let t = this.selectedItems[j];
                        let hdr = {text: t, value: t};
                        // 每一项表头的 tips 都定义在 tips.processes.details.numaMapsHdr 这个对象的同名项中
                        hdr.tips = tips.processes.numaMapsHdr[this.selectedItems[j]];
                        this.headers.push(hdr);
                        break;
                    }
                }
            }
        },
        updateDisplay(){
            this.displayAddr = false;
            this.displayPolicy = false;
            this.displayFile = false;
            this.displayNodes = false;
            this.displayMapType = false;
            this.displayAnon = false;
            this.displayDirty = false;
            this.displayMapped = false;
            this.displayMapMax = false;
            this.displaySwapCache = false;
            this.displayActive = false;
            this.displayWriteBack = false;
            this.displayKernelPageSizeKB = false;
            
            for (let i in this.selectedItems){
                switch(this.selectedItems[i]) {
                    case 'Addr':
                        this.displayAddr = true;
                        break;
                    case 'Policy':
                        this.displayPolicy = true;
                        break;
                    case 'File':
                        this.displayFile = true;
                        break;
                    case 'Nodes':
                        this.displayNodes = true;
                        break;
                    case 'MapType':
                        this.displayMapType = true;
                        break;
                    case 'Anon':
                        this.displayAnon = true;
                        break;
                    case 'Dirty':
                        this.displayDirty = true;
                        break;
                    case 'Mapped':
                        this.displayMapped = true;
                        break;
                    case 'MapMax':
                        this.displayMapMax = true;
                        break;
                    case 'SwapCache':
                        this.displaySwapCache = true;
                        break;
                    case 'Active':
                        this.displayActive = true;
                        break;
                    case 'WriteBack':
                        this.displayWriteBack = true;
                        break;
                    case 'KernelPageSizeKB':
                        this.displayKernelPageSizeKB = true;
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
.numa-maps-container {
    padding: 0 0 3em 0;
}
.file-item {
    word-break: break-all;
}
</style>
