<template>
    <div id="more-container">
        <v-menu bottom left offset-y min-width="180px" id="more-btn">
        	<!--more button-->
            <v-btn slot="activator" icon><icon>moreif</icon></v-btn>

			<!--option list-->
            <v-list>
                <v-tooltip bottom>
                    <span slot="activator">
                        <v-list-tile @click="clearSettingDialog = true"><v-list-tile-title><icon>custom-clear</icon>&ensp;Clear Setting</v-list-tile-title></v-list-tile>
                    </span>
                    <span>clear all custom setting</span>
                </v-tooltip>
                
                <v-list-tile @click='aboutDialog = true'><v-list-tile-title><icon>about-o</icon>&ensp;About</v-list-tile-title></v-list-tile>
            </v-list>
        </v-menu>

		<!--dialogs--> 
        <about v-model="aboutDialog" />
        <v-dialog v-model="clearSettingDialog" max-width="400px">
        	<clear-setting @clearSettingCancel="clearSettingDialog = false" />
        </v-dialog>
    </div>
</template>

<script>
import about from '@/components/common/about'
import clearSetting from '@/components/header/clearSetting'

export default {
    name: 'more',
    components: {
        about,
        clearSetting,
    },
    data() {
        return {
            aboutDialog: false,
            clearSettingDialog: false,
        }
    },
}
</script>

<style scoped>
	#more-container {
		margin-right: 0px;
	}
</style>
