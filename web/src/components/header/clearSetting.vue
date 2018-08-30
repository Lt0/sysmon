<template>
	<v-card>
            <v-card-title class="headline " primary-title > Clear All Setting? </v-card-title>
            <v-card-text>
                All setting will be cleared. Including added servers, customization of resources page and more...
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn style="color:#26a69a" flat @click='$emit("clearSettingCancel")' > Cancel </v-btn>
                <v-spacer></v-spacer>
                <v-btn style="color:grey" flat @click="clearSetting()" > Clear </v-btn>
                <v-spacer></v-spacer>
            </v-card-actions>
       </v-card>
</template>

<script>
	export default {
		name: 'clearSetting',
		methods: {
			clearSetting(){
            	localStorage.clear();
            	this.$router.push("/init");
			},
		},
	}
</script>

<style>
</style>