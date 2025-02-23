import { Connection } from "../types";
import { defineStore } from "pinia";

export const useConfigStore = defineStore("config", {
  state: () => ({
    activeConnection: null as Connection | null,
    connections: [] as Connection[],
  }),
  tauri: {
    saveOnChange: true,
  },
});
