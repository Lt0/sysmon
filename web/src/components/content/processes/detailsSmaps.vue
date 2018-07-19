<template>
    <div class="smaps-container">
        <v-card id="smaps-card" v-resize="onResize">
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
                :items="Smaps"
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
                    <td class="text-xs-left" v-show="displayFile">{{props.item.File}}</td>
                    <td class="text-xs-left" v-show="displaySize">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Size, 1)}}</span>
                            <span>{{props.item.Size}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayRss">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Rss, 1)}}</span>
                            <span>{{props.item.Rss}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayPss">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Pss, 1)}}</span>
                            <span>{{props.item.Pss}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displaySharedClean">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.SharedClean, 1)}}</span>
                            <span>{{props.item.SharedClean}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displaySharedDirty">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.SharedDirty, 1)}}</span>
                            <span>{{props.item.SharedDirty}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayPrivateClean">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.PrivateClean, 1)}}</span>
                            <span>{{props.item.PrivateClean}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayPrivateDirty">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.PrivateDirty, 1)}}</span>
                            <span>{{props.item.PrivateDirty}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayReferenced">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Referenced, 1)}}</span>
                            <span>{{props.item.Referenced}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayAnonymous">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Anonymous, 1)}}</span>
                            <span>{{props.item.Anonymous}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayAnonHugePages">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.AnonHugePages, 1)}}</span>
                            <span>{{props.item.AnonHugePages}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displaySwap">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Swap, 1)}}</span>
                            <span>{{props.item.Swap}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayKernelPageSize">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.KernelPageSize, 1)}}</span>
                            <span>{{props.item.KernelPageSize}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayMMUPageSize">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.MMUPageSize, 1)}}</span>
                            <span>{{props.item.MMUPageSize}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayLocked">
                        <v-tooltip bottom>
                            <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.Locked, 1)}}</span>
                            <span>{{props.item.Locked}} KB</span>
                        </v-tooltip>
                    </td>
                    <td class="text-xs-left" v-show="displayVmFlags">{{props.item.VmFlags}}</td>
                    <td class="text-xs-left" v-show="displayStartAddr">{{props.item.StartAddr}}</td>
                    <td class="text-xs-left" v-show="displayEndAddr">{{props.item.EndAddr}}</td>
                    <td class="text-xs-left" v-show="displayPerm">{{props.item.Perm}}</td>
                    <td class="text-xs-left" v-show="displayOffset">{{props.item.Offset}}</td>
                    <td class="text-xs-left" v-show="displayDev">{{props.item.Dev}}</td>
                    <td class="text-xs-left" v-show="displayINode">{{props.item.INode}}</td>
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
    </div>
</template>

<script>
import selection from '@/components/common/selection'
import cm from '../../../js/common'
import tips from '../../../js/tips'

export default {
    name: 'detailsSmaps',
    props: ['smapsInfo'],
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
            items: ['File', 'Size', 'Rss', 'Pss', 'SharedClean', 'SharedDirty', 'PrivateClean', 'PrivateDirty', 'Referenced', 'Anonymous', 'AnonHugePages', 'Swap', 'KernelPageSize', 'MMUPageSize', 'Locked', 'VmFlags', 'StartAddr', 'EndAddr', 'Perm', 'Offset', 'Dev', 'INode'],  // 所有可显示的项目
            selectedItems: [],  // 实际显示的项目，由 selection 返回
            selectTypes: null,
            displayFile: true,
            displaySize: true,
            displayRss: true,
            displayPss: false,
            displaySharedClean: false,
            displaySharedDirty: false,
            displayPrivateClean: false,
            displayPrivateDirty: false,
            displayReferenced: false,
            displayAnonymous: false,
            displayAnonHugePages: false,
            displaySwap: false,
            displayKernelPageSize: false,
            displayMMUPageSize: false,
            displayLocked: false,
            displayVmFlags: false,
            displayStartAddr: false,
            displayEndAddr: false,
            displayPerm: false,
            displayOffset: false,
            displayDev: false,
            displayINode: false,
            headers: [],
        }
    },
    computed: {
        Smaps: function() {
            if(this.smapsInfo) {
                return this.smapsInfo.Mappings;
            }

            return [];
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
    },
    methods: {
        updateHeaders(){
            this.headers = [];

            for(let i in this.items) {
                for (let j in this.selectedItems){
                    if(this.items[i] == this.selectedItems[j]) {
                        let t = this.selectedItems[j];
                        let hdr = {text: t, value: t};
                        // 每一项表头的 tips 都定义在 tips.processes.detailsSmapsHdr 这个对象的同名项中
                        hdr.tips = tips.processes.smapsHdr[this.selectedItems[j]];
                        this.headers.push(hdr);
                        break;
                    }
                }
            }
        },
        updateDisplay(){
            this.displayFile = false;
            this.displaySize = false;
            this.displayRss = false;
            this.displayPss = false;
            this.displaySharedClean = false;
            this.displaySharedDirty = false;
            this.displayPrivateClean = false;
            this.displayPrivateDirty = false;
            this.displayReferenced = false;
            this.displayAnonymous = false;
            this.displayAnonHugePages = false;
            this.displaySwap = false;
            this.displayKernelPageSize = false;
            this.displayMMUPageSize = false;
            this.displayLocked = false;
            this.displayVmFlags = false;
            this.displayStartAddr = false;
            this.displayEndAddr = false;
            this.displayPerm = false;
            this.displayOffset = false;
            this.displayDev = false;
            this.displayINode = false;
            
            for (let i in this.selectedItems){
                switch(this.selectedItems[i]) {
                    case 'File':
                        this.displayFile = true;
                        break;
                    case 'Size':
                        this.displaySize = true;
                        break;
                    case 'Rss':
                        this.displayRss = true;
                        break;
                    case 'Pss':
                        this.displayPss = true;
                        break;
                    case 'SharedClean':
                        this.displaySharedClean = true;
                        break;
                    case 'SharedDirty':
                        this.displaySharedDirty = true;
                        break;
                    case 'PrivateClean':
                        this.displayPrivateClean = true;
                        break;
                    case 'PrivateDirty':
                        this.displayPrivateDirty = true;
                        break;
                    case 'Referenced':
                        this.displayReferenced = true;
                        break;
                    case 'Anonymous':
                        this.displayAnonymous = true;
                        break;
                    case 'AnonHugePages':
                        this.displayAnonHugePages = true;
                        break;
                    case 'Swap':
                        this.displaySwap = true;
                        break;
                    case 'KernelPageSize':
                        this.displayKernelPageSize = true;
                        break;
                    case 'MMUPageSize':
                        this.displayMMUPageSize = true;
                        break;
                    case 'Locked':
                        this.displayLocked = true;
                        break;
                    case 'VmFlags':
                        this.displayVmFlags = true;
                        break;
                    case 'StartAddr':
                        this.displayStartAddr = true;
                        break;
                    case 'EndAddr':
                        this.displayEndAddr = true;
                        break;
                    case 'Perm':
                        this.displayPerm = true;
                        break;
                    case 'Offset':
                        this.displayOffset = true;
                        break;
                    case 'Dev':
                        this.displayDev = true;
                        break;
                    case 'INode':
                        this.displayINode = true;
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
.smaps-container {
    padding: 0 0 3em 0;
}

#wait-tips {
    padding: 5em 0em;
}
</style>
