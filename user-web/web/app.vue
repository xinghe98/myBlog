<template>
	<el-config-provider :locale="locale">
		<div class="layout" style="display: block">
			<el-container>
				<el-header style="height: 55px; marin: 10px; padding: 0">
					<el-affix>
						<AppHeader />
					</el-affix>
				</el-header>
				<el-main style="padding: 0">
					<NuxtPage />
				</el-main>
				<el-footer style="padding: 0; overflow: hidden">
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
.layout {
	background: linear-gradient(
		90deg,
		#ffffff 1%,
		#ead2ea 10.2%,
		#e4d5ee 19.6%,
		#d1def3 36.8%,
		#b5dee5 62.2%,
		#cfebf3 88.9%,
		#ffffff 99%
	) !important;
}
body {
	background-color: #f0eeee;
}
.el-container {
	min-height: 100vh;
}
</style>
