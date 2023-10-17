<template>
	<main class="main">
		<div class="container">
			<div class="left">
				<el-card style="display: flex" shadow="hover" v-for="item in postData" :key="item.ID">
					<el-image
						fit="cover"
						style="width: 179px; height: 110px; display: flex; border-radius: 5px"
						:src="`${item.image}`"
					></el-image>
					<div class="cardofinfo">
						<div class="title">
							<a :href="`/article/${item.ID}`" target="_blank">
								{{ item.title }}
							</a>
						</div>
						<div class="content">
							{{ item.content.slice(0, 100) }}
						</div>
						<div class="data">
							<span class="tags">
								<el-tag
									style="margin: 3px; display: flex"
									v-for="tag in item.Tags"
									type="success"
									:key="tag.ID"
									effect="light"
									round
								>
									{{ tag.name }}
								</el-tag>
							</span>
							<span class="date">
								<el-icon :size="20" style="display: inline-block; vertical-align: -4px">
									<Edit />
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
			</div>
			<div class="right">
				<RightMain />
			</div>
		</div>
	</main>
</template>

<script lang="ts" setup>
import { Edit } from "@element-plus/icons-vue";
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
let dataurl = "/article/findall?status=1&page=1&limit=10";
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
.main {
	display: flex;
	justify-content: center;
	width: 100%;
	margin-top: 15px;
}
.left {
	width: calc(100% - 400px);
}
.el-card {
	margin: 10px;
	display: flex;
}
.right {
	width: 380px;
}
* {
	box-sizing: border-box;
}
div {
	display: block;
}
.container {
	display: flex;
	justify-content: space-between;
	width: 1300px;
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

.el-card ::v-deep(.el-card__body) {
	display: flex;
}

.el-card ::v-deep(.el-card__body) .cardofinfo {
	padding-left: 20px;
	color: (78, 89, 105);
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	width: 520px;
}

.el-card ::v-deep(.el-card__body) .cardofinfo .title {
	font-size: 20px;
	font-weight: 600;
}
.el-card ::v-deep(.el-card__body) .cardofinfo .content {
	margin-top: 10px;
}
.data {
	display: flex;
	align-items: center;
	margin-top: 5px;
	justify-content: space-between;
}

.tags {
	display: flex;
	flex-wrap: wrap;
}

.date {
	color: #909399;
	font-size: 14px;
	margin-right: 15px;
}
</style>
