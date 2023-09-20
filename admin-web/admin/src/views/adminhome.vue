<template>
	<el-container class="layout-container-demo">
		<el-aside width="200px">
			<el-scrollbar>
				<el-menu :default-openeds="['1']" :router="true">
					<el-sub-menu index="1">
						<template #title>
							<el-icon><message /></el-icon>文章管理
						</template>
						<el-menu-item-group>
							<template #title>文章</template>
							<el-menu-item index="/admin/readall">所有文章</el-menu-item>
							<el-menu-item index="/admin/create">新增文章</el-menu-item>
						</el-menu-item-group>
					</el-sub-menu>
					<el-sub-menu index="2">
						<template #title>
							<el-icon><icon-menu /></el-icon>博客管理
						</template>
						<el-menu-item-group>
							<template #title>标题</template>
							<el-menu-item index="2-1">Option 1</el-menu-item>
							<el-menu-item index="2-2">Option 2</el-menu-item>
						</el-menu-item-group>
					</el-sub-menu>
				</el-menu>
			</el-scrollbar>
		</el-aside>

		<el-container>
			<el-header style="text-align: right; font-size: 12px">
				<div class="toolbar">
					<el-dropdown @command="handleCommand">
						<el-icon style="margin-right: 8px; margin-top: 1px"><setting /></el-icon>
						<template #dropdown>
							<el-dropdown-menu>
								<el-dropdown-item command="logout">登出</el-dropdown-item>
							</el-dropdown-menu>
						</template>
					</el-dropdown>
					<span>Tom</span>
				</div>
			</el-header>

			<router-view />
		</el-container>
	</el-container>
</template>

<script lang="ts" setup>
import { Menu as IconMenu, Message, Setting } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { adminStore } from "@/store/adminStore";
import { useRouter } from "vue-router";
const admin = adminStore();
const router = useRouter();
const handleCommand = (command: number) => {
	ElMessage.warning(`you are ${command}`);
	admin.isAuthenticated = false;
	localStorage.removeItem("adminToken");
	router.push("/");
};
</script>

<style scoped>
html,
body {
	padding: 0%;
	margin: 0%;
}
.layout-container-demo .el-header {
	position: relative;
	background-color: var(--el-color-primary-light-7);
	color: var(--el-text-color-primary);
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
}
</style>
