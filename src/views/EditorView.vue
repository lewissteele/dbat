<script setup lang="ts">
import CodeMirror from "vue-codemirror6";
import Table from "../components/Table.vue";
import router from "../router";
import { MySQL, sql } from "@codemirror/lang-sql";
import { ref } from "vue";
import { useDatabaseStore } from "../stores";

const code = ref("select * from migrations");
const tableData = ref([] as any[]);

const lang = sql({
  dialect: MySQL,
});

const db = useDatabaseStore();

async function handle(): Promise<void> {
  try {
    const result = await db.connection?.select(code.value) as any[];
    tableData.value = result;
  } catch (error: any) {
    tableData.value = [{ error }];
  }
}
</script>

<template>
  <button
    @click='() => router.replace({ name: "browser" })'
    class="btn btn-primary col"
  >
    Browser
  </button>

  <code-mirror
    :dark="true"
    :lang="lang"
    :tab-size="2"
    basic
    tab
    v-model="code"
    wrap
  />

  <button
    @click="handle"
    class="btn btn-primary"
  >
    Run
  </button>

  <Table :data="tableData"></Table>
</template>

<style lang="scss">
.cm-editor, .vue-codemirror {
  height: 200px;
}
</style>
