import App from "./App.vue";
import router from "./router";
import { createApp } from "vue";
import { createPinia } from "pinia";
import { createPlugin } from "tauri-plugin-pinia";
import { useDatabaseStore } from "./stores";

(async () => {
  const app = createApp(App);

  app.use(router);
  app.use(createPinia().use(createPlugin()));

  const db = useDatabaseStore();
  await db.$tauri.start();

  app.mount("#app");

  if (db.connections) {
    router.replace({ name: "connections" });
    return;
  }

  router.replace({ name: "setup" });
})();
