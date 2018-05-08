<template>
    <v-app id="inspire">
        <drawer-left :drawer='drawer' />
        <v-toolbar app fixed clipped-left color="teal lighten-1" dark>
            <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
            <v-toolbar-title>System Monitor</v-toolbar-title>

            <v-spacer></v-spacer>

            <v-btn icon><v-icon>more_vert</v-icon></v-btn>
        </v-toolbar>

        <v-content>
            <v-container fluid fill-height>
                <v-layout justify-center align-center>
                <v-flex shrink>
                    <v-tooltip right>
                    <v-btn
                        slot="activator"
                        :href="source"
                        icon
                        large
                        target="_blank"
                    >
                        <v-icon large>code</v-icon>
                    </v-btn>
                    <span>Source</span>
                    </v-tooltip>
                    <v-tooltip right>
                    <v-btn slot="activator" icon large href="https://codepen.io/johnjleider/pen/qxQWda" target="_blank">
                        <v-icon large>mdi-codepen</v-icon>
                    </v-btn>
                    <span>Codepen</span>
                    </v-tooltip>
                </v-flex>
                </v-layout>
            </v-container>
        </v-content>

        <v-footer app fixed>
            <span>&copy; lightimehpq@gmail.com</span>
        </v-footer>
    </v-app>
</template>

<script>
import DrawerLeft from '@/components/DrawerLeft'

export default {
  name: 'Layout',
  components: {
      DrawerLeft
  },
  data () {
    return {
      drawer: true,
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>