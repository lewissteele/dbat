<script setup lang="ts">
import { RouterView } from "vue-router";
import { getCurrentWindow } from "@tauri-apps/api/window";
import { onMounted } from "vue";

onMounted(() => getCurrentWindow().show());
</script>

<template>
  <RouterView></RouterView>
</template>

<style lang="scss">
$bootstrap-icons-font-dir: "bootstrap-icons/font/fonts";

@import "bootstrap/scss/bootstrap";
@import "bootstrap-icons/font/bootstrap-icons";
</style>
