<template>
  <RouterView v-if="isRouterActive" />
</template>

<script lang="ts" setup>
import { ref, provide, nextTick } from "vue";
const isRouterActive = ref(true);
provide("reload", () => {
  isRouterActive.value = false;
  nextTick(() => {
    isRouterActive.value = true;
  });
});
</script>

<style>
html,
body {
  padding: 0%;
  margin: 0%;
}
</style>
