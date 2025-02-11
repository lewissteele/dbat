import { createRouter, createWebHistory } from "vue-router";
import MainView from "./views/MainView.vue";
import SetupView from "./views/SetupView.vue";
import ConnectionsView from "./views/ConnectionsView.vue";

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      component: MainView,
      name: "main",
      path: "/",
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
  ],
});
