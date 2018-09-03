<template>
	<div>
		<div id="servers-container">
			<v-list id="servers-btn">
				<v-list-tile @click="servers_visible = !servers_visible">
		            <v-list-tile-action><v-icon color="white">computer</v-icon></v-list-tile-action>
		            <v-list-tile-content><v-list-tile-title> Servers </v-list-tile-title></v-list-tile-content>
		            <v-list-tile-action v-if="!servers_visible"><v-icon color="white">arrow_drop_down</v-icon></v-list-tile-action>
		            <v-list-tile-action v-if="servers_visible"><v-icon color="white">arrow_drop_up</v-icon></v-list-tile-action>
		    	</v-list-tile>
			</v-list>
		</div>
		
		<!--server list-->
		<div v-if="servers_visible">
			<v-list>
			    	<v-list-tile v-for="server in servers" :key=server.id @click="selectServer(server)" class="server ">
			            <v-list-tile-content><div class="caption font-weight-thin"> {{ server }} </div> </v-list-tile-content>
			            <v-list-tile-action v-if="server == activeServer"><v-icon>done</v-icon></v-list-tile-action>
			            <v-list-tile-action v-if="server != activeServer && (hostServer != null && server != hostServer)" @click.stop="removeServer(server)"><v-icon>remove</v-icon></v-list-tile-action>
			    	</v-list-tile>
	    	</v-list>
	    	<div>
	    		<v-btn outline class="server-btn" style="color: grey;" @click="clearServersDialog = true"><v-icon>clear_all</v-icon>&ensp;clear</v-btn>
	    		<v-btn outline class="server-btn" style="color: #26a69a;" @click="addServerDialog = true"><v-icon>add</v-icon>&ensp;add</v-btn>
	    		
	    		<!--add server dialog-->
	    		<v-dialog v-model="addServerDialog" width='400px'>
	    			<div><server-input-field @addServerFinish="addServerFinishHandler()" @cancelAddServer="cancelAddServerHandler()" :enableCancel="true"></server-input-field></div>
	    		</v-dialog>
	    		
	    		<!--clear server dialog-->
	    		<v-dialog v-model="clearServersDialog" width="400px">
	    			<v-card>
						<v-card-title class="headline lighten-2" primary-title>Clear All Servers?</v-card-title>
						<v-card-text>All servers will be removed except Active server and Host server</v-card-text>
						<v-card-actions>
							<confirm-btns leftPrimary rightMinor leftPrependVIcon="exit_to_app" rightPrependVIcon="clear_all" rightText="Clear" @clickLeft='clearServersDialog=false' @clickRight="clearServers()" />
						</v-card-actions>
					</v-card>
	    		</v-dialog>
	    	</div>
    	</div>
    	
    	<!--content list-->
	    <v-list v-if="!servers_visible">
	        <v-list-tile @click="showResources">
	            <v-list-tile-action><icon>chart-line</icon></v-list-tile-action>
	            <v-list-tile-content><v-list-tile-title> Resources </v-list-tile-title></v-list-tile-content>
	        </v-list-tile>
	
	        <v-list-tile @click="showProcesses">
	            <v-list-tile-action><icon>process</icon></v-list-tile-action>
	            <v-list-tile-content><v-list-tile-title>Processes</v-list-tile-title></v-list-tile-content>
	        </v-list-tile>
	
	        <v-list-tile @click="showFS">
	            <v-list-tile-action><icon>harddisk</icon></v-list-tile-action>
	            <v-list-tile-content><v-list-tile-title>File Systems</v-list-tile-title></v-list-tile-content>
	        </v-list-tile>
	    </v-list>
    </div>
</template>

<script>
import cm from '../js/common'
import serverInputField from '@/components/common/serverInputField'
import confirmBtns from '@/components/common/confirmBtns'

export default {
    name: 'DrawerLeft',
    components: {
    	serverInputField,
    	confirmBtns, 
    },
    props: ['drawer'],
    data () {
        return {
        	servers: [],
        	hostServer: cm.sapi.getHostServer(),
        	activeServer: "",
            servers_visible: false,
            addServerDialog: false,
            clearServersDialog: false,
        }
    },
    mounted: function() {
    	// init these servers in order
    	this.updateActiveServer();
    	this.updateServers();
    },
    methods: {
    	updateActiveServer() {
    		this.activeServer = cm.sapi.getActiveServer();
    	},
    	updateServers() {
    		this.servers = cm.sapi.getServers();
    	},
    	selectServer(server){
    		cm.sapi.setActiveServer(server);
    		this.updateActiveServer();
    		location.reload();
    	},
    	removeServer(server) {
    		cm.sapi.delServer(server);
    		this.updateServers();
    	},
    	clearServers() {
    		cm.sapi.clearServers();
    		this.updateServers();
    		this.clearServersDialog = false;
    	},
    	addServerFinishHandler() {
    		this.addServerDialog = false;
    		this.updateServers();
    	},
    	cancelAddServerHandler() {
    		this.addServerDialog = false;
    	},
    	
        showResources () {
            cm.bus.$emit('changeContent', 'resources');
        },
        showProcesses () {
            cm.bus.$emit('changeContent', 'processes');
        },
        showFS () {
            cm.bus.$emit('changeContent', 'fs');
        },
        // showsettings () {
        //     cm.bus.$emit('changeContent', 'setting');
        // },
    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.server {
		color: grey;
		
	}
	.server-btn {
		width: 40%;
		color: grey;
	}
	#newServerInput {
		display: flex;
	}
	#newServerProtoInput {
		width: 6em;
	}
	#servers-container {
		
		background-image: url(/static/img/drawerLeft/material.jpg);
		width: 100%;
		height: 14em;
		background-size: 100% 14em;
	}
	#servers-btn {
		padding-top: 10em;
		color: white;
	}
</style>
