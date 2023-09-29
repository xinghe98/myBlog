<template>
	<div class="demo-collapse">
		<el-collapse accordion @change="preview">
			<el-collapse-item v-for="data in datas as any[]" :name="data.ID">
				<template #title> {{ data.title }} </template>
				<div id="preview"></div>
			</el-collapse-item>
		</el-collapse>
	</div>
</template>

<script setup lang="ts">
import Vditor from "vditor";
import { onMounted, ref } from "vue";
import request from "@/util/request";
import { CollapseModelValue } from "element-plus";
let datas = ref<any>([]);
// FIX: 无法预览
const preview = (id: CollapseModelValue) => {
	const data = datas.value.find((item: any) => item.ID === id);
	if (data === undefined) return;
	console.log(data.content);
	// Vditor.preview(document.getElementById("preview") as HTMLDivElement, data.content);
};
onMounted(async () => {
	const res = await request.get("/article/findall");
	datas.value = res.data.data;
	Vditor.preview(document.getElementById("preview") as HTMLDivElement, "## da");
});
</script>
