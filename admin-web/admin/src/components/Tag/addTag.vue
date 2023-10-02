<template>
	<el-button class="btn" type="primary" @click="dialogFormVisible = true">添加</el-button>
	<el-dialog v-model="dialogFormVisible" title="添加标签">
		<el-form :model="form">
			<el-form-item label="标签名" label-width="140px">
				<el-input v-model="form.name" autocomplete="off" />
			</el-form-item>
		</el-form>
		<template #footer>
			<span class="dialog-footer">
				<el-button @click="dialogFormVisible = false">取消</el-button>
				<el-button type="primary" @click="addTag(form)"> 确认 </el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script lang="ts" setup>
import request from "@/util/request";
import { ElMessage } from "element-plus";
import { reactive, ref, inject } from "vue";
const reload: any = inject("reload");

const dialogFormVisible = ref(false);
const form = reactive({
	name: "",
});

const addTag = async (form: { name: string }) => {
	try {
		const res = await request.post("/tags/create", form);
		ElMessage.success(res.data.msg);
		dialogFormVisible.value = false;
		reload();
	} catch (e: any) {
		ElMessage.error(e.response.data.msg);
	}
};
</script>

<style scoped>
.dialog-footer button:first-child {
	margin-right: 10px;
}
.el-input {
	width: 300px;
}
.btn {
	width: 380px;
	margin: 10px auto;
	padding: auto;
}
</style>
