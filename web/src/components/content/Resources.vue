<template>
    <div>
        <rsc-mem :chartData='memChartData' :options='chartOptions' />
        {{ rsc }}
    </div>
</template>

<script>
import cm from '../../js/common'
import RscMem from '@/components/content/resources/mem'

export default {
    name: 'Resources',
    components: {
        RscMem,
    },
    beforeRouteEnter: (to, from, next) => {
        next(vm => {
            console.log("beforeRouteEnter next");
            cm.rsc.stopUpdater();
            //cm.rsc.startUpdater(vm);
        })
    },
    beforeRouteUpdate: (to, from, next) => {
        console.log("beforeRouteUpdate rsc");
        next();
    },
    beforeRouteLeave: (to, from, next) => {
        console.log("beforeRouteLeave rsc");
        cm.rsc.stopUpdater();
        next();
    },

    data () {
        return {
            rsc: new Object(),
        }
    },
    computed: {
        rscStr: function(){
            return JSON.stringify(this.rsc);
        }
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
