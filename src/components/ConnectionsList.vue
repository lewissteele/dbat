<script setup lang="ts">
import { useDatabaseStore } from "../stores";
import { storeToRefs } from "pinia";

const db = useDatabaseStore();
const { connections } = storeToRefs(db);
</script>

<template>
  <li v-for="conn in connections">
    {{ conn.host }} {{ conn.user }} {{ conn.password }} {{ conn.port }}
  </li>
</template>

<style scoped>
</style>
