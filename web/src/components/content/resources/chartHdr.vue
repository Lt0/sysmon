<template>
    <div class="chart-hdr">
            <div class="chart-hdr-title">
                <h4>{{ title }}</h4>
            </div>
            <v-dialog v-model="dialog" max-width="500px" >
                <v-btn slot="activator" flat icon color="teal"><v-icon>list</v-icon></v-btn>

                <v-card>
                    <v-card-title>Select items to show</v-card-title>
                    <v-card-text>
                        <v-radio-group v-model="selectedType" row>
                            <v-radio
                                v-for="t in selectTypes"
                                :key="t.index"
                                :label="t.name"
                                :value="t.name"
                            ></v-radio>
                        </v-radio-group>
                        <v-select 
                            :items="items" 
                            :value='selectedItems'
                            @input="updateSelection"
                            multiple 
                            chips
                            deletable-chips
                        >
                        </v-select>
                    </v-card-text>

                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" flat @click.native="dialog = false">Close</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
        </div>
</template>

<script>
export default {
    name: 'chartHdr',
    props: ['title', 'items', 'selectedItems', 'defaultItemNum'],
    model: {
        prop: 'selectedItems',
        event: 'input',
    },
    data () {
        return {
            dialog: false,
            selectTypes: [{name: 'Default'}, {name: 'Items'}, {name: 'All'}, {name: 'Custom'}],
            selectedType: 'Default',
        }
    },
    methods: {
        updateSelection(arg){
            this.$emit('input', arg);
        },
    },
    watch: {
        selectedType: function(){
            console.log("selectedType: " + this.selectedType);
            switch (this.selectedType) {
                case 'Default':
                    let items = [];
                    for (let i = 0; i < this.defaultItemNum; i++){
                        items.push(this.items[i]);
                    }
                    this.$emit('input', items);
                    break;
                case 'Items':
                    let devs = JSON.parse(JSON.stringify(this.items));
                    for (let i = 0; i < this.defaultItemNum; i++){
                        devs.shift();
                    }
                    this.$emit('input', devs);
                    break;
                case 'All':
                    this.$emit('input', this.items);
                    break;
                default:
                    return;
            }
        },
        selectedItems: function(){
            if (this.selectedItems.length == this.items.length) {
                this.selectedType = this.selectTypes[2].name;
                console.log("2 this.selectedType: " + this.selectedType);
                return;
            }

            if (isDefaultItems(this.selectedItems, this.items, this.defaultItemNum)){
                this.selectedType = this.selectTypes[0].name;
                console.log("0 this.selectedType: " + this.selectedType);
                return;
            }

            if (this.items.length - this.selectedItems.length == this.defaultItemNum && noDefaultItem(this.selectedItems, this.items, this.defaultItemNum)) {
                this.selectedType = this.selectTypes[1].name;
                console.log("1 this.selectedType: " + this.selectedType);
                return;
            }

            this.selectedType = this.selectTypes[3].name;
            console.log("3 this.selectedType: " + this.selectedType);
        },
    },
}

function isDefaultItems(selected, items, defaultItemNum){
    if (selected.length != defaultItemNum){
        return false;
    }
    for (let i = 0; i < defaultItemNum; i++){
        if (selected[i] != items[i]){
            return false;
        }
    }

    return true;
}

function noDefaultItem(seleted, items, defaultItemNum){
    let defaults = [];
    for (let i = 0; i < defaultItemNum; i++){
        defaults.push(items[i]);
    }
    for (let i = 0; i < seleted.length; i++){
        if (defaults.indexOf(seleted[i]) >= 0){
            console.log("selected[i]: " + seleted[i] + " match " + items.indexOf(defaults[i]));
            return false
        }
    }

    return true;
}
</script>

<style scoped>
.chart-hdr{
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}
.chart-hdr-title{
  color: #7e858b;
  min-width: 10em;
  text-align: left;
}
</style>
