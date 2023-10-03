<template>
	<el-table :data="tableData" class="table">
		<el-table-column label="分类名称" width="180">
			<template #default="scope">
				<el-popover effect="light" trigger="hover" placement="top" width="auto">
					<template #reference>
						<el-tag>{{ scope.row.name }}</el-tag>
					</template>
				</el-popover>
			</template>
		</el-table-column>
		<el-table-column label="操作">
			<template #default="scope">
				<el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
				<el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
			</template>
		</el-table-column>
	</el-table>
	<el-button class="btn" type="primary" @click="handleAdd">添加</el-button>
	<!--INFO: 使用pinia 将addtag和edittag 复用为1个-->
	<el-dialog v-model="dialogFormVisible" :title="title">
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
import { onMounted, ref, inject, reactive } from "vue";
import request from "@/util/request";
import { ElMessage } from "element-plus";
import { tagsStore } from "@/store/tagsStore";

const tags = tagsStore();
const reload: any = inject("reload");
const dialogFormVisible = ref(tags.isShow);
const title = ref(tags.windowName);
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
interface tags {
	name: string;
	ID: number;
}

const tableData = ref<tags[]>([]);

const handleEdit = (row: tags) => {
	tags.editTag("编辑分类", row.ID);
	dialogFormVisible.value = true;
	title.value = tags.windowName;
};
const handleDelete = async (row: tags) => {
	try {
		const res = await request.delete(`/tags/${row.ID}`);
		ElMessage.success(res.data.data.msg);
		reload();
	} catch (e: any) {
		ElMessage.error(e.response.data.msg);
	}
};

const handleAdd = () => {
	tags.addTag("添加分类");
	dialogFormVisible.value = true;
	title.value = tags.windowName;
};
onMounted(async () => {
	const res = await request.get("/tags/findall");
	tableData.value = res.data.data;
});
</script>

<style scoped>
.table {
	width: 380px;
	margin: 0 auto;
	padding: auto;
}
.btn {
	width: 380px;
	margin: 10px auto;
	padding: auto;
}
.dialog-footer button:first-child {
	margin-right: 10px;
}
.el-input {
	width: 300px;
}
</style>
