<template>
    <div>
        <v-container id="rsc" fluid>
            <v-layout wrap>
                <v-flex xs12 sm12 md12 lg6 xl6 class="item-flex">
                    <rsc-cpu :cpu='rsc.Cpu' :rscOp='rscOp' :points='points'/>
                </v-flex>

                <v-flex xs12 sm12 md12 lg6 xl6 class="item-flex">
                    <rsc-mem :mem='rsc.Mem' :rscOp='rscOp' :points='points'/>
                </v-flex>

                <v-flex xs12 sm12 md12 lg6 xl6 class="item-flex">
                    <rsc-net :net='rsc.Net' :rscOp='rscOp' :points='points'/>
                </v-flex>

                <v-flex xs12 sm12 md12 lg6 xl6 class="item-flex">
                    <rsc-disk :rsc='rsc' :rscOp='rscOp' :interval='interval'/>
                </v-flex>
            </v-layout>
        </v-container>
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
        intersect: false,   // 设置为 false，则鼠标只要移动到图表内，始终显示 tips，否则，只在鼠标在顶点上时才显示 tips
        displayColors: false,
    },
    animation: {},
    legend: {
        // chartjs 自带的 legend 更新的时候会更新整个图表，为了实现在 legend 显示速率的功能，使用自定义的 lagend 组件
        display: false,
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
        height: '28vh',
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
            rscErr: null,				// 获取系统状态时的错误信息
            interval: 1000,         // 更新时间间隔
        }
    },
    watch: {
    	rsc: function() {
    		cm.bus.$emit("updateInfo", "success");
    	},
    	rscErr: function() {
    		cm.bus.$emit("updateInfo", "failed");
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
        points: function(){
            let self = this;
            let t = this.rscOp.rscChartOp.timeCap;

            let num = Math.floor(t*1000/self.interval);
            if (num < 1) {
                console.log("genSamplePoints: interval too large: " + self.interval);
                return null;
            }

            let ps = [];
            for(let i = 0; i <= num; i++){
                let p = self.interval/1000*i
                ps.push(p.toFixed(1));
            }
            // console.log("recomputed points: " + ps);
            return ps;
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
#rsc{
    width: 99.5%;
    padding: 2em 0 0 0;
}
.item-flex{
    padding: 0 1em 5em 1em;
}
</style>
