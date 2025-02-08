<script setup lang="ts">
  import Connection from "../types/connection";
  import router from "../router";
  import { NFlex, NInput, NForm, NFormItem, NButton } from "naive-ui";
  import { useDatabaseStore } from "../stores";

  const db = useDatabaseStore();

  const conn: Connection = {
    host: "",
    password: "",
    port: 3306,
    user: "",
  };

  function handle() {
    db.save(conn);
    router.push("/");

    console.log(db.connections);
  }
</script>

<template>
  <n-flex justify="center">
    <n-form>
      <h2>Create MySQL Connection</h2>
      <n-form>
        <n-form-item label="Hostname">
          <n-input v-model="conn.host" type="text" placeholder="localhost" />
        </n-form-item>
        <n-form-item label="Username">
          <n-input v-model="conn.user" type="text" placeholder="root" />
        </n-form-item>
        <n-form-item label="Password">
          <n-input v-model="conn.password" type="password" placeholder="" />
        </n-form-item>
        <n-form-item label="Port">
          <n-input v-model="conn.port" type="text" placeholder="" default-value="3306" />
        </n-form-item>
        <n-form-item>
          <n-button @click="handle" type="primary">Save</n-button>
        </n-form-item>
      </n-form>
    </n-form>
  </n-flex>
</template>

<style scoped>
  form {
    flex: 0.5;
  }
</style>