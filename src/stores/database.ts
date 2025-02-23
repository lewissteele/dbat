import Database from "@tauri-apps/plugin-sql";
import { Connection, Driver, Table } from "../types";
import { defineStore } from "pinia";

export const useDatabaseStore = defineStore("database", {
  state: () => ({
    connection: null as Database | null,
    tables: [] as Table[],
  }),
  actions: {
    async connect(connection: Connection): Promise<void> {
      let path = "";

      switch (connection.driver) {
        case Driver.MYSQL:
        case Driver.POSTGRES:
          path =
            `${connection.driver}://${connection.user}@${connection.host}/${connection.database}`;
          break;
      }

      this.connection = await Database.load(path);

      this.refresh();
    },
    async refresh(): Promise<void> {
      if (!this.connection) {
        return;
      }

      const tables = await this.connection.select("show tables") as any[];

      this.tables = tables.map((table) => ({
        name: Object.values(table)[0] as string,
      }));
    },
  },
});
