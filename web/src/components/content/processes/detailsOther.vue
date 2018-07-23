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
                <td>login uid</td>
                <td>{{ this.info.RDLoginUid }}</td>
            </tr>
            <tr>
                <td>session id</td>
                <td>{{ this.info.RDSessionID }}</td>
            </tr>
            <tr>
                <td>setgroup</td>
                <td>{{ this.info.RDSetGroup }}</td>
            </tr>
            <tr>
                <td>cpu set</td>
                <td>{{ this.info.RDCPUSet }}</td>
            </tr>
            <tr>
                <td>coredump filter</td>
                <td>{{ this.info.RDCoredumpFilter }}</td>
            </tr>
            <tr>
                <td>personality</td>
                <td>{{ this.info.RDPersonality }}</td>
            </tr>
            <tr>
                <td>cwd</td>
                <td>{{ this.info.CWD }}</td>
            </tr>
        </table>

        <!-- sched -->
        <div class="info-card">
            <h2>Sched</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ Sched }}</pre>
            </div>
        </div>

        <!-- limits -->
        <div class="info-card">
            <h2>Limits</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <!-- <pre class="rawData">{{ Limits }}</pre> -->
                <div>
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
            </div>
        </div>

        

        <!-- status -->
        <div class="info-card">
            <h2>Status</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ Status }}</pre>
            </div>
        </div>

        <!-- syscall -->
        <div class="info-card">
            <h2>Syscall</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ Syscall }}</pre>
            </div>
        </div>

        <!-- projid map -->
        <div class="info-card">
            <h2>Projid Map</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ ProjidMap }}</pre>
            </div>
        </div>

        <!-- environ -->
        <div class="info-card">
            <h2>Environ</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ Environ }}</pre>
            </div>
        </div>

        <!-- auto group -->
        <div class="info-card">
            <h2>Auto Group</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ AutoGroup }}</pre>
            </div>
        </div>

        <!-- cgroup -->
        <div class="info-card">
            <h2>CGroup</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ CGroup }}</pre>
            </div>
        </div>

        <!-- uid map -->
        <div class="info-card">
            <h2>Uid Map</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ UidMap }}</pre>
            </div>
        </div>

        <!-- gid map -->
        <div class="info-card">
            <h2>Gid Map</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ GidMap }}</pre>
            </div>
        </div>

        <!-- FDs -->
        <div class="info-card">
            <h2>File Descriptors</h2>
            <v-divider></v-divider>
            <br>
            <div v-for="item in FDs" :key=item.id>
                {{ item.Name }} -> {{ item.File }}
            </div>
        </div>

        <!-- map files -->
        <div class="info-card">
            <h2>Map Files</h2>
            <v-divider></v-divider>
            <br>
            <div v-for="item in MapFiles" :key=item.id>
                {{ item.AddrRange }} -> {{ item.File }}
            </div>
        </div>

        <!-- mount info -->
        <div class="info-card">
            <h2>Mount Info</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ MountInfo }}</pre>
            </div>
        </div>

        <!-- mounts -->
        <div class="info-card">
            <h2>Mounts</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ Mounts }}</pre>
            </div>
        </div>

        <!-- mount stats -->
        <div class="info-card">
            <h2>Mount Stats</h2>
            <v-divider></v-divider>
            <br>
            <div>
                <pre class="rawData">{{ MountStats }}</pre>
            </div>
        </div>
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
        Sched: function() {
            return this.info.RDSched || "no Sched to show";
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
    padding: 1em 1em 1em 3em;
}

.rawData {
    font-family: 'Liberation Mono', 'Courier', 'Courier New', '宋体';
    text-align: left;
}

.info-table {
    text-align: left;
    padding: 1em 1em 1em 3em;
}
td {
    padding: 0em 1em 0em 0em;
}

pre {
    text-align: left;
}
</style>

