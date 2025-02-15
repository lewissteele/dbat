<script setup lang="ts">
import CodeMirror from "vue-codemirror6";
import { ref } from "vue";
import { sql } from "@codemirror/lang-sql";

const lang = sql();
const code = ref("");
</script>

<template>
  <code-mirror
    v-model="code"
    :lang="lang"
    :dark="true"
    basic
    wrap
    tab
  />
</template>
