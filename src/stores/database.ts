import { defineStore } from "pinia";
import Connection from "../types/connection";
import Database from "@tauri-apps/plugin-sql";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    _conn: null as Database | null,
    active: null as Connection | null,
    saved: [] as Connection[],
  }),
  getters: {
    async conn(): Promise<Database> {
      if (this._conn) {
        return this._conn;
      }

      return this._conn = await Database.load("mysql://root@localhost/search");
    },
  },
  actions: {
    save(connection: Connection) {
      this.saved.push(connection);
    },
    setActive(connection: Connection) {
      this._conn = null;
      this.active = connection;
    },
  },
  tauri: {
    saveOnChange: true,
  },
});
