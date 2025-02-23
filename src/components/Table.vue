<script setup lang="ts">
import { NTable } from "naive-ui";

const { data } = defineProps<{
  data: any[];
}>();

let headers = [] as any[];

if (data.length) {
  headers = Object.keys(data[0]);
}
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
