<script setup lang="ts">
import router from "../router";
import { Connection, Driver } from "../types";
import { ref } from "vue";
import { useConfigStore } from "../stores";

const config = useConfigStore();

const connection = ref({
  database: "",
  driver: Driver.MYSQL,
  host: "localhost",
  password: "",
  port: "3306",
  user: "root",
});

function saveConnection(): void {
  config.connections.push(connection.value);
  router.replace({ name: "connections" });
}
</script>

<template>
  <div class="container p-2">
    <form @submit.prevent="saveConnection" spellcheck="false">
      <div class="mb-3">
        <select class="form-select" v-model="connection.driver">
          <option v-for="driver in Driver">
            {{ driver }}
          </option>
        </select>
      </div>
      <div class="form-floating mb-3">
        <input
          class="form-control"
          placeholder="localhost"
          required
          type="text"
          v-model="connection.host"
        >
        <label>Host</label>
      </div>
      <div class="form-floating mb-3">
        <input
          class="form-control"
          placeholder="laravel"
          type="text"
          v-model="connection.database"
          required
        >
        <label>Database</label>
      </div>
      <div class="form-floating mb-3">
        <input
          class="form-control"
          placeholder="root"
          required
          type="text"
          v-model="connection.user"
        >
        <label>Password</label>
      </div>
      <div class="form-floating mb-3">
        <input
          class="form-control"
          placeholder="password"
          type="password"
          v-model="connection.password"
        >
        <label>Password</label>
      </div>
      <div class="form-floating mb-3">
        <input
          class="form-control"
          placeholder="3306"
          required
          type="text"
          v-model="connection.port"
        >
        <label>Port</label>
      </div>
      <button type="submit" class="btn btn-primary">
        Save
      </button>
    </form>
  </div>
</template>
