<template>
	<div id="container">
		<v-card flat width="100%">
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
				<confirm-btns leftMinor rightPrimary leftPrependVIcon="exit_to_app" rightPrependVIcon="add" rightText="Add" @clickLeft='cancel()' @clickRight="addServer()" />
			</v-card-actions>
		</v-card>
	</div>
</template>

<script>
	import cm from '../../js/common'
	import confirmBtns from '@/components/common/confirmBtns'
	
	export default {
		name: 'serverInputField',
		props: ['enableCancel'],		
		components: {
			confirmBtns,
		},
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
		width: 100%;
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
	#server-test-field {
		width: 100%;
	}
</style>