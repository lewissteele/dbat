const BaseCommand = require("../base-command");

export default class List extends BaseCommand {
  static description = "show databases";
  async run() {
    const config = await this.getConfig();
    const databases = Object.keys(config.databases);

    if (!databases.length) {
      this.log("no databases");
      return;
    }

    databases.forEach((val) => this.log(val));
  }
}
