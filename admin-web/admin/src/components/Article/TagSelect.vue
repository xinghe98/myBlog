<template>
	<div class="text"><span>选择文章标签:</span></div>
	<div class="tags">
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
const checkedTags = ref<any>([]);
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
	// 对象数组转数组
	tags.value = res.data.data.map((item: any) => item.name);
	// 从articleStore中获取tags
	checkedTags.value = article.tags.map((item: any) => item.name);
});
</script>

<style scoped>
.el-checkbox-group {
	margin-bottom: 10px;
	margin-top: 10px;
	padding-left: 10px;
	padding-right: 10px;
}
div {
	padding-left: 4px;
	padding-right: 5px;
	margin-top: -10px;
	margin-bottom: -10px;
}
.text {
	display: inline-block;
	font-size: 16px;
}
.tags {
	display: inline-block;
}
</style>
