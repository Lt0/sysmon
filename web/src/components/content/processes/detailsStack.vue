<template>
    <v-card>
        <v-card-title>
            <v-text-field
                v-model="search"
                append-icon="search"
                label="Search"
                single-line
                hide-details
            ></v-text-field>
            <v-spacer></v-spacer>
        </v-card-title>
            

        <v-container fluid grid-list-md>
            <v-data-iterator
                :items="stacksInfo"
                select-all
                content-tag="v-layout"
                row
                wrap
                :search="search"
            >
                <v-flex
                    slot="item"
                    slot-scope="props"
                    xs12
                    sm6
                    md4
                    lg3
                    xl2
                >
                    <v-card>
                        <v-card-title><h4>Pid: {{ props.item.Pid }}</h4></v-card-title>
                        
                        <v-divider></v-divider>
                        <v-list dense>
                            <v-list-tile v-for="item in props.item.Stack.Stacks" :key="item.id">
                                <v-tooltip bottom>
                                    <span slot="activator"><v-list-tile-content class="stack-card"> {{item.Name}} </v-list-tile-content></span>
                                    <span>Address: {{item.Address}}</span>
                                </v-tooltip>
                            </v-list-tile>
                        </v-list>
                    </v-card>
                </v-flex>
            </v-data-iterator>
        </v-container>
    </v-card>
</template>

<script>
import tips from '../../../js/tips'

export default {
    name: "detailsStack",
    props: ['stacksInfo'],
    data() {
        return {
            search: null,
        }
    },
    watch: {
        // 由于 v-data-iterator 的搜索功能默认只能直接搜索其 items 绑定的首层数据，所以将所有会被搜索的 stacks 数据转为字符串(这些字符串不用显示也可以用于搜索)
        stacksInfo: function() {
            for(let i in this.stacksInfo) {
                this.stacksInfo[i].searchString = JSON.stringify(this.stacksInfo[i].Stack.Stacks);
            }
        },
    },
}
</script>

<style scoped>
    .stack-card {
        user-select: all;
        word-break: break-all;
    }
    #search {
        max-width: 30em;
    }
</style>
