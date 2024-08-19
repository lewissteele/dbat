import { Args, Command } from "@oclif/core";
import fs from "fs-extra";
import path from "path";

export default class Remove extends Command {
  static args = {
    database: Args.string(),
  };
  static description = "delete database connection";

  async run() {
    const { args } = await this.parse(Remove);

    const configPath = path.join(this.config.configDir, "config.json");
    const config = (await fs.readJson(configPath, { throws: false })) ?? {};

    delete config.databases[args.database];

    await fs.writeJson(configPath, config);
  }
}
