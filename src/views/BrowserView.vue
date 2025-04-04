<script setup lang="ts">
import TableComponent from "../components/Table.vue";
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
  <div class="row flex-nowrap">
    <ul class="list-group list-group-flush col-2">
      <a
        :class='{ "list-group-item-light": table.name == selectedTable }'
        @click="async () => await selectTable(table)"
        class="list-group-item list-group-item-action"
        href="#"
        v-for="table in tables"
      >
        {{ table.name }}
      </a>
    </ul>
    <TableComponent :data="tableData" class="col-10" />
  </div>
</template>

<style lang="scss" scoped>
</style>
