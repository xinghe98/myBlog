<template>
	<el-upload
		:headers="headers"
		:action="action"
		ref="upload"
		list-type="picture-card"
		:on-exceed="handleExceed"
		:on-success="handleSuccess"
		:limit="limit"
	>
		<el-icon><Plus /></el-icon>

		<template #tip>
			<div class="el-upload__tip">只能上传1张图片,建议大小：2550*700px</div>
		</template>
		<template #file="{ file }">
			<div>
				<img class="el-upload-list__item-thumbnail" :src="file.url" alt="" />
				<span class="el-upload-list__item-actions">
					<span v-if="!disabled" class="el-upload-list__item-delete" @click="handleRemove">
						<el-icon><Delete /></el-icon>
					</span>
				</span>
			</div>
		</template>
	</el-upload>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import { Delete, Plus } from "@element-plus/icons-vue";
import { articleStore } from "@/store/articleStore";
import { ElMessage, UploadInstance, genFileId } from "element-plus";
import type { UploadProps, UploadRawFile } from "element-plus";
import request from "@/util/request";

const baseUrl = import.meta.env.VITE_BASE_URL;
const action = baseUrl + "/img/upload";
const disabled = ref(false);
const limit = ref(1);
const upload = ref<UploadInstance>();
const headers = {
	Authorization: "Bearer " + localStorage.getItem("adminToken")!,
};
const article = articleStore();
const handleRemove = async () => {
	try {
		const res = await request.delete(`/img/${article.headimg}`);
		ElMessage.success(res.data.data.msg);
	} catch (e: any) {
		if (e.response.data.data !== null) {
			const keys = Object.keys(e.response.data.data);
			ElMessage.error(e.response.data.data[keys[0]][0]);
		} else {
			ElMessage.error(e.response.data.msg);
		}
	}
	article.headimg = "";
	upload.value!.clearFiles();
};
const handleExceed: UploadProps["onExceed"] = (files) => {
	upload.value!.clearFiles();
	const file = files[0] as UploadRawFile;
	file.uid = genFileId();
	upload.value!.handleStart(file);
};
const handleSuccess: UploadProps["onSuccess"] = (res, _file) => {
	article.headimg = res.data.file_cos_uuid;
	article.isheadlines = true;
};
</script>
