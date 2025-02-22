import { defineStore } from "pinia";
import Connection from "../types/connection";
import Database from "@tauri-apps/plugin-sql";
import { v4 as uuid } from "uuid";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    _reader: null as Database | null,
    _writer: null as Database | null,
    active: null as Connection | null,
    saved: [] as Connection[],
  }),
  actions: {
    save(connection: {
      host: string;
      password: string;
      port: string;
      user: string;
    }) {
      this.saved.push({
        ...connection,
        uuid: uuid(),
      });
    },
    setActive(connection: Connection) {
      this._reader = null;
      this._writer = null;
      this.active = connection;
    },
    async reader(): Promise<Database> {
      return await Database.load("mysql://root@localhost/search");
    },
  },
  tauri: {
    saveOnChange: true,
  },
});
