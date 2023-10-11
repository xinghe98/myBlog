<template>
	<el-main style="display: flex; justify-content: center">
		<el-container>
			<div class="left">
				<el-card shadow="hover" v-for="item in 3" :key="item">
					<el-image
						fit="cover"
						style="width: 179px; height: 110px"
						src="https://i.pinimg.com/564x/06/31/3d/06313d8db718ae9d82503d09caf5cee9.jpg"
					></el-image>
				</el-card>
			</div>
			<div class="right">
				<RightMain />
			</div>
		</el-container>
	</el-main>
</template>

<script lang="ts" setup></script>

<style scoped>
.el-container {
	display: flex;
	justify-content: space-between;
	width: 1200px;
}
.left {
	width: calc(100% - 400px);
}
.el-card {
	margin: 10px;
	padding: 0;
}
.right {
	width: 380px;
}
* {
	box-sizing: border-box;
}
div {
	display: block;
}
</style>
