<template>
    <v-card id="process">
        <v-card-title>
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
                <td class="text-xs-left">{{ props.item.User }}</td>
                <td class="text-xs-left">{{ props.item.Cmd }}</td>
                <td class="text-xs-left">{{ props.item.CPU}}</td>
                <td class="text-xs-left">{{ props.item.CPUTime }}</td>
                <td class="text-xs-left">{{ props.item.MEM }}</td>
                <td class="text-xs-left">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.SZ, 2)}}</span>
                        <span>{{props.item.SZ}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.RSS, 2)}}</span>
                        <span>{{props.item.RSS}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.DRS, 2)}}</span>
                        <span>{{props.item.DRS}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.TRS, 2)}}</span>
                        <span>{{props.item.TRS}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left">
                    <v-tooltip bottom>
                        <span slot="activator">{{cm.fmtSize.fmtKBSize(props.item.VSZ, 2)}}</span>
                        <span>{{props.item.VSZ}}KB</span>
                    </v-tooltip>
                </td>
                <td class="text-xs-left">{{ props.item.Pid }}</td>
                <td class="text-xs-left">{{ props.item.Nlwp }}</td>
                <td class="text-xs-left">{{ props.item.State }}</td>
                <td class="text-xs-left">{{ props.item.Nice }}</td>
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
</template>

<script>
import cm from '../../../js/common'

export default {
    name: 'processList',
    props: ['info'],
    data() {
        return {
            cm: cm,
            search: '',
            headers: [
                { text: 'User', value: 'User' },
                { text: 'Name', value: 'Cmd' },
                { text: 'CPU', value: 'CPU' },
                { text: 'CPUTime', value: 'CPUTime' },
                { text: 'MEM(%)', value: 'MEM' },
                { text: 'SZ', value: 'SZ' },
                { text: 'RSS', value: 'RSS' },
                { text: 'DRS', value: 'DRS' },
                { text: 'TRS', value: 'TRS' },
                { text: 'VSZ', value: 'VSZ' },
                { text: 'Pid', value: 'Pid' },
                { text: 'Nlwp', value: 'Nlwp' },
                { text: 'State', value: 'State' },
                { text: 'Nice', value: 'Nice' },
            ],
        }
    }
}
</script>

<style scoped>
</style>
