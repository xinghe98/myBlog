<template>
	<div id="header">
		<el-menu
			:default-active="route.path"
			:router="true"
			class="el-menu-demo"
			mode="horizontal"
			:ellipsis="false"
			@select="handleSelect"
		>
			<el-menu-item index="0">LOGO</el-menu-item>
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
}
#header {
	height: 100%;
	width: 100%;
}
ul {
	height: 100%;
	width: 100%;
}
</style>
