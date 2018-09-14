<template>
	<div>
    <v-card-title class="headline"> {{ hostname || server }}</v-card-title>
    <v-card-text> {{ server }} </v-card-text>
   </div>
</template>

<script>
	export default {
		name: 'serverInfo',
		props: ['hostname', 'server'],
	}
</script>

<style>
</style>