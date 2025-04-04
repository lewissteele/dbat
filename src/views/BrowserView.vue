<script setup lang="ts">
import { Table } from "../types";
import { ref } from "vue";
import { storeToRefs } from "pinia";
import { useDatabaseStore } from "../stores";

const db = useDatabaseStore();
const { tables } = storeToRefs(db);
const selectedTable = ref("");

function selectTable(table: Table): void {
  selectedTable.value = table.name;
}
</script>

<template>
  <ul class="list-group list-group-flush">
    <a
      :class='{ "list-group-item-light": table.name == selectedTable }'
      @click="selectTable(table)"
      class="list-group-item list-group-item-action"
      href="#"
      v-for="table in tables"
    >
      {{ table.name }}
    </a>
  </ul>
</template>

<style lang="scss" scoped>
</style>
