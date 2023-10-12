<template>
	<el-button type="primary" @click="isShowDialog = true" round>设置为头条</el-button>
	<el-dialog v-model="isShowDialog" title="上传一张封面"></el-dialog>
</template>

<script setup lang="ts">
// FIX: 这里没写完，差一个上传图片的组件
import { ElButton } from "element-plus";
import { ref } from "vue";

const isShowDialog = ref(false);
</script>

<style scoped>
.el-button {
	background-color: #67c23a;
	border-color: #67c23a;
	display: block;
	margin: 10px;
}
</style>
