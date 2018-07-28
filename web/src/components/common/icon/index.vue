<template>
    <svg class="icon" aria-hidden="true">
        <use :xlink:href="iconName"></use>
        <slot></slot>
    </svg>
</template>

<script>
import './iconfont'

export default {
    name: 'icon',
    computed: {
        iconName: function() {
            return "#icon-" +  this.$slots.default[0].text;
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
