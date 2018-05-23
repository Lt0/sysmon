<template>
    <div>
        <process-list :info='info' />
    </div>
</template>

<script>
import cm from '../../js/common'
import processList from '@/components/content/processes/processList'

export default {
    name: 'Processes',
    components: {
        processList, 
    },
    data () {
        return {
            info: new Object(),     // 进程数据，由 processes updater 负责填充
            interval: 1000,         // 更新时间间隔
            type: "all",
        }
    },
    beforeRouteEnter: (to, from, next) => {
        next(vm => {
            // console.log("processes beforeRouteEnter next");
            cm.process.startUpdater(vm, vm.type);
        })
    },
    beforeRouteUpdate: (to, from, next) => {
        // console.log("processes beforeRouteUpdate rsc");
        next();
    },
    beforeRouteLeave: (to, from, next) => {
        // console.log("processes beforeRouteLeave");
        cm.process.stopUpdater();
        next();
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
