<template>
	<el-card shadow="hover" v-for="item in postData" :key="item.ID">
		<el-image fit="cover" style="border-radius: 5px" :src="`${item.image}`"></el-image>
		<div class="cardofinfo">
			<div class="title">
				<a :href="`/article/${item.ID}`" target="_blank">
					{{ item.title }}
				</a>
			</div>
			<div class="content">
				{{ item.content.slice(0, 100).replace(/[#,\-,\+,>,!,\[\],```]/g, " ") }}
			</div>
			<div class="data">
				<span class="tags">
					<el-tag
						style="margin: 3px; display: flex"
						v-for="tag in item.Tags"
						:key="tag.ID"
						effect="light"
						round
					>
						{{ tag.name }}
					</el-tag>
				</span>
				<span class="date">
					<el-icon :size="20" style="display: inline-block; vertical-align: -4px">
						<EditPen />
					</el-icon>
					{{ item.UpdatedAt.slice(0, 10) }}
				</span>
			</div>
		</div>
	</el-card>
	<div class="page">
		<el-pagination
			v-if="total > 10"
			background
			layout="prev, pager, next"
			:total="total"
			:page-size="10"
			@current-change="handleCurrentChange"
		/>
	</div>
</template>

<script lang="ts" setup>
import { EditPen } from "@element-plus/icons-vue";
import { request, articlesData } from "~/util/requests";

type tagdata = {
	ID: number;
	name: string;
};
type artdata = {
	ID: number;
	title: string;
	content: string;
	Tags: tagdata[];
	status: number;
	image: string;
	headimg: string;
	UpdatedAt: string;
};
const route = useRoute();
const name = route.params.tag;
let dataurl = `/tags/${name}?page=1&limit=10`;
if (name === undefined) {
	dataurl = "/article/findall?status=1&page=1&limit=10";
}

const total = ref(0);
const postData = ref<artdata[]>([]);

const handleCurrentChange = async (page: number) => {
	const { data } = await request<artdata[]>(dataurl.replace("page=1", `page=${page}`));
	postData.value = (data.value!.data as articlesData<artdata[]>).articles;
};

const { data } = await request<artdata[]>(dataurl);
const allstatusData: artdata[] = (data.value!.data as articlesData<artdata[]>).articles;
postData.value = allstatusData;
total.value = (data.value!.data as articlesData<artdata[]>).total;
</script>

<style scoped>
a {
	color: #000;
	text-decoration: none;
}
* {
	box-sizing: border-box;
}
div {
	display: block;
}
.page {
	display: flex;
	justify-content: center;
	margin: 24px;
}
.el-pagination {
	display: flex;
	align-items: center;
}
.el-card {
	margin: 0.8vh;
	background-color: rgba(255, 255, 255, 0.6);
}
.el-card .el-image {
	width: 180px;
	height: 110px;
	border-radius: 5px;
}
.el-card ::v-deep(.el-card__body) {
	display: grid;
	grid-template-columns: 2fr 22fr;
	grid-column-gap: 20px;
	align-items: center;
}

.el-card ::v-deep(.el-card__body) .cardofinfo {
	display: grid;
	align-items: center;
	grid-row-gap: 10px;
}

.el-card ::v-deep(.el-card__body) .cardofinfo .title {
	font-size: 1.5rem;
	font-weight: 600;
}
.el-card ::v-deep(.el-card__body) .cardofinfo .data {
	display: grid;
	grid-template-columns: 2fr 22fr;
	justify-items: end;
}
.data .tags {
	display: grid;
	grid-template-columns: 2fr 22fr;
	grid-auto-flow: column;
}
</style>
