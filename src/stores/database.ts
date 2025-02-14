import { defineStore } from "pinia";
import Connection from "../types/connection";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    connections: [] as Connection[],
  }),
  getters: {
    hasConnection(): boolean {
      return this.connections.length > 0;
    },
  },
  actions: {
    save(connection: Connection) {
      this.connections.push(connection);
    },
    empty() {
      this.connections = [] as Connection[];
    },
  },
  tauri: {
    saveOnChange: true,
  },
});
