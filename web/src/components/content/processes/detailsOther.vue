<template>
    <div v-html="Sched">
        <!-- {{ Sched }} -->
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
