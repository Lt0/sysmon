<template>
	<v-card>
            <v-card-title class="headline " primary-title > Clear All Setting? </v-card-title>
            <v-card-text>
                All setting will be cleared. Including added servers, customization of resources page and more...
            </v-card-text>
            <v-card-actions>
            	<confirm-btns leftPrimary rightMinor rightText="Clear" @clickLeft='$emit("clearSettingCancel")' @clickRight="clearSetting()" />
            </v-card-actions>
            
       </v-card>
</template>

<script>
import confirmBtns from '@/components/common/confirmBtns'

	export default {
		name: 'clearSetting',
		components: {
			confirmBtns,
		},
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