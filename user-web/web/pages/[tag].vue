<template>
	<div class="page-header">
		<div class="route-font">
			<span>和{{ name }}相关的文章</span>
		</div>
	</div>
	<main class="main">
		<div class="container">
			<div class="left">
				<el-card shadow="hover" v-for="item in postData" :key="item.ID">
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
							{{ item.content.slice(0, 50).replace(/[#,\-,\+,>,!,\[\],```]/g, " ") }}
						</div>
						<div class="data">
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
			</div>
			<div class="right">
				<RightMain />
			</div>
		</div>
	</main>
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
const total = ref(0);
const postData = ref<artdata[]>([]);
const route = useRoute();
const name = route.params.tag;
let dataurl = `/tags/${name}?page=1&limit=10`;

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
.container {
	max-width: 50%;
	margin: 0 auto;
}
a {
	color: #000;
	text-decoration: none;
}
.el-card {
	margin: 10px;
	display: flex;
	background-color: rgba(255, 255, 255, 0.6);
	width: auto;
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

.el-card ::v-deep(.el-card__body) {
	display: flex;
	width: 100%;
}

.el-card ::v-deep(.el-card__body) .cardofinfo {
	padding-left: 20px;
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	width: 100%;
}

.el-card ::v-deep(.el-card__body) .cardofinfo .title {
	font-size: 20px;
	font-weight: 600;
}
.el-card ::v-deep(.el-card__body) .cardofinfo .title :hover {
	font-size: 20px;
	font-weight: 600;
	color: #5cbfef;
}
.el-card ::v-deep(.el-card__body) .cardofinfo .content {
	margin-top: 10px;
	flex-wrap: wrap;
	color: rgb(115, 130, 151);
}
.data {
	display: flex;
	align-items: center;
	margin-top: 5px;
	justify-content: space-between;
}

.date {
	color: #909399;
	font-size: 14px;
	margin-right: 15px;
}
.left {
	width: calc(100% - 400px);
}
.right {
	width: 380px;
}
.main {
	display: flex;
	justify-content: center;
	width: 100%;
	margin-top: 15px;
}
.container {
	display: flex;
	justify-content: space-between;
	max-width: 1300px;
}
.page-header {
	width: 100%;
	display: flex;
	background-image: url("https://mrzym.gitee.io/blogimg/cover/cute.jpg");
	display: flex;
	justify-content: center;
	align-items: center;
	position: relative;
	background-color: #8a7e90;
	background-position: center center;
	background-size: cover;
	background-repeat: no-repeat;
	height: 26rem;
}

.page-header .route-font {
	font-size: 3.2rem;
	font-weight: 500;
	line-height: 2.4;
	text-align: center;
	color: #fff;
	z-index: 999;
	cursor: pointer;
	transition: all 0.3s;
	text-shadow: 0 1px 0 hsl(174, 5%, 80%), 0 2px 0 hsl(174, 5%, 75%), 0 3px 0 hsl(174, 5%, 70%),
		0 0 5px rgba(0, 0, 0, 0.05), 0 1px 3px rgba(0, 0, 0, 0.2), 0 3px 5px rgba(0, 0, 0, 0.2),
		0 5px 10px rgba(0, 0, 0, 0.2), 0 10px 10px rgba(0, 0, 0, 0.2), 0 20px 20px rgba(0, 0, 0, 0.3);
}
</style>
