<template>
	<div id="container">
		<v-spacer v-if="!noLeft"></v-spacer>
		
        <v-btn v-if="!noLeft"
        	:class="{primaryBtn: leftPrimary, minorBtn: leftMinor}" 
        	flat 
        	@click='$emit("clickLeft")' >
        	<v-icon v-if="leftPrependVIcon">{{ leftPrependVIcon }}</v-icon> 
        	<icon v-if="leftPrependIcon">{{ leftPrependIcon }}</icon> 
        	<span v-if="leftPrependVIcon || leftPrependIcon">&ensp;</span>
        	{{ leftText || "Cancel" }}
        </v-btn>
        
        <v-spacer></v-spacer>
        
        <v-btn  v-if="!noRight"
        	:class="{primaryBtn: rightPrimary, minorBtn: rightMinor}" 
        	flat 
        	@click='$emit("clickRight")'> 
        	<v-icon v-if="rightPrependVIcon">{{ rightPrependVIcon }}</v-icon> 
        	<icon v-if="rightPrependIcon">{{ rightPrependIcon }}</icon> 
        	<span v-if="rightPrependVIcon || rightPrependIcon">&ensp;</span>
        	{{ rightText || "OK" }} 
        </v-btn>
        
        <v-spacer v-if="!noRight"></v-spacer>
	</div>
</template>

<script>
	export default {
		name: 'confirmBtns',
		props: {
			leftPrependIcon: String,
			rightPrependIcon: String,
			leftPrependVIcon: String,
			rightPrependVIcon: String,
			leftText: String,
			rightText: String,
			leftPrimary: Boolean,
			rightPrimary: Boolean,
			leftMinor: Boolean,
			rightMinor: Boolean,
			noLeft: Boolean,
			noRight: Boolean,
		},
	}
</script>

<style>
	#container {
		display: flex;
		width: 100%;
	}
	.primaryBtn {
		color: #26a69a !important;
	}
	.minorBtn {
		color: grey !important;
	}
</style>