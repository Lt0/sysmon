<template>
    <div>
        <v-menu bottom left offset-y min-width="180px">
            <v-btn slot="activator" icon><v-icon>more_vert</v-icon></v-btn>

            <v-list>
                <v-list-tile @click="">
                    <v-list-tile-title @click.stop="aboutDialog = true">About</v-list-tile-title>
                </v-list-tile>
            </v-list>
        </v-menu>

        <v-dialog v-model="aboutDialog" max-width="500px">
            <v-jumbotron color="white">
                <v-container fill-height>
                <v-layout align-center>
                    <v-flex>
                        <h1>About Sysmon</h1>
                        <span>Sysmon is a B/S mode system monitor tool. It's frontend was based on vue and vuetify and backend was based on go and beego.</span>
                        <v-divider class="my-3"></v-divider>
                        <v-btn outline large color="black" href="https://github.com/Lt0/sysmon" target="blank">Github<v-icon right>mdi-github-circle</v-icon></v-btn>
                    </v-flex>
                </v-layout>
                </v-container>
            </v-jumbotron>
        </v-dialog>
    </div>
</template>

<script>
export default {
    name: 'more',
    data() {
        return {
            aboutDialog: false,
        }
    },
}
</script>
