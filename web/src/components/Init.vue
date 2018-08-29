
<template>
	<v-container fluid fill-height id="init-container">
	    <v-layout wrap align-center justify-center>
	        <v-flex xs12 sm6 md4 lg3 xl1>
	        	<server-input-field class="input_field" @addServerFinish="addServerFinishHandler($event)"></server-input-field>
	        </v-flex>
	    </v-layout>
	</v-container>
</template>

<script>
	import cm from '../js/common'
	import serverInputField from '@/components/common/serverInputField'
	
	export default {
		name: 'Init',
		components: {
			serverInputField, 
		},
		data () {
			return {
				pageName: "init page",
			}
		},
		mounted: function() {
			console.log("init component mouonted")
			
			let servers = cm.sapi.getServers();
			if(servers) {
				console.log("don't need to init");
				this.$router.push('main');
				return;
			}
			
			if(location.protocol == "http:" || location.protocol == "https:") {
				console.log("init automatically with protocol ", location.protocol);
				this.autoInit();
				this.$router.push('main');
				return;
			}
		},
		methods: {
			addServerFinishHandler(server) {
				console.log("addServerFinishHandler: server: ", server);
				cm.sapi.setActiveServer(server);
				location.reload();
			},
			autoInit() {
				let hostServer = cm.sapi.getHostServer();
				cm.sapi.setActiveServer(hostServer)
				let servers = [hostServer]
				cm.sapi.setServers(servers);
			}
		}
	}
</script>

<style>
	#init-container {
		/*background-color: white;;*/
	}
	/*.input_field {
		width: 400px;
	}*/
</style>