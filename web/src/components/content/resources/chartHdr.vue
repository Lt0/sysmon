<template>
    <div class="chart-hdr">
            <div class="chart-hdr-title">
                <h4>{{ title }}</h4>
            </div>
            <v-dialog v-model="dialog" max-width="500px" >
                <v-btn slot="activator" flat icon color="teal"><v-icon>list</v-icon></v-btn>

                <v-card>
                    <v-card-title>Select Disks</v-card-title>

                    <v-card-text>
                        <v-select 
                            :items="items" 
                            :value='selectedItems'
                            @input="updateSelection"
                            multiple 
                            chips
                            deletable-chips
                        >
                        </v-select>
                    </v-card-text>

                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" flat @click.native="dialog = false">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </div>
</template>

<script>
export default {
    name: 'chartHdr',
    props: ['title', 'items', 'selectedItems'],
    model: {
        prop: 'selectedItems',
        event: 'input',
    },
    data () {
        return {
            dialog: false,
        }
    },
    methods: {
        updateSelection(arg){
            this.$emit('input', arg);
        },
    },
}
</script>

<style scoped>
.chart-hdr{
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}
.chart-hdr-title{
  color: #7e858b;
  min-width: 10em;
  text-align: left;
}
</style>
