import App from "./App.vue";
import router from "./router";
import { createApp } from "vue";
import { createPinia } from "pinia";
import { createPlugin } from "tauri-plugin-pinia";
import { useConfigStore, useDatabaseStore } from "./stores";

(async () => {
  const app = createApp(App);

  app.use(router);
  app.use(createPinia().use(createPlugin()));

  const config = useConfigStore();
  const db = useDatabaseStore();

  await config.$tauri.start();
  app.mount("#app");

  if (config.activeConnection) {
    await db.connect(config.activeConnection);
    router.replace({ name: "editor" });
    return;
  }

  if (config.connections.length) {
    router.replace({ name: "connections" });
    return;
  }

  router.replace({ name: "setup" });
})();
