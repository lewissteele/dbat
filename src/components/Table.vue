<script setup lang="ts">
import { NTable } from "naive-ui";
import { defineProps } from "vue";

const { data } = defineProps({
  data: Array<Object>,
});

console.log(data);
</script>

<template>
  <n-table>
    <thead>
      <th>Header</th>
    </thead>
    <tbody>
      <tr v-for="row in data">
        <td>Row</td>
        <td>{{ row }}</td>
      </tr>
    </tbody>
  </n-table>
</template>
