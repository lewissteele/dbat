import { defineStore } from "pinia";
import Connection from "../types/connection";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    connections: [] as Connection[],
    active: null as Connection | null,
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
    setActive(connection: Connection) {
      this.active = connection;
    },
  },
  tauri: {
    saveOnChange: true,
  },
});
