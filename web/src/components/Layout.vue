<template>
    <div>
        <v-navigation-drawer app>nav panel</v-navigation-drawer>
        <v-toolbar app>
            <v-toolbar-side-icon></v-toolbar-side-icon>
            <v-toolbar-title class="white--text">Title</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn icon>
            <v-icon>search</v-icon>
            </v-btn>
            <v-btn icon>
            <v-icon>apps</v-icon>
            </v-btn>
            <v-btn icon>
            <v-icon>refresh</v-icon>
            </v-btn>
            <v-btn icon>
            <v-icon>more_vert</v-icon>
            </v-btn>
        </v-toolbar>
        
        <v-content>
            <v-container fluid>
                <div>
                    {{msg}}
                </div>
            </v-container>
        </v-content>
        <v-footer app>lightimehpq@gmail.com</v-footer>
  </div>
</template>

<script>
import Toolbar from '@/components/Toolbar'

export default {
  name: 'Layout',
  components: {
      Toolbar
  },
  data () {
    return {
      msg: 'layout all things here'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>