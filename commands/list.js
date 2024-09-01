const BaseCommand = require("../base-command");

module.exports = class List extends BaseCommand {
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
