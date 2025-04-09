<script setup lang="ts">
import router from "../router";
import { Connection } from "../types";
import { storeToRefs } from "pinia";
import { useConfigStore, useDatabaseStore } from "../stores";

const config = useConfigStore();
const db = useDatabaseStore();

const { connections } = storeToRefs(config);

async function selectConnection(connection: Connection): Promise<void> {
  config.activeConnection = connection;
  await db.connect(config.activeConnection);
  router.replace({ name: "browser" });
}
</script>

<template>
  <div class="container pt-5">
    <ul class="list-group col-12">
      <a
        @click="async () => await selectConnection(connection)"
        class="list-group-item list-group-item-action"
        href="#"
        v-for="connection in connections"
      >
        {{ connection.user }}@{{ connection.host }}
      </a>
    </ul>
  </div>
</template>

<style scoped>
</style>
