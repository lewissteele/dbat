<script setup lang="ts">
import { darkTheme, NConfigProvider } from "naive-ui";
import { RouterView } from "vue-router";
import { onMounted } from "vue";
import { getCurrentWindow } from "@tauri-apps/api/window";

onMounted(() => getCurrentWindow().show());
</script>

<template>
  <n-config-provider :theme="darkTheme">
    app
    <RouterView></RouterView>
  </n-config-provider>
</template>

<style>
:root {
  font-family: Inter, Avenir, Helvetica, Arial, sans-serif;
  font-size: 16px;
  line-height: 24px;
  font-weight: 400;

  color: #0f0f0f;
  background-color: #f6f6f6;

  font-synthesis: none;
  text-rendering: optimizeLegibility;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  -webkit-text-size-adjust: 100%;
}

@media (prefers-color-scheme: dark) {
  :root {
    color: #f6f6f6;
    background-color: #2f2f2f;
  }
}
</style>
