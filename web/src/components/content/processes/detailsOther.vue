<template>
    <div id="sched"> 
        <h2>Sched</h2>
        <v-divider></v-divider>
        <br>
        <div>
            <pre class="rawData">{{ Sched }}</pre>
        </div>
    </div>
</template>

<script>
import cm from '../../../js/common'
import tips from '../../../js/tips'

export default {
    name: "detailsOther",
    props: ['info'],
    data() {
        return {
            
        }
    },
    computed: {
        Sched: function() {
            return this.info.RDSched || "no Sched info to show";
        }
    }
}
</script>

<style scoped>
#sched {
    text-align: left;
}

.rawData {
    font-family: 'Liberation Mono', 'Courier', 'Courier New', '宋体';
}
</style>

