<template>
    <div>
        <v-container fluid class="container">
            <v-layout wrap align-center>
                    <v-flex v-for="dataset in dataCollection.datasets" :key="dataset.label" xs12 sm6 md6 lg6 xl6>
                        <chart-legend 
                            :borderColor=dataset.borderColor 
                            :backgroundColor=dataset.backgroundColor
                            :label=dataset.label
                            :dynamicLabel=dataset.dynamiclabel
                            :dynamiclabelPercent=dataset.dynamiclabelPercent
                            dynamicLabelMinWidth='5em'
                            v-model="dataset.ctrl.hideChart"
                            @updateChart="updateChartHandler"
                            :tips="dataset.tips"
                        >
                        </chart-legend>
                    </v-flex>
            </v-layout>
        </v-container>
    </div>
</template>

<script>
import chartLegend from '@/components/content/resources/chartLegend'

export default {
    name: 'chartLegendBar',
    components: {
        chartLegend
    },
    props:['dataCollection'],
    methods: {
        updateChartHandler(arg){
            this.$emit('updateChart', arg);
        },
    },
}
</script>

<style scoped>
.container{
    padding: 0 1em 0 1em;
}
</style>

