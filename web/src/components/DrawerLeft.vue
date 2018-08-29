<template>
	<div>
		<v-list>
				<v-list-tile @click="servers_visible = !servers_visible">
		            <v-list-tile-action><v-icon>computer</v-icon></v-list-tile-action>
		            <v-list-tile-content><v-list-tile-title> Servers </v-list-tile-title></v-list-tile-content>
		            <v-list-tile-action v-if="!servers_visible"><v-icon>expand_more</v-icon></v-list-tile-action>
		            <v-list-tile-action v-if="servers_visible"><v-icon>expand_less</v-icon></v-list-tile-action>
		    	</v-list-tile>
		</v-list>
		
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
	    		
	    		<v-dialog v-model="addServerDialog" width="400">
	    			<div><server-input-field @addServerFinish="addServerFinishHandler()" @cancelAddServer="cancelAddServerHandler()" :enableCancel="true"></server-input-field></div>
	    		</v-dialog>
	    		<v-dialog v-model="clearServersDialog" width="400">
	    			<v-card>
						<v-card-title class="headline lighten-2" primary-title>Clear All Servers?</v-card-title>
						<v-card-text>All servers will be removed except Active server and Host server</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn @click="clearServersDialog=false" flat style="color: #26a69a;">Cancel</v-btn>
							<v-spacer></v-spacer>
							<v-btn @click="clearServers()" flat style="color: grey;">Clear</v-btn>
							<v-spacer></v-spacer>
						</v-card-actions>
					</v-card>
	    		</v-dialog>
	    	</div>
    	</div>
    	
	    <v-list v-if="!servers_visible">
	    	<v-divider />
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

export default {
    name: 'DrawerLeft',
    components: {
    	serverInputField,
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
</style>
