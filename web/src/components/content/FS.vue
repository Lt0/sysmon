<template>
    <div>
        <div v-for="storage in info.Storage" :key="storage.index">
            <disk-card :storage='storage'></disk-card>
        </div>
    </div>
</template>

<script>
import cm from '../../js/common'
import diskCard from '@/components/content/fs/diskCard'
export default {
    name: 'FS',
    components: {
        diskCard, 
    },
    data () {
        return {
            info: new Object(),     // 进程数据，由 processes updater 负责填充
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
