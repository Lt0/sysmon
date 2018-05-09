<template>
  <div>
    <line-chart :chart-data="datacollection"></line-chart>
    <button @click="fillData()">Randomize</button>
  </div>
</template>

<script>
import LineChart from './LineChart.js'
import cm from '../../../js/common'

export default {
    name: 'RscMem',
    components: {
        LineChart,
    },
    props: ['rsc'],
    data () {
        return {
            a: null
        }
    },
    computed: {
        datacollection: function () {
            // `this` 指向 vm 实例
            if (!this.rsc.Disk) {
                return null;
            }
            console.log("recompute datacollection");
            return {
                labels: [cm.rsc.interval/1000*1, cm.rsc.interval/1000*2, cm.rsc.interval/1000*3, cm.rsc.interval/1000*4, cm.rsc.interval/1000*5, cm.rsc.interval/1000*6],
                datasets: [
                    {
                    label: 'Data One',
                    backgroundColor: '#a87970',
                    data: [this.getRandomInt(), this.getRandomInt(), this.rsc.Disk.Partitions[0].MajorNum]
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
        }
    },
    mounted () {
        //this.fillData();
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