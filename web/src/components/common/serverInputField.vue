<template>
	<div id="container">
		<v-card flat>
			<v-card-title class="headline lighten-2" primary-title>Add Server</v-card-title>
			<v-card-text>
				<div id="newServerInput">
					<div id="newServerProtoInput">
						<v-select v-model="newServerProto" :items="serverProtocols" outline class="serverProtocalSelector"></v-select>
					</div>
					<div id="server-test-field">
						<v-text-field v-model="newServer" label="Server" hint="eg. https://www.xxx.com:2048" prefix="://" clearable autofocus browser-autocomplete @keyup.enter="addServer()"></v-text-field>
					</div>
				</div>
				<div class="error-msg">{{ errMsg }}</div>
			</v-card-text>
			<v-card-actions>
				<v-spacer></v-spacer>
				<v-btn v-show="enableCancel" @click="cancel()" style="color:grey" flat><v-icon>exit_to_app</v-icon>&ensp;Cancel</v-btn>
				<v-spacer v-if="enableCancel"></v-spacer>
				<v-btn @click="addServer()" style="color:#26a69a" flat><v-icon>add</v-icon>&ensp;Add</v-btn>
				<v-spacer></v-spacer>
			</v-card-actions>
		</v-card>
	</div>
</template>

<script>
	import cm from '../../js/common'
	
	export default {
		name: 'serverInputField',
		props: ['enableCancel'],		
		data () {
			return {
				newServer: "",
            	newServerProto: "http",
            	serverProtocols: ["http", "https"],
            	errMsg: " ",
			}
		},
		methods: {
			addServer: function() {
				console.log("this.newServer: ", this.newServer);
				if(this.newServer.length < 3) {
					this.errMsg = "Invalid";
					return;
				}
				let server = this.newServerProto + "://" + this.newServer;
				this.errMsg = cm.sapi.addServer(server);
				if(this.errMsg == null) {
					this.$emit('addServerFinish', server);
				}
			},
			cancel() {
				this.$emit('cancelAddServer');
			},
		}
	}
</script>

<style>
	#container {
		margin: auto
	}
	#newServerInput {
		display: flex;
	}
	#newServerProtoInput {
		/*width: 5em;*/
	}
	.error-msg {
		color: red;
		text-align: right;
	}
</style>