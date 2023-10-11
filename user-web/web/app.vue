<template>
	<el-config-provider :locale="locale">
		<div class="common-layout">
			<el-container>
				<el-container>
					<el-header style="height: 40px,marin:10px">
						<AppHeader />
					</el-header>
					<NuxtPage />
				</el-container>
			</el-container>
			<el-footer>Footer</el-footer>
		</div>
	</el-config-provider>
</template>
<script>
import { defineComponent } from "vue";
import { ElConfigProvider } from "element-plus";

import zhCn from "element-plus/dist/locale/zh-cn.mjs";

export default defineComponent({
	components: {
		ElConfigProvider,
	},
	setup() {
		return {
			locale: zhCn,
		};
	},
});
</script>

<style>
html body {
	height: 100%;
	width: 100%;
	margin: 0;
	padding: 0;
}
.el-header {
	margin: 0;
	padding: 0;
}
</style>
