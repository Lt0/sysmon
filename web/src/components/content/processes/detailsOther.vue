<template>
    <div> 
        <table class="info-table">
            <tr>
                <th></th>
                <th></th>
            </tr>
            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">oom_adj</span>
                        <pre>{{ tips.processes.other.oomAdj }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.OOMAdj }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">oom_score</span>
                        <pre>{{ tips.processes.other.oomScore }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.OOMScore }}</td>
            </tr>
            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">oom_score_adj</span>
                        <pre>{{ tips.processes.other.oomScoreAdj }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.OOMScoreAdj }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">login uid</span>
                        <pre>{{ tips.processes.other.loginUid }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.RDLoginUid }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">session id</span>
                        <pre>{{ tips.processes.other.sessionID }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.RDSessionID }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">setgroup</span>
                        <pre>{{ tips.processes.other.setGroup }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.RDSetGroup }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">cpu set</span>
                        <pre>{{ tips.processes.other.cpuSet }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.RDCPUSet }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">coredump filter</span>
                        <pre>{{ tips.processes.other.coredumpFilter }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.RDCoredumpFilter }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">personality</span>
                        <pre>{{ tips.processes.other.personality }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.RDPersonality }}</td>
            </tr>

            <tr>
                <td>
                    <v-tooltip bottom>
                        <span slot="activator">cwd</span>
                        <pre>{{ tips.processes.other.cwd }}</pre>
                    </v-tooltip>
                </td>
                <td>{{ this.info.CWD }}</td>
            </tr>

        </table>

        <v-expansion-panel expand>
            <!-- sched -->
            <v-expansion-panel-content lazy>
            <div slot="header">
                <v-tooltip bottom>
                    <h5 slot="activator">Sched</h5>
                    <pre>{{ tips.processes.other.schedHdr }}</pre>
                </v-tooltip>
            </div>
            <div class="info-card">
                <table>
                    <tr v-for="item in Sched.Items" :key=item.id>
                        <td>
                            <v-tooltip bottom>
                                <span slot="activator">{{ item.Name }}</span>
                                <pre>{{ tips.processes.other.sched[item.Name] || "no tips" }}</pre>
                            </v-tooltip>
                        </td>
                        <td>{{ item.Val }}</td>
                    </tr>
                </table>
                <br>
                <pre>{{ Sched.Other }}</pre>
            </div>
        </v-expansion-panel-content>

        <!-- limits -->
        <v-expansion-panel-content lazy>
            <div slot="header">
                <v-tooltip bottom>
                    <h5 slot="activator">Limits</h5>
                    <pre>{{ tips.processes.other.limitsHdr }}</pre>
                </v-tooltip>
            </div>
            <div class="info-card">
                <v-data-table
                    :headers="limitsHeaders"
                    :items="Limits"
                    hide-actions
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
                        <td class="text-xs-left">{{props.item.Limit}}</td>
                        <td class="text-xs-left">{{props.item.SoftLimit}}</td>
                        <td class="text-xs-left">{{props.item.HardLimit}}</td>
                        <td class="text-xs-left">{{props.item.Units}}</td>
                    </template>
                    <!-- 没有数据时显示的内容 -->
                    <template slot="no-data">
                        no limits info
                    </template>
                </v-data-table>
            </div>
            </v-expansion-panel-content>

            <!-- status -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Status</h5>
                        <pre>{{ tips.processes.other.statusHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ Status }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- syscall -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Syscall</h5>
                        <pre>{{ tips.processes.other.syscallHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ Syscall }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- environ -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Environ</h5>
                        <pre>{{ tips.processes.other.environHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ Environ }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- FDs -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">File Descriptors</h5>
                        <pre>{{ tips.processes.other.fdHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <div v-for="item in FDs" :key=item.id>
                        {{ item.Name }} -> {{ item.File }}
                    </div>
                </div>
            </v-expansion-panel-content>

            <!-- map files -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Map Files</h5>
                        <pre>{{ tips.processes.other.mapFilesHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <div v-for="item in MapFiles" :key=item.id>
                        {{ item.AddrRange }} -> {{ item.File }}
                    </div>
                </div>
            </v-expansion-panel-content>

            <!-- auto group -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Auto Group</h5>
                        <pre>{{ tips.processes.other.autoGroupHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ AutoGroup }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- cgroup -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">CGroup</h5>
                        <pre>{{ tips.processes.other.cgroupHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ CGroup }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- uid map -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Uid Map</h5>
                        <pre>{{ tips.processes.other.uidMapHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ UidMap }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- gid map -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Gid Map</h5>
                        <pre>{{ tips.processes.other.gidMapHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ GidMap }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- projid map -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Projid Map</h5>
                        <pre>{{ tips.processes.other.projidMapHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ ProjidMap }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- mount info -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Mount Info</h5>
                        <pre>{{ tips.processes.other.mountInfoHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ MountInfo }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- mounts -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Mounts</h5>
                        <pre>{{ tips.processes.other.mountsHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ Mounts }}</pre>
                </div>
            </v-expansion-panel-content>

            <!-- mount stats -->
            <v-expansion-panel-content lazy>
                <div slot="header">
                    <v-tooltip bottom>
                        <h5 slot="activator">Mount Stats</h5>
                        <pre>{{ tips.processes.other.mountStatsHdr }}</pre>
                    </v-tooltip>
                </div>
                <div class="info-card">
                    <pre class="rawData">{{ MountStats }}</pre>
                </div>
            </v-expansion-panel-content>
        </v-expansion-panel>
        
        
    </div>
</template>

<script>
import cm from '../../../js/common'
import tips from '../../../js/tips'

export default {
    name: "detailsOther",
    props: ['info'],
    data() {
        return {
            tips: tips,
            limitsHeaders: [
                {text: "Limit", value: "Limit", tips: tips.processes.limits.Limit},
                {text: "Soft Limit", value: "SoftLimit", tips: tips.processes.limits.SoftLimit},
                {text: "Hard Limit", value: "HardLimit", tips: tips.processes.limits.HardLimit},
                {text: "Units", value: "Units", tips: tips.processes.limits.Units},
            ],
        }
    },
    computed: {
        // Sched: function() {
        //     return this.info.RDSched || "no Sched to show";
        // },
        Sched: function() {
            return this.info.Sched || [];
        },
        Limits: function() {
            if(this.info.Limits) {
                return this.info.Limits.Limits;
            }
            return [];
        },
        Status: function() {
            return this.info.RDStatus || "no status to show";
        },
        Syscall: function() {
            return this.info.RDSyscall || "no syscall to show";
        },
        ProjidMap: function() {
            return this.info.RDProjidMap || "no projid map filter to show";
        },
        Environ: function() {
            return this.info.RDEnviron || "no environ to show";
        },
        AutoGroup: function() {
            return this.info.RDAutoGroup || "no auto group to show";
        },
        CGroup: function() {
            return this.info.RDCGroup || "no cgroup to show";
        },
        UidMap: function() {
            return this.info.RDUidMap || "no uid map to show";
        },
        GidMap: function() {
            return this.info.RDGidMap || "no gid map to show";
        },
        FDs: function() {
            if(this.info.FD) {
                return this.info.FD.FDs;
            }
            return [];
        },
        MapFiles: function() {
            if(this.info.MapFiles) {
                return this.info.MapFiles.Mappings;
            }
            return [];
        },
        MountInfo: function() {
            return this.info.RDMountInfo || "no mount info to show";
        },
        Mounts: function() {
            return this.info.RDMounts || "no mounts to show";
        },
        MountStats: function() {
            return this.info.RDMountStats || "no mount stats to show";
        },
    }
}
</script>

<style scoped>
.info-card {
    text-align: left;
    padding: 1em 1em 1em 1em;
}

.rawData {
    /* font-family: 'Liberation Mono', 'Courier', 'Courier New', '宋体'; */
    text-align: left;
}

.info-table {
    text-align: left;
    padding: 1em 1em 1em 1em;
}
td {
    padding: 0em 1em 0em 0em;
}

pre {
    text-align: left;
}

#sched-content {
    display: flex;
}
.sched-tips {
    color: grey;
}
</style>

