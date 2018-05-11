<template>
    <div>
        <rsc-mem :rsc='rsc' :lineChartOptions='lineChartOptions' :vueChartOp='vueChartOp' :interval='interval'/>
    </div>
</template>

<script>
import cm from '../../js/common'
import RscMem from '@/components/content/resources/mem'

// line chart 的配置模板
let op = {
    maintainAspectRatio: false,
    title: {
        display: true,
    },
    tooltips: {
        intersect: false,   // 设置为 false，则鼠标只要移动到图表内，都适中显示 tips，否则，只在鼠标在顶点上时才显示 tips
    },
    animation: {
    },
    legend: {
        // 图标位置
        position: 'bottom',
    },
    elements: {
        line: {
            borderWidth: 1,
            //borderColor: 'rgba(0, 0, 0, 0.5)',    // 设置线条颜色
        },
        point: {
            radius: 0,  // 设置数据点的顶点大小
        },
    },
    scales: {
        // y 轴设置
        yAxes: [{
            position: 'right',
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
            },
            scaleLabel: {},
        }],
    }
};

// vue-chartjs 的配置选项
let vueChartOp = {
    styles: {
        width: '100%',
        height: '23vh',
        position: 'relative',
    },
}

export default {
    name: 'Resources',
    components: {
        RscMem,
    },
    data () {
        return {
            rsc: new Object(),
            interval: 1000,         // 更新时间间隔
            vueChartOp: vueChartOp,
        }
    },
    computed: {
        rscStr: function(){
            return JSON.stringify(this.rsc);
        },
        lineChartOptions: function(){
            let o = Object.assign({}, op);
            // 默认的动画时长为 1000 毫秒，如果动画时长大于 interval，则会导致显示的 tips 在鼠标移开图表后无法自动关闭
            o.animation.duration = this.interval - 1;
            return o;
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
