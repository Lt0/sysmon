<template>
    <div>
        <v-container fluid class="container">
            <v-layout wrap align-center>
                    <v-flex v-for="dataset in dataCollection.datasets" :key="dataset.label" xs6 sm3 md3 lg2 >
                        <chart-legend 
                            :borderColor=dataset.borderColor 
                            :backgroundColor=dataset.backgroundColor
                            :label=dataset.label
                            :dynamicLabel=dataset.dynamiclabel
                            dynamicLabelMinWidth='5em'
                            v-model="dataset.ctrl.hideChart"
                            @updateChart="updateChartHandler"
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
    padding: 0px;
}
</style>

