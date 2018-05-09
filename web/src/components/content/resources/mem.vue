<template>
  <div>
    <line-chart :chart-data="datacollection"></line-chart>
    <button @click="fillData()">Randomize</button>
  </div>
</template>

<script>
import LineChart from './LineChart.js'

export default {
    name: 'RscMem',
    components: {
        LineChart,
    },
    data () {
        return {
            datacollection: null
        }
    },
    mounted () {
        this.fillData();
    },
    methods: {
      fillData () {
        this.datacollection = {
          labels: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt(),],
          datasets: [
            {
              label: 'Data One',
              backgroundColor: '#a87970',
              data: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt()]
            }, {
              label: 'Data One',
              backgroundColor: '#f87979',
              data: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt()]
            }, {
              label: 'Data Third',
              backgroundColor: '#587979',
              data: [this.getRandomInt(), this.getRandomInt(), this.getRandomInt()]
            }
          ]
        }
      },
      getRandomInt () {
        return Math.floor(Math.random() * (50 - 5 + 1)) + 5
      }
    },
}
</script>

<style scoped>

</style>