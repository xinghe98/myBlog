<template>
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
					<img
						src="https://pic.52112.com/2019/07/23/JPG-190723_759/mfvvEHO7sx_small.jpg"
						class="image"
					/>
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

type artdata = {
	ID: number;
	title: string;
	content: string;
	Tags: never[];
	status: number;
};

const edit = async (article: artdata) => {
	// 将article 传递给articleStore
	artStore.ID = article.ID;
	artStore.title = article.title;
	artStore.content = article.content;
	artStore.tags = article.Tags;
	artStore.status = article.status;
	router.push(`/admin/edit/${article.ID}`);
};

const del = async (article: artdata) => {
	const res = await request.delete(`/article/${article.ID}`);
	if (res.data.code === 200) {
		data.value = data.value.filter((i) => i.ID !== article.ID);
	}
};

const data = ref<artdata[]>([]);
const props = defineProps(["dataurl"]);
onMounted(async () => {
	const res = await request.get(props.dataurl);
	const allstatusData: artdata[] = res.data.data;
	// 过滤出status为1的数据
	data.value = allstatusData;
	loading.value = false;
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
	display: block;
}
</style>
