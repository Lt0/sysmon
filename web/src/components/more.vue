<template>
    <div>
        <v-menu bottom left offset-y min-width="180px">
            <v-btn slot="activator" icon><v-icon>more_vert</v-icon></v-btn>

            <v-list>
                <v-tooltip bottom>
                    <span slot="activator">
                        <v-list-tile @click="clearSetting">
                            <v-list-tile-title>Clear Setting</v-list-tile-title>
                        </v-list-tile>
                    </span>
                    <span>clear all custom setting</span>
                </v-tooltip>
                
                <v-list-tile @click="aboutDialog = true">
                    <v-list-tile-title>About</v-list-tile-title>
                </v-list-tile>
            </v-list>
        </v-menu>

        <v-dialog v-model="aboutDialog" max-width="500px">
            <v-card>
                <v-card-title class="headline " primary-title > About Sysmon </v-card-title>
                <v-card-text>
                    Sysmon is a B/S mode system monitor tool. Created by lightimehpq@gmail.com.
                </v-card-text>
                <v-btn outline large color="black" href="https://github.com/Lt0/sysmon" target="blank">Github<v-icon right>mdi-github-circle</v-icon></v-btn>
                <br><br><br>
                <v-divider></v-divider>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary" flat @click="aboutDialog = false" > OK </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

        <v-dialog v-model="clearSettingDialog" max-width="500px">
            <v-card>
                <v-card-title class="headline " primary-title > Success </v-card-title>
                <v-card-text>
                    All custon setting was cleared. Refresh the page to take effect.
                </v-card-text>
                <v-card-actions>
                    
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
export default {
    name: 'more',
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
