<script setup lang="ts">
import SqlEditor from "../components/SqlEditor.vue";
import router from "../router";
</script>

<template>
  <button
    @click='() => router.replace({ name: "browser" })'
    class="btn btn-primary col"
  >
    Browser
  </button>
  <SqlEditor />
</template>
