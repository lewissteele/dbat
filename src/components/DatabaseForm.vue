<script setup lang="ts">
  import router from "../router";
  import { NFlex, NInput, NForm, NFormItem, NButton } from "naive-ui";
  import { useDatabaseStore } from "../stores";
  import { ref } from "vue";

  const db = useDatabaseStore();

  const conn = ref({
    host: "",
    password: "",
    port: "3306",
    user: "",
  });

  function handle() {
    db.save(conn.value);
    router.replace({ name: "main" });
  }
</script>

<template>
  <n-flex justify="center">
    <n-form>
      <h2>MySQL Connection</h2>
      <n-form :model="conn">
        <n-form-item label="Hostname">
          <n-input v-model:value="conn.host" type="text" placeholder="localhost" spellcheck="false" />
        </n-form-item>
        <n-form-item label="Username">
          <n-input v-model:value="conn.user" type="text" placeholder="root" spellcheck="false" />
        </n-form-item>
        <n-form-item label="Password">
          <n-input v-model:value="conn.password" type="password" placeholder="" />
        </n-form-item>
        <n-form-item label="Port">
          <n-input v-model:value="conn.port" type="text" placeholder="" />
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