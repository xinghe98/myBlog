<!-- 该视图用于显示主页面的基本架构 -->
<template>
	<el-container class="layout-container-demo">
		<el-aside width="auto">
			<el-scrollbar>
				<el-menu
					:default-openeds="['1']"
					:default-active="route.path"
					:router="true"
					:collapse="isCollapse"
					class="el-menu-container"
				>
					<el-sub-menu index="1">
						<template #title>
							<el-icon><Notebook /></el-icon>
							<span>文章管理</span>
						</template>
						<el-menu-item-group>
							<template #title>文章</template>
							<el-menu-item index="/admin/readall">已发布</el-menu-item>
							<el-menu-item index="/admin/readsaved">草稿箱</el-menu-item>
							<el-menu-item index="/admin/create">新增文章</el-menu-item>
						</el-menu-item-group>
					</el-sub-menu>
					<el-sub-menu index="2">
						<template #title>
							<el-icon><Setting /></el-icon>
							<span>博客管理</span>
						</template>
						<el-menu-item-group>
							<template #title>标题</template>
							<el-menu-item index="/admin/tags">分类管理</el-menu-item>
						</el-menu-item-group>
					</el-sub-menu>
				</el-menu>
			</el-scrollbar>
		</el-aside>
		<el-container>
			<el-header>
				<el-link v-if="isCollapse" class="collpasebut" :icon="Expand" @click="transformColl" />
				<el-link v-else class="collpasebut" :icon="Fold" @click="transformColl" />
				<div class="toolbar">
					<el-dropdown @command="handleCommand">
						<el-icon style="margin-right: 8px; margin-top: 1px"><setting /></el-icon>
						<template #dropdown>
							<el-dropdown-menu>
								<el-dropdown-item command="logout">登出</el-dropdown-item>
							</el-dropdown-menu>
						</template>
					</el-dropdown>
					<span>{{ admin.username }}</span>
				</div>
			</el-header>
			<el-card class="card">
				<router-view />
			</el-card>
		</el-container>
	</el-container>
</template>

<script lang="ts" setup>
import { Notebook, Setting, Fold, Expand } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { adminStore } from "@/store/adminStore";
import { useRouter } from "vue-router";
import { ref } from "vue";
import { useRoute } from "vue-router";
const route = useRoute();
let isCollapse = ref<boolean>(false);
const admin = adminStore();
const router = useRouter();
const transformColl = () => {
	isCollapse.value = !isCollapse.value;
};
const handleCommand = (command: number) => {
	ElMessage.warning(`you are ${command}`);
	admin.isAuthenticated = false;
	localStorage.removeItem("adminToken");
	router.push("/");
};
</script>

<style lang="scss" scoped>
html,
body {
	padding: 0%;
	margin: 0%;
}
.layout-container-demo .el-header {
	position: relative;
	background-color: var(--el-color-primary-light-7);
	color: var(--el-text-color-primary);
	padding: 0;
}
.layout-container-demo .el-aside {
	background: transparent;
}
.layout-container-demo .el-menu {
	border-right: none;
}
.layout-container-demo .el-main {
	padding: 0;
}
.layout-container-demo .toolbar {
	display: inline-flex;
	align-items: center;
	justify-content: center;
	height: 100%;
	right: 20px;
	font-size: 12px;
	float: right;
}
.el-menu-container:not(.el-menu--collapse) {
	width: 200px;
	height: calc(100vh);
}
.collpasebut {
	width: 30px;
	margin-top: 11px;
	font-size: 32px;
}
.card {
	margin: 10px;
	margin-top: 10px;
	height: 100%;
	box-shadow: var(--el-box-shadow-light);
}
.el-aside {
	box-shadow: var(--el-box-shadow-lighter);
}
</style>
