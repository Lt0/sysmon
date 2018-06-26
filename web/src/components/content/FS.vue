<template>
    <div>
        {{rsc.Storage}}
    </div>
</template>

<script>
import cm from '../../js/common'

export default {
    name: 'FS',
    data () {
        return {
            interval: 1000,
        }
    },
    beforeRouteEnter: (to, from, next) => {
        next(vm => {
            //console.log("processes beforeRouteEnter next");
            cm.fs.startUpdater(vm);
        })
    },
    beforeRouteUpdate: (to, from, next) => {
        // console.log("processes beforeRouteUpdate rsc");
        next();
    },
    beforeRouteLeave: (to, from, next) => {
        //console.log("processes beforeRouteLeave");
        cm.fs.stopUpdater();
        next();
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
