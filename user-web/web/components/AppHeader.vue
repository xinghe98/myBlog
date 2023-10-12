<template>
	<div id="header">
		<el-menu
			:default-active="route.path"
			:router="true"
			class="el-menu-demo"
			mode="horizontal"
			active-text-color="#95d475"
			:ellipsis="false"
			@select="handleSelect"
		>
			<img src="https://vuejs.org/images/logo.png" alt="Vue logo" />
			<div class="flex-grow" />
			<el-menu-item index="/">主页</el-menu-item>
			<el-menu-item index="/post">文章</el-menu-item>
		</el-menu>
	</div>
</template>
<script lang="ts" setup>
import { useRoute } from "vue-router";
const route = useRoute();
const handleSelect = (key: string, keyPath: string[]) => {
	console.log(key, keyPath);
};
</script>

<style scoped>
.flex-grow {
	flex-grow: 1;
	background-color: transparent;
}
.el-menu-demo {
	background-color: transparent;
}
#header img {
	height: 40px;
	margin-top: 10px;
}
#header {
	height: 100%;
	width: 100%;
	margin: 0;
	padding: 0;
}
ul {
	height: 100%;
	width: 100%;
}
.el-menu--horizontal.el-menu {
	border-bottom: 0;
}
</style>
