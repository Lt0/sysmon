<template>
    <v-dialog v-model="dialog" max-width="480px" >
        <v-card>
            <v-card-title class="headline " primary-title > About Sysmon </v-card-title>
            <v-card-text>
                Sysmon is a B/S mode system monitor tool. Created by lightimehpq@gmail.com.
            </v-card-text>
            <v-btn outline large color="black" href="https://github.com/Lt0/sysmon" target="blank">Github &ensp;<icon style="font-size: 2em">github</icon></v-btn>
            <br><br><br>
            <v-divider></v-divider>
            <v-card-actions>
                <confirm-btns noLeft rightPrimary @clickRight="$emit('input', !value)" />
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import confirmBtns from '@/components/common/confirmBtns'
export default {
    name: "about",
    props: ['value'],
    components: {
    	confirmBtns,
    },
    model: {
        prop: 'value',
        event: 'input',
    },
    data() {
        return {
            dialog: false,
        }
    },
    watch: {
        value: function() {
            this.dialog = this.value;
        },
        dialog: function() {
            this.$emit('input', this.dialog)
        },
    },
}
</script>
