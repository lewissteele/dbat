<script setup lang="ts">
import { useDatabaseStore } from "../stores";
import { storeToRefs } from "pinia";
import { NButton } from "naive-ui";
import Connection from "../types/connection";

const db = useDatabaseStore();
const { saved } = storeToRefs(db);

function handle(conn: Connection): void {
  db.setActive(conn);
}
</script>

<template>
  <li v-for="conn in saved">
    <n-button @click="handle(conn)">{{ conn.user }}@{{ conn.host }}</n-button>
  </li>
</template>

<style scoped>
</style>
