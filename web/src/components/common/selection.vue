<template>
    <div>
        <v-dialog v-model="dialog" max-width="500px" >
            <v-btn slot="activator" flat icon color="teal">
                <icon>tools</icon>
            </v-btn>
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
                	<confirm-btns noLeft rightPrimary @clickRight="dialog = false" />
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
import confirmBtns from '@/components/common/confirmBtns'
export default {
    name: 'appSelection',
    // items: 所有的选项
    // selectedItems: 已选项
    // defaultItemNum: 默认选项的数量
    // dataKey: 用来自动记录已选项的 key
    // lazyInit: 是否等待 items 发生变化时才初始化已选项（根据本地自动保存的记录进行初始化）
    props: ['items', 'selectedItems', 'defaultItemNum', 'dataKey', 'lazyInit'],
    components: {
    	confirmBtns,
    },
    model: {
        prop: 'selectedItems',
        event: 'input',
    },
    data(){
        return {
            dialog: false,
            selectTypes: [{name: 'Default'}, {name: 'Items'}, {name: 'All'}, {name: 'Custom'}],
            selectedType: 'Default',
        }
    },
    created() {
        this.initSelectedItems();
    },
    methods: {
        updateSelection(arg){
            this.$emit('input', arg);
            this.saveSelectedItems(arg);
        },
        initSelectedItems() {
            if(this.dataKey) {
                if(!localStorage[this.dataKey]) {
                    return;
                }

                let items = JSON.parse(localStorage[this.dataKey]);
                // console.log("selection init " + this.dataKey + ": " + items);
                this.$emit('input', items);
            }
        }, 
        saveSelectedItems(items) {
            if(this.dataKey) {
                // console.log("go to save " + this.dataKey);
                localStorage.setItem(this.dataKey, JSON.stringify(items))
                // console.log("saved data: " + localStorage[this.dataKey]);
            }
        }
    },
    watch: {
        items: function() {
            this.initSelectedItems();
        },
        selectedType: function(){
            // console.log("selectedType: " + this.selectedType);
            let items = [];
            switch (this.selectedType) {
                case 'Default':
                    for (let i = 0; i < this.defaultItemNum; i++){
                        items.push(this.items[i]);
                    }
                    break;
                case 'Items':
                    items = JSON.parse(JSON.stringify(this.items));
                    for (let i = 0; i < this.defaultItemNum; i++){
                        items.shift();
                    }
                    break;
                case 'All':
                    items = this.items;
                    break;
                default:
                    items = this.selectedItems;
            }

            this.$emit('input', items);
            this.saveSelectedItems(items);
        },
        selectedItems: function(){
            // console.log("selection: selectedItems: " + this.selectedItems);
            if (this.selectedItems.length == this.items.length && this.items.length != 1) {
                this.selectedType = this.selectTypes[2].name;
                // console.log("2 this.selectedType: " + this.selectedType);
                return;
            }

            if (isDefaultItems(this.selectedItems, this.items, this.defaultItemNum)){
                this.selectedType = this.selectTypes[0].name;
                // console.log("0 this.selectedType: " + this.selectedType);
                return;
            }

            if (this.items.length - this.selectedItems.length == this.defaultItemNum && noDefaultItem(this.selectedItems, this.items, this.defaultItemNum)) {
                this.selectedType = this.selectTypes[1].name;
                // console.log("1 this.selectedType: " + this.selectedType);
                return;
            }

            this.selectedType = this.selectTypes[3].name;
            // console.log("3 this.selectedType: " + this.selectedType);
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
            // console.log("selected[i]: " + seleted[i] + " match " + items.indexOf(defaults[i]));
            return false
        }
    }

    return true;
}
</script>
