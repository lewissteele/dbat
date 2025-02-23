<script setup lang="ts">
import { storeToRefs } from 'pinia';
import { useDatabaseStore } from '../stores';
import { NList, NListItem } from 'naive-ui';

const db = useDatabaseStore();

const { tables } = storeToRefs(db);

</script>

<template>
  <n-list>
    <n-list-item v-for="table in tables">
      {{ table.name }}
    </n-list-item>
  </n-list>
</template>