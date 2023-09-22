<template>
	<div id="vditor"></div>
	<div class="but">
		<el-button type="primary" @click="submitarticle">发布</el-button>
	</div>
</template>
<script setup lang="ts">
import { ref, onMounted } from "vue";
import Vditor from "vditor";
import "vditor/dist/index.css";
const submitarticle = () => {
	console.log(vditor.value!.getValue());
};

const vditor = ref<Vditor | null>(null);
onMounted(() => {
	vditor.value = new Vditor("vditor", {
		height: 800,
		toolbarConfig: {
			pin: true,
		},
		after: () => {
			// vditor.value is a instance of Vditor now and thus can be safely used here
			vditor.value!.setValue("## hello world");
		},
	});
});
</script>

<style>
.but .el-button {
	float: right;
	margin-right: 10px;
	margin-top: 10px;
}
</style>
