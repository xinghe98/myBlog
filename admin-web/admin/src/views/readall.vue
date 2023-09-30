<template>
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
					src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
					class="image"
				/>
				<div style="padding: 10px">
					<span>{{ i.title }}</span>
					<div class="bottom">
						<time class="time">{{ i.content.substring(1, 30) + "..." }}</time>
						<el-button type="primary" :icon="Edit" circle />
						<el-button type="danger" :icon="Delete" circle />
					</div>
				</div>
			</el-card>
		</el-space>
	</el-row>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Edit, Delete } from "@element-plus/icons-vue";
import request from "@/util/request";
type artdata = {
	ID: number;
	title: string;
	content: string;
	tags: string[];
};
const data = ref<artdata[]>([]);
onMounted(async () => {
	const res = await request.get("/article/findall");
	data.value = res.data.data;
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
