<template>
    <div id="fs">
        <v-container id="fs-container">
            <v-layout wrap align-center>
                    <v-flex v-for="storage in Storages" :key="storage.index" xs12 sm4 md4 lg3 xl2>
                        <disk-card :storage='storage'></disk-card>
                    </v-flex>
            </v-layout>
        </v-container>
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
    computed: {
        Storages: function() {
            if(!this.info || !this.info.Storage) {
                return [];
            }

            let valideStorages = [];
            for(let i = 0; i < this.info.Storage.length; i++) {
                if(this.info.Storage[i].Fs != "none") {
                    valideStorages.push(this.info.Storage[i])
                }
            }
            return valideStorages;
        },
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#fs-container {
    padding: 1em;
}
</style>
