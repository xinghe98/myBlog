<template>
	<div class="block">
		<el-card class="box-card">
			<template #header>
				<div class="card-header">
					<span>所有标签</span>
					<!--
					<el-button class="button" text>Operation button</el-button>
					-->
				</div>
			</template>
			<div v-for="tag in tags" :key="tag.ID" class="text item">
				<el-button @click="handleTags(tag.name)" size="small" type="info" plain>{{
					tag.name
				}}</el-button>
			</div>
		</el-card>
	</div>
</template>

<script lang="ts" setup>
import { request } from "~/util/requests";
import { useTagsStore } from "~/stores/tagsStore";

type tagdata = {
	ID: number;
	name: string;
};

const tags = ref<tagdata[]>([]);
const tagsStore = useTagsStore();

const { data } = await request<tagdata[]>("/tags/findall");
tags.value = data.value!.data as tagdata[];

const handleTags = (name: string) => {
	tagsStore.tags = name;
};
</script>

<style scoped>
.block {
	margin: 10px;
}
div {
	display: block;
}
.card-header {
	display: flex;
	justify-content: center;
	align-items: center;
}

.text {
	font-size: 14px;
}

.item {
	margin-bottom: 18px;
	margin: 10px;
}
.el-card ::v-deep(.el-card__body) {
	display: flex;
	flex-wrap: wrap;
	justify-content: space-between;
}
</style>
