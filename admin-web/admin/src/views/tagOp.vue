<template>
	<el-table :data="tableData" class="table">
		<el-table-column label="分类名称" width="180">
			<template #default="scope">
				<el-popover effect="light" trigger="hover" placement="top" width="auto">
					<template #default>
						<div>name: {{ scope.row.name }}</div>
						<div>address: {{ scope.row.address }}</div>
					</template>
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
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import request from "@/util/request";

interface tags {
	name: string;
	ID: number;
}

const tableData = ref<tags[]>([]);

const handleEdit = (row: tags) => {
	console.log(row);
};
const handleDelete = (row: tags) => {
	console.log(row);
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
</style>
