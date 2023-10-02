<template>
	<tagOp />
</template>

<script lang="ts" setup>
import { defineAsyncComponent } from "vue";
const tagOp = defineAsyncComponent(() => import("@/components/Tag/TagsOp.vue"));
</script>
