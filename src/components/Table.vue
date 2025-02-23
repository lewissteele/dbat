<script setup lang="ts">
import { NTable } from "naive-ui";
import { computed } from "vue";

const { data } = defineProps<{
  data: any[];
}>();

const headers = computed(() => {
  if (data.length) {
    return Object.keys(data[0]);
  }

  return [];
});
</script>

<template>
  <n-table>
    <thead>
      <th v-for="header in headers">
        {{ header }}
      </th>
    </thead>
    <tbody>
      <tr v-for="row in data">
        <td v-for="data in row">
          {{ data }}
        </td>
      </tr>
    </tbody>
  </n-table>
</template>
