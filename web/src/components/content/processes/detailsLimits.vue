<template>
    <div>
        <v-data-table
            :headers="headers"
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
</template>

<<script>
import tips from '../../../js/tips'

export default {
    props: ['limitsInfo'],
    data() {
        return {
            headers: [
                {text: "Limit", value: "Limit", tips: tips.processes.limits.Limit},
                {text: "Soft Limit", value: "SoftLimit", tips: tips.processes.limits.SoftLimit},
                {text: "Hard Limit", value: "HardLimit", tips: tips.processes.limits.HardLimit},
                {text: "Units", value: "Units", tips: tips.processes.limits.Units},
            ],
        }
    },
    computed: {
        Limits: function() {
            if(this.limitsInfo) {
                return this.limitsInfo.Limits;
            }

            return [];
        },
    },
}
</script>>
