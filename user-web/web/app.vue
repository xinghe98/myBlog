<template>
	<el-config-provider :locale="locale">
		<div style="display: block">
			<el-container>
				<el-header style="height: 55px; marin: 10px; padding: 0">
					<el-affix>
						<AppHeader />
					</el-affix>
				</el-header>
				<el-main style="padding: 0">
					<NuxtPage />
				</el-main>
				<el-footer style="padding: 0">
					<AppFooter />
				</el-footer>
			</el-container>
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
body {
	background-color: #f0eeee;
}
</style>
