<template>
    <div class="chart-hdr">
            <div class="chart-hdr-title">
                <h3>{{ title }}</h3>
            </div>
            <selection :items="items" :selectedItems='selectionSelectedItems' :defaultItemNum='defaultItemNum' @input="$emit('input', $event)" :dataKey='dataKey' />            
    </div>
</template>

<script>
import selection from '@/components/common/selection'

export default {
    name: 'chartHdr',
    components: {
        selection, 
    },
    props: ['title', 'items', 'selectedItems', 'defaultItemNum', 'dataKey'],
    model: {
        prop: 'selectedItems',
        event: 'input',
    },
    computed: {
        selectionSelectedItems: function() {
            return this.selectedItems
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
  min-width: 8em;
  text-align: left;
}
</style>
