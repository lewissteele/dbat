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
@import "bootstrap/scss/bootstrap";
</style>
