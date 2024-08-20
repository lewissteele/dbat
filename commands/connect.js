import fs from "fs-extra";
import path from "path";
import { Args, Command } from "@oclif/core";
import readline from "node:readline/promises";
import { stdin, stdout } from "node:process";
import { Sequelize } from "sequelize";

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
      dialect: 'sqlite',
      logging: false,
      storage: 'database.sqlite',
    });

    const rl = readline.createInterface({
      input: stdin,
      output: stdout,
    });

    rl.prompt();

    rl.on("line", async (line) => {
      const [results, meta] = await sequelize.query(line)
      this.log(results, meta)
      rl.prompt();
    });
  }
}
