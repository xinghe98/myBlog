<!-- 读取所有已保存的文章 -->
<template>
	<readall :dataurl="url" />
</template>

<script lang="ts" setup>
import { defineAsyncComponent } from "vue";

const readall = defineAsyncComponent(() => import("@/components/Article/ReadAll.vue"));
const url = "/article/findall?status=-1";
</script>
