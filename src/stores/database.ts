import Database from "@tauri-apps/plugin-sql";
import { Connection, Driver } from "../types";
import { defineStore } from "pinia";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    connection: null as Database | null,
  }),
  actions: {
    async connect(connection: Connection): Promise<void> {
      let path = "";

      switch (connection.driver) {
        case Driver.MYSQL:
        case Driver.POSTGRES:
          path = `${connection.driver}://${connection.user}@${connection.host}/${connection.database}`;
          break;
      }

      this.connection = await Database.load(path);
    },
  },
});
