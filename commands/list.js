import fs from "fs-extra";
import path from "path";
import { Command } from "@oclif/core";

export default class List extends Command {
  static description = "show databases";
  async run() {
    const configPath = path.join(this.config.configDir, "config.json");
    const config = (await fs.readJson(configPath, { throws: false })) ?? {};
    const databases = Object.keys(config.databases);

    if (!databases.length) {
      this.log("no databases");
      return;
    }

    databases.forEach((val) => this.log(val));
  }
}
