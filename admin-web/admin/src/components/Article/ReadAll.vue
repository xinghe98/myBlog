<template>
	<el-empty v-if="nodata" description="这里是空的" />
	<el-skeleton style="width: 240px" :loading="loading" animated>
		<el-row>
			<el-space wrap size="large">
				<el-card
					v-for="i in data"
					:key="i.ID"
					class="box-card"
					style="width: 250px"
					:body-style="{ padding: '5px' }"
				>
					<img :src="`${i.image}`" class="image" />
					<div style="padding: 10px">
						<span>{{ i.title }}</span>
						<div class="bottom">
							<time class="time">{{ i.content.substring(1, 30) + "..." }}</time>
							<el-button type="primary" :icon="Edit" @click="edit(i)" circle />
							<el-popconfirm
								@confirm="del(i)"
								:icon="CloseBold"
								title="确定吗？"
								icon-color="red"
								confirm-button-text="是"
								cancel-button-text="否"
								confirm-button-type="danger"
							>
								<template #reference>
									<el-button type="danger" :icon="Delete" circle />
								</template>
							</el-popconfirm>
						</div>
					</div>
				</el-card>
			</el-space>
			<el-pagination
				v-if="total > 10 && !nodata"
				style="margin-top: 30px; margin-left: auto"
				background
				layout="prev, pager, next"
				:total="total"
				:page-size="10"
				@current-change="handleCurrentChange"
			/>
		</el-row>
	</el-skeleton>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { CloseBold } from "@element-plus/icons-vue";
import { Edit, Delete } from "@element-plus/icons-vue";
import request from "@/util/request";
import { articleStore } from "@/store/articleStore";
import router from "@/router";

const loading = ref(true);
const artStore = articleStore();
const nodata = ref(false);
const total = ref(0);
const data = ref<artdata[]>([]);
const props = defineProps(["dataurl"]);

type artdata = {
	ID: number;
	title: string;
	content: string;
	Tags: never[];
	status: number;
	image: string;
	headimg: string;
};

const edit = async (article: artdata) => {
	// 将article 传递给articleStore
	artStore.ID = article.ID;
	artStore.title = article.title;
	artStore.content = article.content;
	artStore.tags = article.Tags;
	artStore.status = article.status;
	artStore.img = article.image;
	artStore.headimg = article.headimg;
	router.push(`/admin/edit/${article.ID}`);
};

const del = async (article: artdata) => {
	const res = await request.delete(`/article/${article.ID}`);
	if (res.data.code === 200) {
		data.value = data.value.filter((i) => i.ID !== article.ID);
	}
};

const handleCurrentChange = async (page: number) => {
	const res: any = await request.get(props.dataurl.replace("page=1", `page=${page}`));
	data.value = res.data.data.articles;
};

onMounted(async () => {
	const res = await request.get(props.dataurl);
	const allstatusData: artdata[] = res.data.data.articles;
	// 过滤出status为1的数据
	data.value = allstatusData;
	total.value = res.data.data.total;
	loading.value = false;
	// 判断数据是否为空
	if (data.value.length === 0) {
		nodata.value = true;
	}
});
</script>

<style scoped>
.time {
	font-size: 12px;
	color: #999;
}

.bottom {
	margin-top: 13px;
	line-height: 12px;
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.button {
	padding: 0;
	min-height: auto;
}

.image {
	width: 100%;
	height: 150px;
	display: block;
}
</style>
