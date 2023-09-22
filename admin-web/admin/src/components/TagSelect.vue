<template>
	<div>
		<el-checkbox v-model="checked1" label="Option1" size="large" border />
		<el-checkbox v-model="checked2" label="Option2" size="large" border />
	</div>
</template>

<script lang="ts" setup>
import { ref } from "vue";

const checked1 = ref(true);
const checked2 = ref(false);
const checkboxGroup1 = ref(["Option1"]);
</script>
