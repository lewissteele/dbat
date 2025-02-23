import Database from "@tauri-apps/plugin-sql";
import { Connection } from "../types";
import { defineStore } from "pinia";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    connection: null as Database | null,
  }),
  actions: {
    async connect(connection: Connection): Promise<void> {
      console.log(connection);
      this.connection = await Database.load("mysql://root@localhost/search");
    },
  },
});
