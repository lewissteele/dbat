import BrowserView from "./views/BrowserView.vue";
import ConnectionsView from "./views/ConnectionsView.vue";
import EditorView from "./views/EditorView.vue";
import SetupView from "./views/SetupView.vue";
import { createRouter, createWebHistory } from "vue-router";

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: "main",
      path: "/",
      redirect: "/connections",
    },
    {
      component: SetupView,
      name: "setup",
      path: "/setup",
    },
    {
      component: ConnectionsView,
      name: "connections",
      path: "/connections",
    },
    {
      component: BrowserView,
      name: "browser",
      path: "/browser",
    },
    {
      component: EditorView,
      name: "editor",
      path: "/editor",
    },
  ],
});
