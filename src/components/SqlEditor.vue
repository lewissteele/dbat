<script setup lang="ts">
import CodeMirror from "vue-codemirror6";
import { ref } from "vue";
import { MySQL, sql } from "@codemirror/lang-sql";
import { NButton } from "naive-ui";
import { useDatabaseStore } from "../stores";
import Table from "./Table.vue";

const code = ref("");
const tableData = ref([] as any[]);

const lang = sql({
  dialect: MySQL,
});

const db = useDatabaseStore();

async function handle(): Promise<void> {
  const result = await db.connection?.select(code.value) as any[];
  tableData.value = result;
}
</script>

<template>
  <code-mirror
    v-model="code"
    :lang="lang"
    :dark="true"
    :tab-size="2"
    basic
    wrap
    tab
  />
  <n-button @click="handle">Run</n-button>

  <Table :data="tableData"></Table>
</template>

<style>
.cm-editor {
  height: 100px;
}
</style>
