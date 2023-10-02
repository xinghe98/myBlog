<script lang="ts" setup>
import { defineAsyncComponent } from "vue";

const titleEdit = defineAsyncComponent(() => import("@/components/Article/titleEdit.vue"));
const TagSelect = defineAsyncComponent(() => import("@/components/Article/TagSelect.vue"));
const articleedit = defineAsyncComponent(() => import("@/components/Article/ArticleEdit.vue"));
</script>

<template>
	<titleEdit />
	<TagSelect />
	<articleedit />
</template>
