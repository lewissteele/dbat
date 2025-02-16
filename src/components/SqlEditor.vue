<script setup lang="ts">
import CodeMirror from "vue-codemirror6";
import { ref } from "vue";
import { sql, MySQL } from "@codemirror/lang-sql";
import { NButton } from "naive-ui";

const lang = sql({
  dialect: MySQL,
});
const code = ref("");

function handle(): void {
  console.log(code.value);
}
</script>

<template>
  <code-mirror
    v-model="code"
    :lang=lang
    :dark=true
    :tab-size=2
    basic
    wrap
    tab
  />
  <n-button @click="handle">Run</n-button>
</template>

<style>
.cm-editor {
  height: 500px;
}
</style>
