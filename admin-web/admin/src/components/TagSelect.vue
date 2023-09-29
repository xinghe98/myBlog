<template>
	<div><span>选择文章标签:</span></div>
	<div>
		<el-checkbox-group v-model="checkedTags" @change="selectTags">
			<el-checkbox v-for="tag in tags" :key="tag" :label="tag">{{ tag }}</el-checkbox>
		</el-checkbox-group>
	</div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import request from "@/util/request";
import { articleStore } from "@/store/articleStore";
import { CheckboxValueType } from "element-plus";
const checkedTags = ref([]);
const tags = ref([]);
const article = articleStore();
const selectTags = (value: CheckboxValueType[]) => {
	//将数组转为[{name:tag1},{name:tag2}]的形式
	let tags = value.map((item) => {
		return { name: item };
	});
	article.tags = tags as never[];
};
onMounted(async () => {
	const res = await request.get("/tags/findall");
	// 对象数字转数组
	tags.value = res.data.data.map((item: any) => item.name);
});
</script>

<style scoped>
.el-checkbox-group {
	margin-bottom: 10px;
	margin-top: 10px;
}
div {
	background-color: #f2f2f2;
}
</style>
