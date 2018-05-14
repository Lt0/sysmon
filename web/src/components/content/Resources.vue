<template>
    <div style="width: 99%">
        <rsc-cpu :rsc='rsc' :rscOp='rscOp' :interval='interval'/>
        <rsc-mem :rsc='rsc' :rscOp='rscOp' :interval='interval'/>
        <rsc-net :rsc='rsc' :rscOp='rscOp' :interval='interval'/>
        <rsc-disk :rsc='rsc' :rscOp='rscOp' :interval='interval'/>
    </div>
</template>

<script>
import cm from '../../js/common'
import RscMem from '@/components/content/resources/mem'
import RscCpu from '@/components/content/resources/cpu'
import RscNet from '@/components/content/resources/net'
import RscDisk from '@/components/content/resources/disk'

// chartjs 的配置模板
let chartJsOpTpl = {
    maintainAspectRatio: false,
    title: {
        display: false,
    },
    tooltips: {
        intersect: false,   // 设置为 false，则鼠标只要移动到图表内，都适中显示 tips，否则，只在鼠标在顶点上时才显示 tips
        displayColors: false,
    },
    animation: {},
    legend: {
        display: true,
        // 图标位置
        position: 'bottom',
    },
    elements: {
        line: {
            borderWidth: 1,
            tension: 0,
        },
        point: {
            radius: 0,  // 设置数据点的顶点大小
        },
    },
    scales: {
        // y 轴设置
        yAxes: [{
            position: 'left',
            ticks: {
                beginAtZero: true,
                max: 100,
                stepSize: 20,
            },
            scaleLabel: {},
        }],
        xAxes: [{
            ticks: {
                // 在 enable autoSkip 的前提下，调整 x 轴的 padding 距离，仅对 x 轴有效
                autoSkipPadding: 100,
                // 设置 x 轴坐标轴文字最大旋转角度
                maxRotation: 0,
            },
            scaleLabel: {},
        }],
    }
};

// vue-chartjs 的配置模板
let vueChartOpTpl = {
    styles: {
        width: '100%',
        height: '21vh',
        position: 'relative',
    },
}

// resources chart 配置模板
let rscChartOpTpl = {
    timeCap: 60,   // 显示时长
}

export default {
    name: 'Resources',
    components: {
        RscMem, RscCpu, RscNet, RscDisk, 
    },
    data () {
        return {
            rsc: new Object(),      // 系统信息状态数据，由 resources updater 负责填充
            interval: 1000,         // 更新时间间隔
        }
    },
    computed: {
        rscStr: function(){
            return JSON.stringify(this.rsc);
        },
        rscOp: function(){
            let chartJsOp = Object.assign({}, chartJsOpTpl);
            // 默认的动画时长为 1000 毫秒，如果动画时长大于 interval，则会导致显示的 tips 在鼠标移开图表后无法自动关闭
            // 由于更新 legend 会重绘整个图表，导致动画重新开始，暂时先通过全局配置关闭动画效果。
            chartJsOp.animation.duration = this.interval - 1;
            let vueChartOp = Object.assign({}, vueChartOpTpl);
            let rscChartOp = Object.assign({}, rscChartOpTpl);
            return {chartJsOp, vueChartOp, rscChartOp};
        },
    },
    created () {
        cm.rsc.startUpdater(this);
    },
    beforeRouteEnter: (to, from, next) => {
        next(vm => {
            //console.log("beforeRouteEnter next");
        })
    },
    beforeRouteUpdate: (to, from, next) => {
        //console.log("beforeRouteUpdate rsc");
        next();
    },
    beforeRouteLeave: (to, from, next) => {
        //console.log("beforeRouteLeave rsc");
        next();
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
