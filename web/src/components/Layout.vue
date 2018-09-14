<template>
    <v-app id="inspire">
        <v-navigation-drawer v-model="drawer" clipped fixed app ><drawer-left /></v-navigation-drawer>
        <v-toolbar app fixed clipped-left color="teal lighten-1" dark>
            <v-toolbar-side-icon id="side-btn" @click.stop="drawer = !drawer"></v-toolbar-side-icon>
            <v-spacer></v-spacer>
            <v-toolbar-title class="hdr-text" @click="serverDialog = true" :class="{uif: updateInfoFailed}">{{ hostname || activeServer}}</v-toolbar-title>

            <v-spacer></v-spacer>

            <more />
        </v-toolbar>

        <v-content>
            <v-container fluid fill-height class="container" id="main-container">
                <v-layout justify-center id="main-layout">
                <v-flex id="main-flex">
                    <!-- https://cn.vuejs.org/v2/api/#keep-alive -->
                    <keep-alive include="Resources">
                        <router-view></router-view>
                    </keep-alive>
                </v-flex>
                </v-layout>
            </v-container>
        </v-content>

        <!-- <v-footer app fixed>
            <div>&copy; lightimehpq@gmail.com</div>
        </v-footer> -->

        <!--<about v-model="aboutDialog" />-->
        <v-dialog v-model="serverDialog" width='400px'>
        	<v-card>
        		<server-info :hostname="hostname" :server="activeServer"></server-info>
        		<v-card-actions><confirm-btns noLeft rightPrimary @clickRight="serverDialog = false"></confirm-btns></v-card-actions>
        	</v-card>
		</v-dialog>
    </v-app>
</template>

<script>
import cm from '../js/common'
import DrawerLeft from '@/components/DrawerLeft'
import more from '@/components/more'
import serverInfo from '@/components/header/serverInfo'
import confirmBtns from '@/components/common/confirmBtns'

export default {
    name: 'Layout',
    components: {
        DrawerLeft,
        more, 
        serverInfo, 
        confirmBtns, 
    },
    data () {
        return {
            drawer: false,
            infoAll: null,
            serverDialog: false,
            activeServer: cm.sapi.getActiveServer(),
            hostname: null,
            updateInfoFailed: false,
        }
    },
    beforeCreate () {
    	cm.sapi.getActiveHostname(this);
    },
    created () {
        cm.bus.$on('changeContent', this.changeContentHandler);
        cm.bus.$on('updateInfo', this.updateInfoHandler);
    },
    methods: {
        changeContentHandler (r) {
        	console.log("this.$router: ", this.$router);
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
		updateInfoHandler(arg) {
			switch(arg) {
				case "success":
					this.updateInfoFailed = false;
					break;
				case "failed":
					this.updateInfoFailed = true;
					break;
				default:
					console.log("updateInfoHandler: unknown status: ", arg);
			}
		}
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

.hdr-text {
    cursor: pointer;
    font-size: 1em;
    text-align: center;
    margin: 0;
}

#side-btn {
	margin: 6px;
}

.uif {
	color: #f7abab;
}
</style>