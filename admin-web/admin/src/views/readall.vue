<template>
	<readall />
</template>
<script lang="ts" setup>
import { defineAsyncComponent } from "vue";
const readall = defineAsyncComponent(() => import("@/components/Article/ReadAll.vue"));
</script>
