<template>
    <div class="card">
        <div>
            <h2>{{storage.Fs}}</h2>
        </div>
        <div class="chart">
            <doughnut :chart-data="datacollection"></doughnut>
        </div>
        <div class="detail">
            <h3>Details</h3>
            <div class="details-item">Type: {{storage.Type}}</div>
            <div class="details-item">Total: {{storage.Size}}</div>
            <div class="details-item">Used: {{storage.Used}}</div>
            <div class="details-item">Avail: {{storage.Avail}}</div>
            <div class="details-item">Use%: {{storage.UsePercent}}</div>
            <div class="details-item">MountPoint: {{storage.MountPoint}}</div>
        </div>
        
    </div>
</template>

<script>
import Doughnut from '../../common/Doughnut.js'

export default {
    name: 'diskCard',
    components: {Doughnut},
    props: ['storage'],
    data () {
        return {
            datacollection: null,
        }
    },
    mounted () {
      this.fillData()
    },
    methods: {
        fillData () {
            this.datacollection = {
                labels: ["Used", "Avail"],
                datasets: [
                    {
                        label: 'Total',
                        backgroundColor: ['#36a2eb', '#ffce56'],
                        data: [parseInt(this.storage.Used), parseInt(this.storage.Avail)]
                    }
                ]
            }
        },
    }
}
</script>

<style scoped>
    .card {
        background-color: #ffffff;
        /* height: 40em; */
        min-width: 16em;
        max-width: 30em;
        
        margin: 1em;
        padding: 1em;

        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    .chart {
        width: 95%;
        padding: 1em;
    }

    .detail {
        padding: 1em;

        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: flex-start;

        color: #888888;
        font-size: 0.8em;
    }

    .details-item {
        padding: 0.3em 0em 0em 0em;
    }
</style>
