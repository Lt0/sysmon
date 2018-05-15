<template>
    <div class="chart-legend" @click="$emit('toggleChart', !hideChart); $emit('updateChart')">
        <div class="color-indicator" :style="{borderColor: borderColor, backgroundColor: backgroundColor}"></div>
        <div class="chart-label" :class="{'hide-chart': hideChart}" :style="{color: labelColor}">
             {{ label }} 
             <div :style="{'min-width': dynamicLabelMinWidth}" class="dynamic-label">
                &ensp;{{ dynamicLabel }}
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'chartLegend',
    // borderColor: 图案边缘颜色
    // backgroundColor: 图案背景颜色
    // hideChart: 是否隐藏 chartjs 中对应的数据图，由调用者通过 v-model 传入
    // label: 文字内容
    props:['borderColor', 'backgroundColor', 'label', 'dynamicLabel', 'dynamicLabelMinWidth', 'labelColor', 'hideChart'],
    model: {
        prop: 'hideChart',
        event: "toggleChart",
    },
}
</script>

<style scoped>
.chart-legend {
    display: flex;
    align-items: center;

    cursor: pointer;
}
.color-indicator {
    height: 1em;
    width: 2em;

    border-width: 2px;
    border-style: solid;
    border-color: #bfbfbf;
}
.chart-label {
    display: flex;
    padding: 0 0 0 0.5em;

    user-select: none;
    -moz-user-select:none;

    cursor:pointer
}
.dynamic-label {
    text-align: left;
}
.hide-chart {
    text-decoration: line-through;
}
</style>