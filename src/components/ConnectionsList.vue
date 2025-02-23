<script setup lang="ts">
import { Connection } from "../types";
import { NButton } from "naive-ui";
import { storeToRefs } from "pinia";
import { useConfigStore, useDatabaseStore } from "../stores";
import router from "../router";

const config = useConfigStore();
const db = useDatabaseStore();

const { connections } = storeToRefs(config);

async function handle(connection: Connection): Promise<void> {
  config.activeConnection = connection;
  await db.connect(config.activeConnection);
  router.replace({ name: "main" });
}
</script>

<template>
  <li v-for="connection in connections">
    <n-button @click="handle(connection)">
      {{ connection.user }}@{{ connection.host }}
    </n-button>
  </li>
</template>

<style scoped>
</style>
