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
			    	<v-list-tile v-for="server in servers" :key=server.id @click="selectServer(server)">
			            <v-list-tile-content><v-list-tile-title> {{ server }} </v-list-tile-title></v-list-tile-content>
			            <v-list-tile-action v-if="server == activeServer"><v-icon>done</v-icon></v-list-tile-action>
			            <v-list-tile-action v-if="server != activeServer && (hostServer != null && server != hostServer)" @click.stop="removeServer(server)"><v-icon>remove</v-icon></v-list-tile-action>
			    	</v-list-tile>
	    	</v-list>
	    	<div>
	    		<v-btn outline class="server-btn" @click="clearServers()"><v-icon>clear_all</v-icon></v-btn>
	    		<v-btn outline class="server-btn" @click="addServerDialog = true"><v-icon>add</v-icon></v-btn>
	    		
	    		<v-dialog v-model="addServerDialog" width="400">
	    			<v-card>
	    				<v-card-title class="headline grey lighten-2" primary-title>Add Server</v-card-title>
	    				<v-card-text>
	    					<div id="newServerInput">
		    					<div id="newServerProtoInput">
		    						<v-select v-model="newServerProto" :items="serverProtocals" outline class="serverProtocalSelector"></v-select>
		    					</div>
		    					<v-text-field v-model="newServer" label="Server" hint="eg. https://www.xxx.com:2048" prefix="://" clearable @keyup.enter="addServer()"></v-text-field>
		    				</div>
	    				</v-card-text>
	    				<v-card-actions>
	    					<v-spacer></v-spacer>
	    					<!--<v-btn @click="addServerDialog = false">Cancel</v-btn>-->
	    					<v-btn @click="addServer()" outline>Add</v-btn>
	    					
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
	
	        <!-- <v-list-tile @click="showsettings">
	            <v-list-tile-action><v-icon>mdi-settings-outline</v-icon></v-list-tile-action>
	            <v-list-tile-content><v-list-tile-title>Settings</v-list-tile-title></v-list-tile-content>
	        </v-list-tile> -->
	    </v-list>
    </div>
</template>

<script>
import cm from '../js/common'

export default {
    name: 'DrawerLeft',
    props: ['drawer'],
    data () {
        return {
        	servers: [],
        	activeServer: "",
        	hostServer: null,
            servers_visible: false,
            addServerDialog: false,
            newServer: "",
            newServerProto: "http",
            serverProtocals: ["http", "https"],
        }
    },
    watch: {
    	servers: function() {
    		this.saveServers(this.servers);
    	},
    	activeServer: function() {
    		this.saveActiveServer(this.activeServer);
    	},
    },
    mounted: function() {
    	// init these servers in order
    	this.updateHostServer();
    	this.updateActiveServer();
    	this.updateServers();
    },
    methods: {
    	saveServers(servers){
    		localStorage.setItem("servers", JSON.stringify(servers));
    	},
    	saveActiveServer(server) {
    		localStorage.setItem("activeServer", server);
    	},
    	updateHostServer() {
    		if(location.protocol != "http:" && location.protocol  != "https:") {
    			this.hostServer = null
    		} else {
    			this.hostServer = location.protocol + "//" + location.hostname + ":" + location.port;
    		}
    		console.log("updateHostServer: this.hostServer: ", this.hostServer);
    	},
    	updateActiveServer() {
    		this.activeServer = localStorage.getItem("activeServer");
    		if(!this.activeServer) {
    			this.activeServer = this.hostServer
    		}
    		console.log("updateActiveServer to ", this.activeServer);
    	},
    	updateServers() {
    		let foundActive = false;
    		let foundHost = false;
    		
    		let localServers = JSON.parse(localStorage.getItem("servers"));
    		if(!localServers) {
    			console.log("no localServers, init a new one\n");
    			localServers = [];
    		}
    		for(let i in localServers) {
    			if(localServers[i] == this.activeServer) {
    				foundActive = true;
    			}
    			if(localServers[i] == this.hostServer) {
    				foundHost = true;
    			}
    		}

    		if(!foundActive) {
    			localServers.push(this.activeServer);
    			console.log("updateServers: add active server: ", localServers);
    		}
    		if(!foundHost && this.hostServer != null && this.hostServer != this.activeServer) {
    			localServers.push(this.hostServer);
    			console.log("updateServers: add host server: ", localServers);
    		}
    		console.log("updateServers: ", localServers);
    		// watcher of this.servers will save it to localStorage automatically.
    		this.servers = localServers;
    	},
    	selectServer(server){
    		// watcher of this.activeServer will save it to localStorage automatically.
    		this.activeServer = server;
    		location.reload();
    	},
    	addServer() {
    		this.addServerDialog = false;
    		let server = this.newServerProto + "://" + this.newServer;
    		for(let i in this.servers) {
    			if(server == this.servers[i]) {
    				return;
    			}
    		}
    		this.servers.push(server)
    	},
    	removeServer(server) {
    		console.log("remove ", server);
    		for(let item in this.servers) {
    			console.log("checking ", item);
    			if(this.servers[item] == server) {
    				this.servers.splice(item, 1);
    				break;
    			}
    		}
    	},
    	clearServers() {
    		if(this.hostServer) {
    			this.servers = [this.hostServer, this.activeServer]
    		} else {
    			this.servers = [this.activeServer]
    		}
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
	.server-btn {
		width: 40%;
		color: #26a69a;
	}
	#newServerInput {
		display: flex;
	}
	#newServerProtoInput {
		width: 6em;
	}
</style>
