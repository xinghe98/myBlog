<template>
	<main>
		<div class="container">
			<div class="aside">
				<ElCard shadow="never">
					<template #header>
						<div class="card-header">
							<span style="font-size: 18px; font-weight: 600">文章目录</span>
						</div>
					</template>
					<client-only>
						<MdCatalog :editorId="md_moudle" :scrollElement="scrollElement" />
					</client-only>
				</ElCard>
			</div>
			<div class="article">
				<MdPreview
					:preview-theme="preview_theme"
					:code-theme="code_theme"
					:editor-id="md_moudle"
					:model-value="article.content"
					style="background-color: rgba(255, 255, 255, 0.8)"
				/>
			</div>
		</div>
	</main>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { articlesData, request } from "~/util/requests";
import { MdPreview, MdCatalog } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";

const route = useRoute();
const id = route.params.id;
const preview_theme = "cyanosis";
const code_theme = "atom";
const scrollElement = ref<HTMLElement>();

onMounted(() => {
	// 必须挂载dom之后才能操作dom
	scrollElement.value = document.documentElement as HTMLElement;
});

interface article {
	ID: number;
	title: string;
	content: string;
	status: number;
	image: string;
	headimg: string;
	UpdatedAt: string;
}

const md_moudle = "preview-only";

const { data } = await request<article>("/article/" + id);
const article = (data.value!.data as articlesData<article>).articles;
</script>

<style scoped>
main {
	width: 100%;
	display: flex;
	justify-content: center;
	padding-top: 20px;
	padding-bottom: 20px;
}

main .container {
	width: 1200px;
	display: flex;
	justify-content: space-between;
}
.container .article {
	width: calc(100% - 320px);
}
.container .aside {
	width: 300px;
}
.el-card ::v-deep(.el-card__body) {
	max-height: calc(100vh - 300px);
	overflow: auto;
	padding: 20px;
}
.el-card {
	position: fixed;
	width: 300px;
	background-color: rgba(255, 255, 255, 0.5);
}
</style>
