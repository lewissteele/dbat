<script setup lang="ts">
import router from "../router";
import { Driver } from "../types";
import { ref } from "vue";
import { useConfigStore } from "../stores";
import {
  NButton,
  NFlex,
  NForm,
  NFormItem,
  NInput,
  NSelect,
  NSpace,
} from "naive-ui";

const config = useConfigStore();
const drivers = Object.values(Driver).map((driver) => ({
  label: driver,
  selected: driver == Driver.MYSQL,
  value: driver,
}));

const conn = ref({
  database: "",
  driver: Driver.MYSQL,
  host: "",
  password: "",
  port: "3306",
  user: "",
});

function handle(): void {
  config.connections.push(conn.value);
  router.replace({ name: "connections" });
}
</script>

<template>
  <n-flex justify="center">
    <n-form>
      <h2>Setup</h2>
      <n-form :model="conn">
        <n-form-item label="Driver">
          <n-select
            v-model:value="conn.driver"
            :options="drivers"
          />
        </n-form-item>
        <n-form-item label="Hostname">
          <n-input
            v-model:value="conn.host"
            type="text"
            placeholder="localhost"
            spellcheck="false"
          />
        </n-form-item>
        <n-form-item label="Username">
          <n-input
            v-model:value="conn.user"
            type="text"
            placeholder="root"
            spellcheck="false"
          />
        </n-form-item>
        <n-form-item label="Password">
          <n-input
            v-model:value="conn.password"
            type="password"
            placeholder=""
          />
        </n-form-item>
        <n-form-item label="Database">
          <n-input
            v-model:value="conn.database"
            type="text"
            placeholder="laravel"
          />
        </n-form-item>
        <n-form-item label="Port">
          <n-input v-model:value="conn.port" type="text" placeholder="" />
        </n-form-item>
        <n-space justify="end">
          <n-button @click="handle" type="primary" size="large">Save</n-button>
        </n-space>
      </n-form>
    </n-form>
  </n-flex>
</template>

<style scoped>
form {
  flex: 0.5;
}
</style>
