<script setup lang="ts">
import SqlEditor from "../components/SqlEditor.vue";
import TableList from "../components/TableList.vue";
import { NTabPane, NTabs } from "naive-ui";
</script>

<template>
  <n-tabs type="segment" animated>
    <n-tab-pane name="browser" tab="Browser">
      <TableList />
    </n-tab-pane>
    <n-tab-pane name="editor" tab="Editor">
      <SqlEditor />
    </n-tab-pane>
  </n-tabs>
</template>
