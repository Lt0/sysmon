<template>
    <svg class="icon" aria-hidden="true">
        <use :xlink:href="iconName"></use>
    </svg>
</template>

<script>
import './iconfont'

export default {
    name: 'icon',
    props: {iconClass: {
        type: String, 
        required: true,
    }},

    computed: {
        iconName: function() {
            let n = "#icon-" + this.iconClass;
            console.log("name: " + n)
            return "#icon-" + this.iconClass;
        }
    },
}
</script>


<style scoped>
.icon {
    width: 0.8em; height: 0.8em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
 }
</style>
