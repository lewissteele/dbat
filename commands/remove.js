import BaseCommand from "../base-command.js";
import { Args } from "@oclif/core";

export default class Remove extends BaseCommand {
  static args = {
    database: Args.string(),
  };
  static description = "delete database connection";

  async run() {
    const { args } = await this.parse(Remove);

    const config = await this.getConfig();

    delete config.databases[args.database];

    await this.setConfig(config);
  }
}
