<!-- 该视图用于显示编辑/新增文章的页面 -->
<script lang="ts" setup>
import { defineAsyncComponent } from "vue";

const titleEdit = defineAsyncComponent(() => import("@/components/Article/titleEdit.vue"));
const TagSelect = defineAsyncComponent(() => import("@/components/Article/TagSelect.vue"));
const HeaderlineSelect = defineAsyncComponent(
	() => import("@/components/Article/HeaderlineSelect.vue"),
);
const articleedit = defineAsyncComponent(() => import("@/components/Article/ArticleEdit.vue"));
</script>

<template>
	<titleEdit />
	<TagSelect />
	<HeaderlineSelect />
	<articleedit />
</template>
