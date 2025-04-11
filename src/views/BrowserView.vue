<script setup lang="ts">
import TableComponent from "../components/Table.vue";
import router from "../router";
import { Table } from "../types";
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useDatabaseStore } from "../stores";

const db = useDatabaseStore();
const { tables } = storeToRefs(db);

const selectedTable = ref("");
const tableData = ref([] as any[]);

async function selectTable(table: Table): Promise<void> {
  selectedTable.value = table.name;

  tableData.value = await db.connection?.select(
    `select * from ${table.name}`,
  ) as any[];
}
</script>

<template>
  <div class="row">
    <button
      @click='() => router.replace({ name: "connections" })'
      class="btn btn-primary col"
    >
      Connections
    </button>
    <button
      @click='() => router.replace({ name: "editor" })'
      class="btn btn-primary col"
    >
      Editor
    </button>
  </div>

  <div class="row flex-nowrap">
    <ul class="list-group list-group-flush col-2">
      <button
        :class="{ active: table.name == selectedTable }"
        @click="async () => await selectTable(table)"
        class="list-group-item list-group-item-action"
        type="button"
        v-for="table in tables"
      >
        {{ table.name }}
      </button>
    </ul>
    <TableComponent :data="tableData" class="col-10" />
  </div>
</template>

<style lang="scss" scoped>
</style>
