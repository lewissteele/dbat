<script setup lang="ts">
import { useDatabaseStore } from '../stores';

const db = useDatabaseStore();
</script>

<template>
  <li v-for="conn in db.connections">
    {{ conn.host }}
    {{ conn.user }}
    {{ conn.password }}
    {{ conn.port }}
  </li>
</template>

<style scoped>
</style>