<template>
	<el-form-item label="标题:">
		<el-input @blur="addtitle" v-model="input" placeholder="在此处填写你的文章标题"></el-input>
	</el-form-item>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { articleStore } from "@/store/articleStore";
const input = ref("");
const article = articleStore();
const addtitle = () => {
	article.title = input.value;
};
onMounted(() => {
	input.value = article.title;
});
</script>

<style scoped>
.el-input {
	margin-bottom: 10px;
	padding-right: 10px;
}
.el-form-item {
	margin-top: 10px;
}
</style>
