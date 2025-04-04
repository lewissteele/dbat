<script setup lang="ts">
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
  <table class="table table-striped table-bordered">
    <thead>
      <tr>
        <th v-for="header in headers">
          {{ header }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="row in data">
        <td v-for="data in row">
          {{ data }}
        </td>
      </tr>
    </tbody>
  </table>
</template>
