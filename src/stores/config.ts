import { Connection } from "../types";
import { defineStore } from "pinia";

export const useConfigStore = defineStore("config", {
  state: () => ({
    connections: [] as Connection[],
  }),
  tauri: {
    saveOnChange: true,
  },
});