<template>
    <v-list>
        <v-list-tile @click="showResources">
            <v-list-tile-action><v-icon>mdi-chart-line-variant</v-icon></v-list-tile-action>
            <v-list-tile-content><v-list-tile-title> Resources </v-list-tile-title></v-list-tile-content>
        </v-list-tile>

        <v-list-tile @click="showProcesses">
            <v-list-tile-action><v-icon>mdi-format-list-bulleted</v-icon></v-list-tile-action>
            <v-list-tile-content><v-list-tile-title>Processes</v-list-tile-title></v-list-tile-content>
        </v-list-tile>

        <v-list-tile @click="showFS">
            <v-list-tile-action><v-icon>mdi-harddisk</v-icon></v-list-tile-action>
            <v-list-tile-content><v-list-tile-title>File Systems</v-list-tile-title></v-list-tile-content>
        </v-list-tile>

        <!-- <v-list-tile @click="showsettings">
            <v-list-tile-action><v-icon>settings</v-icon></v-list-tile-action>
            <v-list-tile-content><v-list-tile-title>Settings</v-list-tile-title></v-list-tile-content>
        </v-list-tile> -->
    </v-list>
</template>

<script>
import cm from '../js/common'

export default {
    name: 'DrawerLeft',
    props: ['drawer'],
    data () {
        return {
            a: null,
        }
    },
    methods: {
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

</style>
