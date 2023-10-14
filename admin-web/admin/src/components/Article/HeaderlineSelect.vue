<template>
	<el-popconfirm
		width="220"
		confirm-button-text="OK"
		cancel-button-text="No, Thanks"
		:icon="WarningFilled"
		icon-color="red"
		title="上传封面的机会只有一次，确定吗？"
		@confirm="showdialog"
	>
		<template #reference>
			<el-button :disabled="isDisable" type="primary" @click="" round>设置为头条</el-button>
		</template>
	</el-popconfirm>
	<el-dialog v-model="isShowDialog" title="上传一张封面">
		<upload-img />
	</el-dialog>
</template>

<script setup lang="ts">
import { ElButton } from "element-plus";
import { defineAsyncComponent, ref } from "vue";
import { articleStore } from "@/store/articleStore";
import { WarningFilled } from "@element-plus/icons-vue";

const article = articleStore();
const isShowDialog = ref(false);
const isDisable = ref();
if (article.headimg !== "") {
	isDisable.value = true;
} else {
	isDisable.value = false;
}

const showdialog = () => {
	isShowDialog.value = true;
};
const uploadImg = defineAsyncComponent(() => import("@/components/Article/uploadimg.vue"));
</script>

<style scoped>
.el-button {
	background-color: #67c23a;
	border-color: #67c23a;
	display: block;
	margin: 10px;
}
</style>
