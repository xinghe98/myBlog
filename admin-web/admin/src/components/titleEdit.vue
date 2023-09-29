<template>
	<el-input @blur="addtitle" v-model="input" placeholder="在此处填写你的文章标题"></el-input>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { articleStore } from "@/store/articleStore";
const input = ref("");
const article = articleStore();
const addtitle = () => {
	article.title = input.value;
};
</script>

<style scoped>
.el-input {
	margin-bottom: 10px;
	margin-top: 10px;
}
</style>
