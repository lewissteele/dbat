import { createRouter, createWebHistory } from "vue-router";
import MainView from "./views/MainView.vue";
import SetupView from "./views/SetupView.vue";
import ConnectionsView from "./views/ConnectionsView.vue";

export default createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "main",
      component: MainView,
    },
    {
      path: "/setup",
      name: "setup",
      component: SetupView,
    },
    {
      path: "/connections",
      name: "connections",
      component: ConnectionsView,
    },
  ],
});
