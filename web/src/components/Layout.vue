<template>
    <v-app id="inspire">
        <v-navigation-drawer v-model="drawer" clipped fixed app ><drawer-left /></v-navigation-drawer>
        <v-toolbar app fixed clipped-left color="teal lighten-1" dark>
            <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
            <v-toolbar-title>System Monitor</v-toolbar-title>

            <v-spacer></v-spacer>

            <more />
        </v-toolbar>

        <v-content>
            <v-container fluid fill-height class="container" id="main-container">
                <v-layout justify-center id="main-layout">
                <v-flex id="main-flex">
                    <keep-alive>
                        <router-view></router-view>
                    </keep-alive>
                </v-flex>
                </v-layout>
            </v-container>
        </v-content>

        <!-- <v-footer app fixed>
            <div>&copy; lightimehpq@gmail.com</div>
        </v-footer> -->
    </v-app>
</template>

<script>
import cm from '../js/common'
import DrawerLeft from '@/components/DrawerLeft'
import more from '@/components/more'

export default {
    name: 'Layout',
    components: {
        DrawerLeft,
        more
    },
    data () {
        return {
            drawer: false,
            infoAll: null,
        }
    },
    beforeCreate () {
        
    },
    created () {
        console.log("created layout");
        cm.updateSysInfoAll();
        cm.bus.$on('changeContent', this.changeContentHandler);
    },
    methods: {
        changeContentHandler (r) {
            switch (r) {
                case 'resources':
                    this.$router.push('resources');
                    break;
                case 'processes':
                    this.$router.push('processes');
                    break;
                case 'fs':
                    this.$router.push('fs');
                    break;
                // case 'setting':
                //     this.$router.push('setting');
                //     break;
                default:
                    console.log('unknow page to route: ', r);
            }
        },

    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container{
    padding: 0;
}

#main-flex {
    max-width: 100%;
}
</style>