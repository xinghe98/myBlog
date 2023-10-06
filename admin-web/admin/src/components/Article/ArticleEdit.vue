<template>
	<div id="vditor"></div>
	<div class="but">
		<el-popconfirm @confirm="submitarticle" :icon="WarningFilled" title="确定吗？">
			<template #reference>
				<el-button type="primary">发布</el-button>
			</template>
		</el-popconfirm>
		<el-button type="primary" @click="savearticle">保存</el-button>
	</div>
</template>
<script setup lang="ts">
import { ref, onMounted } from "vue";
import Vditor from "vditor";
import "vditor/dist/index.css";
import { articleStore } from "@/store/articleStore";
import { WarningFilled } from "@element-plus/icons-vue";
import router from "@/router";
const article = articleStore();
const submitarticle = () => {
	article.content = vditor.value!.getValue();
	article.status = 1;
	if (article.ID === 0) {
		article.createArticle();
	} else {
		article.updateArticle(article.ID);
	}
	router.push("/admin/readall");
	// console.log(vditor.value!.getValue());
};

const savearticle = () => {
	article.content = vditor.value!.getValue();
	article.status = -1;
	article.createArticle();
	router.push("/admin/readsaved");
	// console.log(vditor.value!.getValue());
};

const vditor = ref<Vditor | null>(null);
onMounted(() => {
	vditor.value = new Vditor("vditor", {
		height: 800,
		toolbarConfig: {
			pin: true,
		},
		icon: "material",
		after: () => {
			// vditor.value is a instance of Vditor now and thus can be safely used here
			vditor.value!.setValue(article.content);
		},
	});
});
</script>

<style>
.but .el-button {
	float: right;
	margin-right: 10px;
	margin-top: 10px;
}
</style>
