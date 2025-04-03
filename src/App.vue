<script setup lang="ts">
import { darkTheme, NConfigProvider } from "naive-ui";
import { RouterView } from "vue-router";
import { onMounted } from "vue";
import { getCurrentWindow } from "@tauri-apps/api/window";

onMounted(() => getCurrentWindow().show());
</script>

<template>
  <n-config-provider :theme="darkTheme">
    <RouterView></RouterView>
  </n-config-provider>
</template>

<style lang="scss">
@import "bootstrap/scss/bootstrap";
</style>
