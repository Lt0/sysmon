<template>
    <v-navigation-drawer v-model="drawer" clipped fixed app >
        <v-list>
            <v-list-tile @click="">
                <v-list-tile-action><v-icon>show_chart</v-icon></v-list-tile-action>
                <v-list-tile-content><v-list-tile-title> Resources </v-list-tile-title></v-list-tile-content>
            </v-list-tile>

            <v-list-tile @click="">
                <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                <v-list-tile-content><v-list-tile-title>Processes</v-list-tile-title></v-list-tile-content>
            </v-list-tile>

            <v-list-tile @click="">
                <v-list-tile-action><v-icon>storage</v-icon></v-list-tile-action>
                <v-list-tile-content><v-list-tile-title>File Systems</v-list-tile-title></v-list-tile-content>
            </v-list-tile>

            <v-list-tile @click="">
                <v-list-tile-action><v-icon>settings</v-icon></v-list-tile-action>
                <v-list-tile-content><v-list-tile-title>Settings</v-list-tile-title></v-list-tile-content>
            </v-list-tile>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
export default {
    name: 'DrawerLeft',
    props: ['drawer'],
    data () {

    },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
