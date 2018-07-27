<template>
    <v-dialog :value="value" max-width="500px" @input="$emit('input', value)" persistent>
        <v-card>
            <v-card-title class="headline " primary-title > About Sysmon </v-card-title>
            <v-card-text>
                Sysmon is a B/S mode system monitor tool. Created by lightimehpq@gmail.com.
            </v-card-text>
            <v-btn outline large color="black" href="https://github.com/Lt0/sysmon" target="blank">Github &ensp;<icon iconClass="github" style="font-size: 2em" /></v-btn>
            <br><br><br>
            <v-divider></v-divider>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" flat @click="$emit('input', !value)" > Close </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    name: "about",
    props: ['value'],
}
</script>
