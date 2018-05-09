<template>
    <v-app id="inspire">
        <v-navigation-drawer v-model="drawer" clipped fixed app ><drawer-left /></v-navigation-drawer>
        <v-toolbar app fixed clipped-left color="teal lighten-1" dark>
            <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
            <v-toolbar-title>System Monitor</v-toolbar-title>

            <v-spacer></v-spacer>

            <v-btn icon><v-icon>more_vert</v-icon></v-btn>
        </v-toolbar>

        <v-content>
            <v-container fluid fill-height>
                <v-layout justify-center align-center>
                <v-flex shrink>
                    <router-view></router-view>
                </v-flex>
                </v-layout>
            </v-container>
        </v-content>

        <v-footer app fixed>
            <div>&copy; lightimehpq@gmail.com</div>
        </v-footer>
    </v-app>
</template>

<script>
import cm from '../js/common'
import DrawerLeft from '@/components/DrawerLeft'

export default {
    name: 'Layout',
    components: {
        DrawerLeft
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
                case 'setting':
                    this.$router.push('setting');
                    break;
                default:
                    console.log('unknow page to route: ', r);
            }
        },

    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>