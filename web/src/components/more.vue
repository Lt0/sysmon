<template>
    <div>
        <v-menu bottom left offset-y min-width="180px">
            <v-btn slot="activator" icon><v-icon>mdi-dots-vertical</v-icon></v-btn>

            <v-list>
                <v-tooltip bottom>
                    <span slot="activator">
                        <v-list-tile @click="clearSetting">
                            <v-list-tile-title>Clear Setting</v-list-tile-title>
                        </v-list-tile>
                    </span>
                    <span>clear all custom setting</span>
                </v-tooltip>
                
                <v-list-tile @click='aboutDialog = true'>
                    <v-list-tile-title>About</v-list-tile-title>
                </v-list-tile>
            </v-list>
        </v-menu>

        <about v-model="aboutDialog" />

        <v-dialog v-model="clearSettingDialog" max-width="500px">
            <v-card>
                <v-card-title class="headline " primary-title > Success </v-card-title>
                <v-card-text>
                    All custon setting was cleared. Refresh the page to take effect.
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary" flat @click="clearSettingDialog = false" > OK </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import about from '@/components/common/about'

export default {
    name: 'more',
    components: {
        about,
    },
    data() {
        return {
            aboutDialog: false,
            clearSettingDialog: false,
        }
    },
    methods: {
        clearSetting() {
            console.log("clear Setting");
            localStorage.clear();
            this.clearSettingDialog = true;
        },
    }, 
}
</script>
