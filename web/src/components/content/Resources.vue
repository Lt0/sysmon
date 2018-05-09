<template>
    <div>
        {{ rsc }}
    </div>
</template>

<script>
import cm from '../../js/common'
export default {
    name: 'Resources',
    beforeRouteEnter: (to, from, next) => {
        next(vm => {
            console.log("beforeRouteEnter next");
            cm.rsc.stopUpdater();
            cm.rsc.startUpdater(vm);
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
