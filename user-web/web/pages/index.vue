<template>
	<el-container>
		<el-main>
			<el-carousel height="700px" :interval="8000" arrow="hover">
				<el-carousel-item>
					<img
						src="http://image.fengfengzhidao.com/1001/20231006164747__shouhuiyuhangyuan123.png"
						alt=""
					/>
				</el-carousel-item>
			</el-carousel>
			<SomeArticle />
		</el-main>
	</el-container>
</template>
<style scoped>
.el-main {
	padding: 0;
	margin: 0;
}
.el-carousel__item img {
	width: 100%;
	height: 100%;
	object-fit: cover;
	margin: 0;
	padding: 0;
}
</style>
