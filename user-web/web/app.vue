<template>
	<div class="common-layout">
		<el-container>
			<el-container>
				<el-header>
					<AppHeader />
				</el-header>
				<NuxtPage />
			</el-container>
		</el-container>
		<el-footer>Footer</el-footer>
	</div>
</template>
<style scoped>
.el-header {
	height: 50px;
	margin-top: 7px;
}
</style>
