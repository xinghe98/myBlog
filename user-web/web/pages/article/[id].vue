<template>
	<main>
		<el-row :gutter="10" justify="center">
			<el-col :sm="24" :md="4" :lg="4">
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
			</el-col>
			<el-col :sm="24" :md="10" :lg="10">
				<div class="article">
					<MdPreview
						:preview-theme="preview_theme"
						:code-theme="code_theme"
						:editor-id="md_moudle"
						:model-value="article.content"
						style="
							border-radius: 5px 5px 8px 8px;
							box-shadow: var(--card-box-shadow) !important;
							background: rgba(255, 255, 255, 0.3);
							transition: all 0.3s;
							padding: 20px;
						"
					/>
				</div>
			</el-col>
		</el-row>
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
.el-card ::v-deep(.el-card__body) {
	max-height: 80vh;
	overflow: auto;
	padding: 20px;
	font-size: 0.8rem;
}
.el-card {
	position: fixed;
	width: 13vw;
	background-color: rgba(255, 255, 255, 0.5);
	border: none;
	border-radius: 5px 5px 8px 8px;
	box-shadow: var(--card-box-shadow) !important;
	transition: all 0.3s;
}
</style>
