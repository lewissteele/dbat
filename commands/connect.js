import fs from "fs-extra";
import path from "path";
import { Args, Command } from "@oclif/core";
import readline from "node:readline/promises";
import { stdin, stdout } from "node:process";
import { Sequelize } from "sequelize";
import Table from "cli-table3";

export default class Connect extends Command {
  static args = {
    database: Args.string(),
  };

  async run() {
    const { args } = await this.parse(Connect);

    const configPath = path.join(this.config.configDir, "config.json");
    const config = await fs.readJson(configPath, { throws: false });
    const database = config.databases[args.database];

    if (database == undefined) {
      this.log("no such database");
      return;
    }

    const sequelize = new Sequelize({
      ...database,
      logging: false,
    });

    const rl = readline.createInterface({
      input: stdin,
      output: stdout,
    });

    rl.prompt();

    rl.on("line", async (line) => {
      const [results, meta] = await sequelize.query(line, { raw: true });

      if (!results.length) {
        rl.prompt();
        return;
      }

      const headers = Object.keys(results[0]);
      const rows = results.map((row) => Object.values(row));

      const table = new Table({
        head: headers,
        rows: rows,
      });

      table.push(...rows);

      this.log(table.toString());

      rl.prompt();
    });
  }
}
